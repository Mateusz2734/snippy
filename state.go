package main

type State struct {
	// Snippet data
	localSnippets  map[string]*Snippet
	globalSnippets map[string]*Snippet

	// Flag data
	InputFile    string
	Name         string
	Extension    string
	Content      string
	Directory    string
	UseClipboard bool
	CurrentPage  int
	PageSize     int
	NoMetadata   bool
	Global       bool
	NoFormatting bool
}

func (state *State) UseGlobal() bool {
	return state.Global || state.localSnippets == nil
}

func (state *State) GetSnippets() map[string]*Snippet {
	if state.UseGlobal() {
		return state.globalSnippets
	}
	return state.localSnippets
}

func (state *State) SetSnippets(snippets map[string]*Snippet) {
	if state.UseGlobal() {
		state.globalSnippets = snippets
	} else {
		state.localSnippets = snippets
	}
}

func (state *State) InitializeSnippets(globalSnippets map[string]*Snippet, localSnippets map[string]*Snippet) {
	state.globalSnippets = globalSnippets
	state.localSnippets = localSnippets
}

func NewState() *State {
	state := &State{}

	return state
}
