package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func SetupTestApp() (*State, *cli.App, *bytes.Buffer, *bytes.Buffer) {
	global := make(map[string]*Snippet)
	local := make(map[string]*Snippet)

	state := NewState()

	state.InitializeSnippets(global, local)

	app := NewApp(state, []cli.Flag{WithGlobal(state), WithNoFormatting(state)})

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	app.Before = func(cCtx *cli.Context) error { return nil }
	app.ExitErrHandler = func(cCtx *cli.Context, err error) {}
	app.Writer = stdout
	app.ErrWriter = stderr

	for _, command := range app.Commands {
		command.After = func(*cli.Context) error { return nil }
		for _, subcommand := range command.Subcommands {
			subcommand.After = func(*cli.Context) error { return nil }
		}
	}

	return state, app, stdout, stderr
}

func TestDefaultAction(t *testing.T) {
	state, app, stdout, stderr := SetupTestApp()

	err := app.Run([]string{"snippy"})
	assert.Nil(t, err, "error should be nil")
	assert.Empty(t, stdout.String(), "stdout should be empty")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "test"})
	assert.Nil(t, err, "error should be nil")
	assert.Empty(t, stdout.String(), "stdout should be empty")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	state.localSnippets["test"] = &Snippet{Content: "test content"}

	err = app.Run([]string{"snippy", "test"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "Copied!", "stdout should contain success message")
	assert.Empty(t, stderr.String(), "stderr should be empty")
}

func TestGetCommand(t *testing.T) {
	state, app, stdout, stderr := SetupTestApp()

	err := app.Run([]string{"snippy", "get", "-h"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "snippy get", "stdout should contain help message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "get"})
	assert.NotNil(t, err, "error should not be nil")
	assert.Empty(t, stdout.String(), "stdout should be empty")
	assert.Contains(t, stderr.String(), "Name is required", "stderr should contain error message")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "get", "--name", "test"})
	assert.Nil(t, err, "error should be nil")
	assert.Empty(t, stdout.String(), "stdout should be empty")
	assert.Contains(t, stderr.String(), "Snippet not found", "stderr should contain error message")

	stdout.Reset()
	stderr.Reset()

	state.localSnippets["test"] = &Snippet{Content: "test content", Extension: "go"}

	err = app.Run([]string{"snippy", "get", "-n", "test", "--nf"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "Metadata", "stdout should contain snippet metadata")
	assert.Contains(t, stdout.String(), "Extension: go", "stdout should contain snippet extension")
	assert.Contains(t, stdout.String(), "test content", "stdout should contain snippet content")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "get", "-n", "test", "--no-metadata", "--nf"})
	assert.Nil(t, err, "error should be nil")
	assert.NotContains(t, stdout.String(), "Metadata", "stdout should not contain snippet metadata")
	assert.Contains(t, stdout.String(), "test content", "stdout should contain snippet content")
	assert.Empty(t, stderr.String(), "stderr should be empty")
}

