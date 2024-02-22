package main

import (
	"io"
	"os"

	"github.com/urfave/cli/v2"
	"golang.design/x/clipboard"
)

func DefaultAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		clipboard.Init()

		snippet := state.Snippets[cCtx.Args().First()]

		if snippet != "" {
			clipboard.Write(clipboard.FmtText, []byte(snippet))
			cCtx.App.Writer.Write([]byte("Copied!\n"))
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
			cCtx.App.Writer.Write([]byte(snippet))
			return nil
		}

		cCtx.App.ErrWriter.Write([]byte("Snippet not found\n"))
		return nil
	}
}

func ListAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		for name := range state.Snippets {
			cCtx.App.Writer.Write([]byte(name + "\n"))
		}
		return nil
	}
}

func AddAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		var content []byte
		var err error

		if state.Name == "" {
			cCtx.App.ErrWriter.Write([]byte("Name is required\n"))
			return cli.Exit("", 1)
		}

		if state.InputFile == "" {
			content, err = io.ReadAll(os.Stdin)
		} else {
			content, err = os.ReadFile(state.InputFile)
		}

		if err != nil {
			cCtx.App.ErrWriter.Write([]byte("Cannot read content\n"))
			return cli.Exit("", 1)
		}

		if len(content) == 0 {
			cCtx.App.ErrWriter.Write([]byte("Snippet content is required\n"))
			return cli.Exit("", 1)
		}

		state.Snippets[state.Name] = string(content)
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
