package main

type State struct {
	Snippets  map[string]*Snippet
	InputFile string
	Name      string
	Extension string
}
