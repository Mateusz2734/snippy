package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strings"
	"time"

	"github.com/alecthomas/chroma/quick"
	"github.com/andrew-d/go-termutil"
	"github.com/urfave/cli/v2"
	"golang.design/x/clipboard"
)

func DefaultAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		err := clipboard.Init()

		if err != nil {
			cCtx.App.ErrWriter.Write([]byte("Cannot initialize clipboard\n"))
			return cli.Exit("", 1)
		}

		snippet, ok := state.Snippets[cCtx.Args().First()]

		if !ok {
			return nil
		}

		if snippet.Content != "" {
			clipboard.Write(clipboard.FmtText, []byte(snippet.Content))
			cCtx.App.Writer.Write([]byte("Copied!\n"))
			return nil
		}

		return nil
	}
}

func GetAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		if state.Name == "" {
			cCtx.App.ErrWriter.Write([]byte("Name is required\n"))
			return cli.Exit("", 1)
		}

		if snippet, ok := state.Snippets[state.Name]; ok {
			if !state.NoMetadata {
				printMetadata(cCtx, snippet)
			}

			quick.Highlight(cCtx.App.Writer, snippet.Content, snippet.Extension, "terminal256", "dracula")
			return nil
		}

		cCtx.App.ErrWriter.Write([]byte("Snippet not found\n"))
		return nil
	}
}

func ListAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		if len(state.Snippets) == 0 {
			cCtx.App.Writer.Write([]byte("No snippets found\n"))
			return nil
		}

		sortedSnippets := sortByFavorite(state.Snippets, false)

		iterator := (state.CurrentPage - 1) * state.PageSize

		if iterator >= len(sortedSnippets) {
			cCtx.App.Writer.Write([]byte("This page is empty\n"))
			return nil
		}

		for i := iterator; i < len(sortedSnippets) && i < iterator+state.PageSize; i++ {
			cCtx.App.Writer.Write([]byte(sortedSnippets[i] + "\n"))
		}

		return nil
	}
}

func AddAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		var content []byte
		var err error

		if err = clipboard.Init(); err != nil {
			cCtx.App.ErrWriter.Write([]byte("Cannot initialize clipboard\n"))
			return cli.Exit("", 1)
		}

		if state.Name == "" {
			cCtx.App.ErrWriter.Write([]byte("Name is required\n"))
			return cli.Exit("", 1)
		}

		if state.InputFile != "" {
			content, err = os.ReadFile(state.InputFile)
		} else if !termutil.Isatty(os.Stdin.Fd()) {
			content, err = io.ReadAll(os.Stdin)
		} else if state.UseClipboard {
			content = clipboard.Read(clipboard.FmtText)
		}

		if err != nil {
			cCtx.App.ErrWriter.Write([]byte("Cannot read content\n"))
			return cli.Exit("", 1)
		}

		if content == nil || len(content) == 0 {
			cCtx.App.ErrWriter.Write([]byte("Snippet content is required\n"))
			return cli.Exit("", 1)
		}

		snippet := &Snippet{Content: string(content), CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}

		if state.Extension != "" {
			snippet.Extension = state.Extension
		}

		state.Snippets[state.Name] = snippet

		cCtx.App.Writer.Write([]byte("Snippet added successfully\n"))
		return nil
	}
}

func DeleteAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		if state.Name == "" {
			cCtx.App.ErrWriter.Write([]byte("Name is required\n"))
			return cli.Exit("", 1)
		}

		if _, ok := state.Snippets[state.Name]; ok {
			delete(state.Snippets, state.Name)
			cCtx.App.Writer.Write([]byte("Snippet deleted successfully\n"))
			return nil
		}

		cCtx.App.ErrWriter.Write([]byte("Snippet not found\n"))
		return nil
	}
}

func EditAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		var ok bool
		var snippet *Snippet
		var extension = "txt"

		if state.Name == "" {
			cCtx.App.ErrWriter.Write([]byte("Name is required\n"))
			return cli.Exit("", 1)
		}

		snippets := state.Snippets

		if snippet, ok = snippets[state.Name]; !ok {
			cCtx.App.ErrWriter.Write([]byte("Snippet not found\n"))
			return nil
		}

		if state.Extension != "" {
			snippet.Extension = state.Extension
			cCtx.App.Writer.Write([]byte("Snippet updated successfully\n"))
			return nil
		}

		if snippet.Extension != "" {
			extension = snippet.Extension
		}

		tmpFile, err := os.CreateTemp("", fmt.Sprintf("snippet-*.%s", extension))
		if err != nil {
			cCtx.App.ErrWriter.Write([]byte("Cannot create snippet file\n"))
			return cli.Exit("", 1)
		}
		defer os.Remove(tmpFile.Name())

		if _, err := tmpFile.WriteString(snippet.Content); err != nil {
			cCtx.App.ErrWriter.Write([]byte("Cannot write snippet to file\n"))
			return cli.Exit("", 1)
		}

		tmpFile.Close()

		err = launchTextEditor(tmpFile.Name())
		if err != nil {
			cCtx.App.ErrWriter.Write([]byte("Error launching text editor\n"))
			return cli.Exit("", 1)
		}

		modifiedSnippet, err := os.ReadFile(tmpFile.Name())
		if err != nil {
			cCtx.App.ErrWriter.Write([]byte("Error reading modified snippet\n"))
			return cli.Exit("", 1)
		}

		elem := state.Snippets[state.Name]
		elem.Content = string(modifiedSnippet)
		elem.UpdatedAt = time.Now().Unix()
		cCtx.App.Writer.Write([]byte("Snippet updated successfully\n"))
		return nil
	}
}

