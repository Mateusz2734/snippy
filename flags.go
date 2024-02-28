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

func WithExtension(state *State) cli.Flag {
	return &cli.StringFlag{
		Name:        "extension",
		Usage:       "Set the programming language extension of the snippet.",
		Aliases:     []string{"e", "ext"},
		Destination: &state.Extension,
	}
}

func WithClipboard(state *State) cli.Flag {
	return &cli.BoolFlag{
		Name:        "clipboard",
		Usage:       "Import content from the clipboard.",
		Aliases:     []string{"c", "clip"},
		Destination: &state.UseClipboard,
		EnvVars:     []string{"SNIPPY_USE_CLIPBOARD"},
	}
}
