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

	return state, app, stdout, stderr
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
