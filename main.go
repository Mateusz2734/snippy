package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	myMap := map[string]string{
		"aa": "aaa",
		"bb": "bbb",
		"cc": "ccc",
	}

	app := &cli.App{

		Action: func(cCtx *cli.Context) error {
			fmt.Println(myMap[cCtx.Args().First()])
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "browse",
				Usage: "browse snippets",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("browse snippets")
					return nil
				},
			},
			{
				Name:  "add",
				Usage: "add snippet",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("add snippet")
					return nil
				},
			},
			{
				Name:  "edit",
				Usage: "edit snippet",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("edit snippet")
					return nil
				},
			},
			{
				Name:  "get",
				Usage: "get snippet",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("get snippet")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
