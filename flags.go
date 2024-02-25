package main

import "github.com/urfave/cli/v2"

func WithInputFile(state *State) cli.Flag {
	return &cli.StringFlag{
		Name:        "file",
		Usage:       "Set the content from `FILE` as the snippet.",
		Aliases:     []string{"f"},
		Destination: &state.InputFile,
	}
}

func WithName(state *State) cli.Flag {
	return &cli.StringFlag{
		Name:        "name",
		Usage:       "Set the name of the snippet.",
		Aliases:     []string{"n"},
		Destination: &state.Name,
	}
}

func WithLanguage(state *State) cli.Flag {
	return &cli.StringFlag{
		Name:        "language",
		Usage:       "Set the language of the snippet.",
		Aliases:     []string{"l", "lang"},
		Destination: &state.Language,
	}
}
