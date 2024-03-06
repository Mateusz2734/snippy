package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewState(t *testing.T) {
	state := NewState()
	assert.NotNil(t, state)

	assert.Nil(t, state.localSnippets, "localSnippets should be nil")
	assert.Nil(t, state.globalSnippets, "globalSnippets should be nil")

	assert.Empty(t, state.InputFile, "InputFile should be empty")
	assert.Empty(t, state.Name, "Name should be empty")
	assert.Empty(t, state.Extension, "Extension should be empty")

	assert.Zero(t, state.CurrentPage, "CurrentPage should be 0")
	assert.Zero(t, state.PageSize, "PageSize should be 0")

	assert.False(t, state.UseClipboard, "UseClipboard should be false")
	assert.False(t, state.NoMetadata, "NoMetadata should be false")
	assert.False(t, state.UseGlobal, "UseGlobal should be false")
	assert.False(t, state.NoFormatting, "NoFormatting should be false")
}

func TestGetSnippets(t *testing.T) {
	state := NewState()

	state.UseGlobal = true
	state.globalSnippets = make(map[string]*Snippet)
	state.localSnippets = nil

	assert.Equal(t, state.globalSnippets, state.GetSnippets(), "Should return globalSnippets")

	state.UseGlobal = true
	state.globalSnippets = nil
	state.localSnippets = make(map[string]*Snippet)

	assert.Equal(t, state.globalSnippets, state.GetSnippets(), "Should return globalSnippets")

	state.UseGlobal = false
	state.globalSnippets = map[string]*Snippet{"test": &Snippet{Content: "test"}}
	state.localSnippets = make(map[string]*Snippet)

	assert.Equal(t, state.localSnippets, state.GetSnippets(), "Should return localSnippets")

	state.UseGlobal = false
	state.globalSnippets = make(map[string]*Snippet)
	state.localSnippets = nil

	assert.Equal(t, state.globalSnippets, state.GetSnippets(), "Should return globalSnippets")

	state.UseGlobal = false
	state.globalSnippets = nil
	state.localSnippets = make(map[string]*Snippet)

	assert.Equal(t, state.localSnippets, state.GetSnippets(), "Should return localSnippets")
}
