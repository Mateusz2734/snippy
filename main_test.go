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

func TestGlobalUserFlow(t *testing.T) {
	homeDir, _ := os.UserHomeDir()
	snippyPath := filepath.Join(homeDir, ".snippy")

	content, err := os.ReadFile(snippyPath)

	if err != nil && !os.IsNotExist(err) {
		t.Fatal(err)
	}

	os.Remove(snippyPath)

	t.Cleanup(func() {
		os.WriteFile(snippyPath, content, 0644)
	})

	state := NewState()
	buff := bytes.NewBuffer(nil)

	app := NewApp(state, []cli.Flag{WithGlobal(state), WithNoFormatting(state)})
	app.Writer = buff
	app.ErrWriter = buff

	app.Run([]string{"snippy", "add", "-n", "test", "--content", "echo test"})
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

func TestMixedUserFlow(t *testing.T) {
	initialDir, _ := os.Getwd()
	homeDir, _ := os.UserHomeDir()
	snippyPath := filepath.Join(homeDir, ".snippy")
	tempDir := t.TempDir()

	content, err := os.ReadFile(snippyPath)

	if err != nil && !os.IsNotExist(err) {
		t.Fatal(err)
	}

	os.Chdir(tempDir)
	os.Remove(snippyPath)

	t.Cleanup(func() {
		os.WriteFile(snippyPath, content, 0644)
		os.Chdir(initialDir)
	})

	state := NewState()
	buff := bytes.NewBuffer(nil)

	app := NewApp(state, []cli.Flag{WithGlobal(state), WithNoFormatting(state)})
	app.Writer = buff
	app.ErrWriter = buff

	app.Run([]string{"snippy", "add", "-n", "global", "--content", "echo global"})
	assert.Contains(t, state.GetSnippets(), "global", "snippy add should add a snippet to the global state")

	app.Run([]string{"snippy", "init"})
	assert.FileExists(t, filepath.Join(tempDir, ".snippy"), "snippy init should create a local .snippy file")

	app.Run([]string{"snippy", "add", "-n", "local", "--content", "echo local"})
	assert.Contains(t, state.GetSnippets(), "local", "snippy add should add a snippet to the local state")

	buff.Reset()
	app.Run([]string{"snippy", "list"})
	assert.Contains(t, buff.String(), "local", "snippy list should list the local snippet")

	buff.Reset()
	app.Run([]string{"snippy", "list", "--global"})
	assert.Contains(t, buff.String(), "global", "snippy list --global should list the global snippet")
}
