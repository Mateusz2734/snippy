package main

import (
	"github.com/urfave/cli/v2"
)

func NewApp(state *State, globalFlags []cli.Flag) *cli.App {
	app := &cli.App{
		Name:            "snippy",
		Usage:           "CLI snippet manager",
		HideHelpCommand: true,
		Before: func(cCtx *cli.Context) error {
			state.InitializeSnippets(ReadSnippets())
			return nil
		},
		Action: DefaultAction(state),
		Commands: []*cli.Command{
			AddCommand(state),
			GetCommand(state),
			ListCommand(state),
			DeleteCommand(state),
			EditCommand(state),
			SearchCommand(state),
			FavoriteCommand(state),
			InitCommand(state),
		},
	}

	for _, command := range app.Commands {
		applyGlobalFlags(globalFlags, command)
	}

	return app
}

// It applies the globalFlags to a command and its direct subcommands.
// IMPORTANT: It does NOT apply flags recursively.
func applyGlobalFlags(globalFlags []cli.Flag, command *cli.Command) {
	if command.Subcommands != nil {
		for _, subcommand := range command.Subcommands {
			subcommand.Flags = append(subcommand.Flags, globalFlags...)
		}
	}
	command.Flags = append(command.Flags, globalFlags...)
}