func SearchAction(state *State) func(cCtx *cli.Context) error {
	extensionMatches := func(snippet *Snippet) bool {
		return state.Extension == "" || snippet.Extension == state.Extension
	}

	nameMatches := func(key string) bool {
		return state.Name == "" || strings.Contains(key, state.Name)
	}

	return func(cCtx *cli.Context) error {
		filteredSnippets := make([]string, 0, len(state.Snippets))

		for key, snippet := range state.Snippets {
			if extensionMatches(snippet) && nameMatches(key) {
				filteredSnippets = append(filteredSnippets, key)
			}
		}

		if len(filteredSnippets) == 0 {
			cCtx.App.Writer.Write([]byte("No snippets found\n"))
			return nil
		}

		for _, snippet := range filteredSnippets {
			cCtx.App.Writer.Write([]byte(snippet + "\n"))
		}

		return nil
	}
}

func FavoriteAddAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		name := state.Name

		if name == "" {
			cCtx.App.ErrWriter.Write([]byte("Name is required\n"))
			return cli.Exit("", 1)
		}

		if snippet, ok := state.Snippets[name]; ok {
			snippet.Favorite = true
			cCtx.App.Writer.Write([]byte("Favorite added successfully\n"))
			return nil
		}

		cCtx.App.ErrWriter.Write([]byte("Snippet not found\n"))
		return nil
	}
}

func FavoriteDeleteAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		name := state.Name

		if name == "" {
			cCtx.App.ErrWriter.Write([]byte("Name is required\n"))
			return cli.Exit("", 1)
		}

		if snippet, ok := state.Snippets[name]; ok {
			snippet.Favorite = false
			cCtx.App.Writer.Write([]byte("Favorite deleted successfully\n"))
			return nil
		}

		cCtx.App.ErrWriter.Write([]byte("Snippet not found\n"))
		return nil
	}
}

func FavoriteListAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		filteredSnippets := sortByFavorite(state.Snippets, true)

		if len(filteredSnippets) == 0 {
			cCtx.App.Writer.Write([]byte("No favorite snippets found\n"))
			return nil
		}

		iterator := (state.CurrentPage - 1) * state.PageSize

		if iterator >= len(filteredSnippets) {
			cCtx.App.Writer.Write([]byte("This page is empty\n"))
			return nil
		}

		for i := iterator; i < len(filteredSnippets) && i < iterator+state.PageSize; i++ {
			cCtx.App.Writer.Write([]byte(filteredSnippets[i] + "\n"))
		}

		return nil
	}
}

func launchTextEditor(filepath string) error {
	flag := ""

	editor := existingEditor()
	if editor == "" {
		return fmt.Errorf("no text editor found")
	}

	if editor == "code" {
		flag = "--wait"
	}

	cmd := exec.Command(editor, flag, filepath)

	return cmd.Run()
}

func existingEditor() string {
	var cmd *exec.Cmd
	editors := []string{"code", "nano", "vim", "emacs", "vi", "notepad"}

	for _, editor := range editors {
		cmd = exec.Command(editor, "--version")

		if editor == "notepad" && runtime.GOOS == "windows" {
			cmd = exec.Command("where", "notepad")
		}

		err := cmd.Run()
		if err == nil {
			return editor
		}
	}
	return ""
}

func printMetadata(cCtx *cli.Context, snippet *Snippet) {
	cCtx.App.Writer.Write([]byte("Metadata = { "))
	if snippet.CreatedAt != 0 {
		cCtx.App.Writer.Write([]byte(fmt.Sprintf("Created: %d ", snippet.CreatedAt)))
	}
	if snippet.UpdatedAt != 0 {
		cCtx.App.Writer.Write([]byte(fmt.Sprintf("Updated: %d ", snippet.UpdatedAt)))
	}
	if snippet.Extension != "" {
		cCtx.App.Writer.Write([]byte(fmt.Sprintf("Extension: %s ", snippet.Extension)))
	}
	cCtx.App.Writer.Write([]byte("}\n\n"))
}

func sortByFavorite(snippets map[string]*Snippet, onlyFavorites bool) []string {
	favorites := make([]string, 0, len(snippets))
	rest := make([]string, 0, len(snippets))

	for key, snippet := range snippets {
		if snippet.Favorite {
			favorites = append(favorites, key)
		} else if !onlyFavorites {
			rest = append(rest, key)
		}
	}

	slices.Sort(favorites)
	slices.Sort(rest)

	return slices.Concat(favorites, rest)
}
