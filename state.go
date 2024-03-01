package main

type State struct {
	// Snippet data
	localSnippets  map[string]*Snippet
	globalSnippets map[string]*Snippet

	// Flag data
	InputFile    string
	Name         string
	Extension    string
	UseClipboard bool
	CurrentPage  int
	PageSize     int
	NoMetadata   bool
	UseGlobal    bool
	NoFormatting bool
}

func (state *State) GetSnippets() map[string]*Snippet {
	if !state.UseGlobal && state.localSnippets != nil {
		return state.localSnippets
	}
	return state.globalSnippets
}

func (state *State) InitializeSnippets(globalSnippets map[string]*Snippet, localSnippets map[string]*Snippet) {
	state.globalSnippets = globalSnippets
	state.localSnippets = localSnippets
}

func NewState() *State {
	state := &State{}

	return state
}
