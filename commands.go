package main

import (
	"github.com/urfave/cli/v2"
)

func AddCommand(state *State) *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add snippet",
		After:   saveFunc(state),
		Action:  AddAction(state),
		Flags: []cli.Flag{
			WithInputFile(state),
			WithName(state),
		},
	}
}

func ListCommand(state *State) *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "list snippets",
		Action:  ListAction(state),
	}
}

func EditCommand(state *State) *cli.Command {
	return &cli.Command{
		Name:    "edit",
		Aliases: []string{"e"},
		Usage:   "edit snippet",
		After:   saveFunc(state),
		Action:  EditAction(state),
		Flags: []cli.Flag{
			WithName(state),
		},
	}
}

func GetCommand(state *State) *cli.Command {
	return &cli.Command{
		Name:   "get",
		Usage:  "get snippet",
		Action: GetAction(state),
		Flags: []cli.Flag{
			WithName(state),
		},
	}
}

func DeleteCommand(state *State) *cli.Command {
	return &cli.Command{
		Name:    "delete",
		Aliases: []string{"d"},
		Usage:   "delete snippet",
		After:   saveFunc(state),
		Action:  DeleteAction(state),
		Flags: []cli.Flag{
			WithName(state),
		},
	}
}

func saveFunc(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		WriteSnippets(state.Snippets)
		return nil
	}
}
