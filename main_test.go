package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestLocalUserFlow(t *testing.T) {
	initialDir, _ := os.Getwd()
	state := NewState()
	buff := bytes.NewBuffer(nil)

	app := NewApp(state, []cli.Flag{WithGlobal(state), WithNoFormatting(state)})
	app.Writer = buff
	app.ErrWriter = buff

	dir := t.TempDir()
	os.Chdir(dir)
	defer os.Chdir(initialDir)

	app.Run([]string{"snippy", "init"})
	assert.FileExists(t, filepath.Join(dir, ".snippy"), "snippy init should create a .snippy file")

	app.Run([]string{"snippy", "add", "--name", "test", "--content", "echo test"})
	assert.Contains(t, state.GetSnippets(), "test", "snippy add should add a snippet to the state")

	app.Run([]string{"snippy", "add", "-n", "newTest", "--content", "echo new test", "-ext", "sh"})
	assert.Contains(t, state.GetSnippets(), "newTest", "snippy add should add a snippet to the state")
	snippet := state.GetSnippets()["newTest"]
	assert.Equal(t, "sh", snippet.Extension, "snippy add should add a snippet with the correct extension")
	assert.Equal(t, "echo new test", snippet.Content, "snippy add should add a snippet with the correct content")

	buff.Reset()
	app.Run([]string{"snippy", "list"})
	assert.Contains(t, buff.String(), "test", "snippy list should list the snippet")
	assert.Contains(t, buff.String(), "newTest", "snippy list should list the snippet")

	app.Run([]string{"snippy", "edit", "-n", "newTest", "--ext", "go"})
	assert.NotNil(t, state.GetSnippets(), "snippy edit should not remove the snippet from the state")
	assert.Contains(t, state.GetSnippets(), "newTest", "snippy edit should not remove the snippet from the state")
	snippet = state.GetSnippets()["newTest"]
	assert.Equal(t, "go", snippet.Extension, "snippy edit should edit the snippet with the correct extension")

	app.Run([]string{"snippy", "delete", "-n", "test"})
	assert.NotContains(t, state.GetSnippets(), "test", "snippy delete should remove the snippet from the state")

	app.Run([]string{"snippy", "favorite", "add", "-n", "newTest"})
	assert.True(t, state.GetSnippets()["newTest"].Favorite, "snippy favorite add should make the snippet favorite")

	buff.Reset()
	app.Run([]string{"snippy", "favorite", "list"})
	assert.Contains(t, buff.String(), "newTest", "snippy favorite list should list the snippet")

	app.Run([]string{"snippy", "favorite", "delete", "-n", "newTest"})
	assert.False(t, state.GetSnippets()["newTest"].Favorite, "snippy favorite remove should remove the snippet from favorite")

	buff.Reset()
	app.Run([]string{"snippy", "search", "--ext", "go"})
	assert.Contains(t, buff.String(), "newTest", "snippy search should list the snippet with the correct extension")

	buff.Reset()
	app.Run([]string{"snippy", "get", "-n", "newTest", "--nf"})
	assert.Contains(t, buff.String(), "echo new test", "snippy get should print the snippet content")
	assert.Contains(t, buff.String(), "Metadata = {", "snippy get should print the snippet metadata")
}
