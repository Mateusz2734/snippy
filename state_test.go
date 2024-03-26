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
	assert.False(t, state.Global, "UseGlobal should be false")
	assert.False(t, state.NoFormatting, "NoFormatting should be false")
}

func TestUseGlobal(t *testing.T) {
	state := NewState()

	state.Global = true
	state.localSnippets = nil
	assert.True(t, state.UseGlobal(), "Should return true")

	state.Global = true
	state.localSnippets = make(map[string]*Snippet)
	assert.True(t, state.UseGlobal(), "Should return true")

	state.Global = false
	state.localSnippets = nil
	assert.True(t, state.UseGlobal(), "Should return true")

	state.Global = false
	state.localSnippets = make(map[string]*Snippet)
	assert.False(t, state.UseGlobal(), "Should return false")
}

func TestGetSnippets(t *testing.T) {
	state := NewState()

	state.Global = true
	state.globalSnippets = make(map[string]*Snippet)
	state.localSnippets = nil

	assert.Equal(t, state.globalSnippets, state.GetSnippets(), "Should return globalSnippets")

	state.Global = false
	state.globalSnippets = nil
	state.localSnippets = make(map[string]*Snippet)

	assert.Equal(t, state.localSnippets, state.GetSnippets(), "Should return localSnippets")
}

func TestSetSnippets(t *testing.T) {
	state := NewState()

	globalSnippets := make(map[string]*Snippet)
	localSnippets := make(map[string]*Snippet)
	globalSnippets["global"] = &Snippet{Content: "global"}
	localSnippets["local"] = &Snippet{Content: "local"}

	state.Global = true
	state.SetSnippets(globalSnippets)
	assert.Equal(t, globalSnippets, state.globalSnippets, "Should set globalSnippets")
	_, ok := state.globalSnippets["global"]
	assert.True(t, ok, "Should contain 'global' snippet")

	state.Global = false
	state.localSnippets = make(map[string]*Snippet)
	state.SetSnippets(localSnippets)
	assert.Equal(t, localSnippets, state.localSnippets, "Should set localSnippets")
	_, ok = state.localSnippets["local"]
	assert.True(t, ok, "Should contain 'local' snippet")
}
