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

func WithPage(state *State) cli.Flag {
	return &cli.IntFlag{
		Name:        "page",
		Usage:       "Specify current page for the snippet list.",
		Aliases:     []string{"p"},
		Value:       1,
		Destination: &state.CurrentPage,
		Action: func(cCtx *cli.Context, newValue int) error {
			if newValue < 1 {
				return cli.Exit("Page number should be greater than 0", 1)
			}
			return nil
		},
	}
}

func WithPageSize(state *State) cli.Flag {
	return &cli.IntFlag{
		Hidden:      true,
		Value:       20,
		Destination: &state.PageSize,
		EnvVars:     []string{"SNIPPY_PAGE_SIZE"},
	}
}

func WithNoMetadata(state *State) cli.Flag {
	return &cli.BoolFlag{
		Name:        "no-metadata",
		Usage:       "Do not print metadata about the snippet.",
		Destination: &state.NoMetadata,
		EnvVars:     []string{"SNIPPY_NO_METADATA"},
		Action: func(cCtx *cli.Context, val bool) error {
			state.NoMetadata = val
			return nil
		},
	}
}

func WithMetadata(state *State) cli.Flag {
	return &cli.BoolFlag{
		Name:        "metadata",
		Usage:       "Print metadata about the snippet.",
		Destination: &state.NoMetadata,
		Action: func(cCtx *cli.Context, val bool) error {
			state.NoMetadata = !val
			return nil
		},
	}
}

func WithGlobal(state *State) cli.Flag {
	return &cli.BoolFlag{
		Name:        "global",
		Usage:       "Use global snippets.",
		Aliases:     []string{"g"},
		Destination: &state.UseGlobal,
	}
}

func WithNoFormatting(state *State) cli.Flag {
	return &cli.BoolFlag{
		Name:        "no-formatting",
		Usage:       "Do not format the snippet content.",
		Aliases:     []string{"nf"},
		Destination: &state.NoFormatting,
		EnvVars:     []string{"SNIPPY_NO_FORMATTING"},
	}
}
