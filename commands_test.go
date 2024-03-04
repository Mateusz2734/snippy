package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func setupTestApp() (*State, *cli.App, *bytes.Buffer, *bytes.Buffer) {
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

func TestListCommand(t *testing.T) {
	state, app, stdout, stderr := setupTestApp()

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

func TestFavoriteCommand(t *testing.T) {
	_, app, stdout, stderr := setupTestApp()

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
	state, app, stdout, stderr := setupTestApp()

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
	state, app, stdout, stderr := setupTestApp()

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
	state, app, stdout, stderr := setupTestApp()

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
	_, app, stdout, stderr := setupTestApp()

	dir := t.TempDir()
	os.Chdir(dir)

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

	os.Chdir(initialDir)
}
