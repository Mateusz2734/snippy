package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func DefaultAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		fmt.Println(state.Snippets[cCtx.Args().First()])
		fmt.Println(cCtx.FlagNames(), state.InFile)
		return nil
	}
}

func GetAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		fmt.Println("get snippet")
		return nil
	}
}

func ListAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		fmt.Println("list snippets")
		return nil
	}
}

func AddAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		fmt.Println("edit snippet")
		return nil
	}
}

func EditAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		fmt.Println("edit snippet")
		return nil
	}
}

func DeleteAction(state *State) func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		fmt.Println("delete snippet")
		return nil
	}
}
