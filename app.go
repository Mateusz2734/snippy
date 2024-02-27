package main

import (
	"github.com/urfave/cli/v2"
)

func NewApp(state *State) *cli.App {
	return &cli.App{
		Name:            "snippy",
		Usage:           "CLI snippet manager",
		HideHelpCommand: true,
		Before: func(cCtx *cli.Context) error {
			state.Snippets = ReadSnippets()
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
		},
	}
}
