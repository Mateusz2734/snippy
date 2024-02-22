package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	state := &State{}

	app := &cli.App{
		Before: func(cCtx *cli.Context) error {
			state.Snippets = ReadSnippets()
			return nil
		},
		Action: DefaultAction(state),
		Commands: []*cli.Command{
			AddCommand(state),
			ListCommand(state),
			EditCommand(state),
			GetCommand(state),
			DeleteCommand(state),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
