package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	state := NewState()

	app := NewApp(state, []cli.Flag{WithGlobal(state), WithNoFormatting(state)})

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
