package main

type State struct {
	Snippets     map[string]*Snippet
	InputFile    string
	Name         string
	Extension    string
	UseClipboard bool
	CurrentPage  int
	PageSize     int
	NoMetadata   bool
}
