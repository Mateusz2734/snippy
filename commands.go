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
			WithExtension(state),
			WithClipboard(state),
			WithContent(state),
		},
	}
}

func ListCommand(state *State) *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "list snippets",
		Action:  ListAction(state),
		Flags: []cli.Flag{
			WithPage(state),
			WithPageSize(state),
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
			WithMetadata(state),
			WithNoMetadata(state),
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

func EditCommand(state *State) *cli.Command {
	return &cli.Command{
		Name:    "edit",
		Aliases: []string{"e"},
		Usage:   "edit snippet",
		After:   saveFunc(state),
		Action:  EditAction(state),
		Flags: []cli.Flag{
			WithName(state),
			WithExtension(state),
		},
	}
}

func SearchCommand(state *State) *cli.Command {
	return &cli.Command{
		Name:    "search",
		Aliases: []string{"s"},
		Usage:   "search snippets",
		Action:  SearchAction(state),
		Flags: []cli.Flag{
			WithExtension(state),
			WithName(state),
		},
	}
}

func FavoriteCommand(state *State) *cli.Command {
	return &cli.Command{
		Name:            "favorite",
		Aliases:         []string{"f"},
		Usage:           "manage favorite snippets",
		HideHelpCommand: true,
		Subcommands: []*cli.Command{
			FavoriteAddCommand(state),
			FavoriteDeleteCommand(state),
			FavoriteListCommand(state),
		},
	}
}

func FavoriteAddCommand(state *State) *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add snippet to favorites",
		After:   saveFunc(state),
		Action:  FavoriteAddAction(state),
		Flags:   []cli.Flag{WithName(state)},
	}
}

func FavoriteDeleteCommand(state *State) *cli.Command {
	return &cli.Command{
		Name:    "delete",
		Aliases: []string{"d"},
		Usage:   "delete snippet from favorites",
		After:   saveFunc(state),
		Action:  FavoriteDeleteAction(state),
		Flags:   []cli.Flag{WithName(state)},
	}
}

func FavoriteListCommand(state *State) *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "list favorites",
		Action:  FavoriteListAction(state),
		Flags: []cli.Flag{
			WithPage(state),
			WithPageSize(state),
		},
	}
}

func InitCommand(state *State) *cli.Command {
	return &cli.Command{
		Name:   "init",
		Usage:  "initialize snippy",
		Action: InitAction(state),
	}
}

func saveFunc(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		WriteSnippets(state.GetSnippets(), state.UseGlobal || state.localSnippets == nil)
		return nil
	}
}
