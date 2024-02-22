package main

import (
	"log"
	"os"
)

func main() {
	state := &State{}

	app := NewApp(state)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