func TestListCommand(t *testing.T) {
	state, app, stdout, stderr := SetupTestApp()

	err := app.Run([]string{"snippy", "list", "-h"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "snippy list", "stdout should contain help message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "list"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "No snippets found", "stdout should be empty")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	state.localSnippets["test"] = &Snippet{Content: "test content"}

	err = app.Run([]string{"snippy", "list"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "test", "stdout should contain snippet name")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "list", "-g"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "No snippets found", "stdout should contain snippet name")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "list", "-p", "4"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "This page is empty", "stdout should be empty")
	assert.Empty(t, stderr.String(), "stderr should be empty")
}

func TestAddCommand(t *testing.T) {
	state, app, stdout, stderr := SetupTestApp()

	err := app.Run([]string{"snippy", "add", "-h"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "snippy add", "stdout should contain help message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "add"})
	assert.NotNil(t, err, "error should not be nil")
	assert.Empty(t, stdout.String(), "stdout should be empty")
	assert.Contains(t, stderr.String(), "Name is required", "stderr should contain error message")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "add", "--name", "test"})
	assert.NotNil(t, err, "error should not be nil")
	assert.Empty(t, stdout.String(), "stdout should be empty")
	assert.Contains(t, stderr.String(), "Snippet content is required", "stderr should contain error message")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "add", "-n", "test", "--e", "go", "--content", "test content"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "Snippet added successfully", "stdout should contain success message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	snippet := state.localSnippets["test"]

	assert.Equal(t, "test content", snippet.Content, "snippet content should be set")
	assert.Equal(t, "go", snippet.Extension, "snippet extension should be set")
	assert.NotZero(t, snippet.CreatedAt, "snippet CreatedAt should be set")
	assert.NotZero(t, snippet.UpdatedAt, "snippet UpdatedAt should be set")

	stdout.Reset()
	stderr.Reset()
	delete(state.localSnippets, "test")

	err = app.Run([]string{"snippy", "add", "-n", "test", "--content", "test content", "--global"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "Snippet added successfully", "stdout should contain success message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	snippet = state.globalSnippets["test"]
	_, ok := state.localSnippets["test"]

	assert.False(t, ok, "local snippet should not exist")
	assert.Equal(t, "test content", snippet.Content, "snippet content should be set")
	assert.NotZero(t, snippet.CreatedAt, "snippet CreatedAt should be set")
	assert.NotZero(t, snippet.UpdatedAt, "snippet UpdatedAt should be set")
}

func TestDeleteCommand(t *testing.T) {
	state, app, stdout, stderr := SetupTestApp()

	err := app.Run([]string{"snippy", "delete", "-h"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "snippy delete", "stdout should contain help message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "delete"})
	assert.NotNil(t, err, "error should not be nil")
	assert.Empty(t, stdout.String(), "stdout should be empty")
	assert.Contains(t, stderr.String(), "Name is required", "stderr should contain error message")

	stdout.Reset()
	stderr.Reset()

	state.localSnippets["test"] = &Snippet{Content: "test content"}

	err = app.Run([]string{"snippy", "delete", "-n", "test"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "Snippet deleted successfully", "stdout should contain success message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "delete", "--name", "test"})
	assert.Nil(t, err, "error should be nil")
	assert.Empty(t, stdout.String(), "stdout should be empty")
	assert.Contains(t, stderr.String(), "Snippet not found", "stderr should contain error message")
}

func TestEditCommand(t *testing.T) {
	state, app, stdout, stderr := SetupTestApp()

	err := app.Run([]string{"snippy", "edit", "-h"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "snippy edit", "stdout should contain help message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "edit"})
	assert.NotNil(t, err, "error should not be nil")
	assert.Empty(t, stdout.String(), "stdout should be empty")
	assert.Contains(t, stderr.String(), "Name is required", "stderr should contain error message")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "edit", "--name", "test"})
	assert.Nil(t, err, "error should be nil")
	assert.Empty(t, stdout.String(), "stdout should be empty")
	assert.Contains(t, stderr.String(), "Snippet not found", "stderr should contain error message")

	stdout.Reset()
	stderr.Reset()

	state.localSnippets["test"] = &Snippet{Content: "test content"}

	err = app.Run([]string{"snippy", "edit", "-n", "test", "--ext", "go"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "Snippet updated successfully", "stdout should contain snippet content")
	assert.Empty(t, stderr.String(), "stderr should be empty")
	assert.Equal(t, "go", state.localSnippets["test"].Extension, "snippet extension should be updated")
}

func TestSearchCommand(t *testing.T) {
	state, app, stdout, stderr := SetupTestApp()

	err := app.Run([]string{"snippy", "search", "-h"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "snippy search", "stdout should contain help message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "search", "--name", "test"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "No snippets found", "stdout should be empty")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	state.localSnippets["test"] = &Snippet{Content: "test content"}
	state.localSnippets["test1"] = &Snippet{Content: "test content", Extension: "go"}

	err = app.Run([]string{"snippy", "search", "-n", "test"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "test", "stdout should contain snippet name")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "search", "--ext", "go"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "test1", "stdout should contain snippet name")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "search"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "test", "stdout should contain name of first snippet")
	assert.Contains(t, stdout.String(), "test1", "stdout should contain name of second snippet")
	assert.Empty(t, stderr.String(), "stderr should be empty")
}

func TestFavoriteCommand(t *testing.T) {
	_, app, stdout, stderr := SetupTestApp()

	err := app.Run([]string{"snippy", "favorite", "-h"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "USAGE", "stdout should contain help message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "favorite"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "USAGE", "stdout should contain help message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "favorite", "add", "-h"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "snippy favorite add", "stdout should contain help message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "favorite", "delete", "-h"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "snippy favorite delete", "stdout should contain help message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "favorite", "list", "-h"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "snippy favorite list", "stdout should contain help message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

}

func TestFavoriteListCommand(t *testing.T) {
	state, app, stdout, stderr := SetupTestApp()

	err := app.Run([]string{"snippy", "favorite", "list", "-h"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "snippy favorite list", "stdout should contain help message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "favorite", "list"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "No favorite snippets found", "stdout should be empty")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	state.localSnippets["test"] = &Snippet{Content: "test content", Favorite: false}

	err = app.Run([]string{"snippy", "favorite", "list"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "No favorite snippets found", "stdout should contain snippet name")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	state.localSnippets["test"] = &Snippet{Content: "test content", Favorite: true}

	err = app.Run([]string{"snippy", "favorite", "list"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "test", "stdout should contain snippet name")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "favorite", "list", "-p", "4"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "This page is empty", "stdout should be empty")
	assert.Empty(t, stderr.String(), "stderr should be empty")
}

func TestFavoriteAddCommand(t *testing.T) {
	state, app, stdout, stderr := SetupTestApp()

	err := app.Run([]string{"snippy", "favorite", "add", "-h"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "snippy favorite add", "stdout should contain help message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "favorite", "add"})
	assert.NotNil(t, err, "error should not be nil")
	assert.Empty(t, stdout.String(), "stdout should be empty")
	assert.Contains(t, stderr.String(), "Name is required", "stderr should contain error message")

	stdout.Reset()
	stderr.Reset()

	state.localSnippets["test"] = &Snippet{Content: "test content"}

	err = app.Run([]string{"snippy", "favorite", "add", "--name", "test"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "Favorite added successfully", "stdout should contain success message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	state.localSnippets["test"] = &Snippet{Content: "test content", Favorite: true}

	err = app.Run([]string{"snippy", "favorite", "add", "-n", "test"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "Favorite added successfully", "stdout should contain success message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "favorite", "add", "-n", "test1"})
	assert.Nil(t, err, "error should be nil")
	assert.Empty(t, stdout.String(), "stdout should be empty")
	assert.Contains(t, stderr.String(), "Snippet not found", "stderr should contain error message")
}

func TestFavoriteDeleteCommand(t *testing.T) {
	state, app, stdout, stderr := SetupTestApp()

	err := app.Run([]string{"snippy", "favorite", "delete", "-h"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "snippy favorite delete", "stdout should contain help message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "favorite", "delete"})
	assert.NotNil(t, err, "error should not be nil")
	assert.Empty(t, stdout.String(), "stdout should be empty")
	assert.Contains(t, stderr.String(), "Name is required", "stderr should contain error message")

	stdout.Reset()
	stderr.Reset()

	state.localSnippets["test"] = &Snippet{Content: "test content", Favorite: true}

	err = app.Run([]string{"snippy", "favorite", "delete", "--name", "test"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "Favorite deleted successfully", "stdout should contain success message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "favorite", "delete", "-n", "test"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "Favorite deleted successfully", "stdout should contain success message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	state.localSnippets["test"] = &Snippet{Content: "test content", Favorite: false}

	err = app.Run([]string{"snippy", "favorite", "delete", "-n", "test"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "Favorite deleted successfully", "stdout should contain success message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "favorite", "delete", "-n", "test1"})
	assert.Nil(t, err, "error should be nil")
	assert.Empty(t, stdout.String(), "stdout should be empty")
	assert.Contains(t, stderr.String(), "Snippet not found", "stderr should contain error message")
}

func TestInitCommand(t *testing.T) {
	initialDir, _ := os.Getwd()
	_, app, stdout, stderr := SetupTestApp()

	dir := t.TempDir()
	os.Chdir(dir)
	defer os.Chdir(initialDir)

	err := app.Run([]string{"snippy", "init", "-h"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "snippy init", "stdout should contain help message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "init"})
	assert.Nil(t, err, "error should be nil")
	assert.Contains(t, stdout.String(), "Snippy initialized successfully", "stdout should contain success message")
	assert.Empty(t, stderr.String(), "stderr should be empty")

	stdout.Reset()
	stderr.Reset()

	err = app.Run([]string{"snippy", "init"})
	assert.NotNil(t, err, "error should not be nil")
	assert.Empty(t, stdout.String(), "stdout should be empty")
	assert.Contains(t, stderr.String(), "Snippy already initialized", "stderr should contain error message")

}
