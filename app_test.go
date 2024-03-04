package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestNewApp(t *testing.T) {
	app := NewApp(&State{}, []cli.Flag{})

	assert.NotNil(t, app, "app should not be nil")

	assert.Equal(t, "snippy", app.Name, "app name should be snippy")

	assert.NotNil(t, app.Before, "app BeforeFunc should not be nil")

	assert.NotNil(t, app.Action, "app ActionFunc should not be nil")

	assert.Len(t, app.Commands, 8, "app should have 8 commands")

	assert.True(t, app.HideHelpCommand, "app should hide help command")
}

func TestApplyGlobalFlagsSimple(t *testing.T) {
	cmd := &cli.Command{
		Subcommands: []*cli.Command{
			{Name: "subCmd1"},
			{Name: "subCmd2"},
		},
	}

	applyGlobalFlags([]cli.Flag{&cli.StringFlag{Name: "test"}}, cmd)

	assert.Len(t, cmd.Flags, 1, "command should have 1 flag")

	for _, subcommand := range cmd.Subcommands {
		assert.Len(t, subcommand.Flags, 1, "subcommand should have 1 flag")
	}
}

func TestApplyGlobalFlagsComplex(t *testing.T) {
	// 1 layer deep
	cmd1 := &cli.Command{Name: "cmd1"}

	// 2 layers deep
	cmd2 := &cli.Command{
		Name:        "cmd2",
		Subcommands: []*cli.Command{{Name: "subCmd2"}},
	}

	// 3 layers deep
	cmd3 := &cli.Command{
		Name: "cmd3",
		Subcommands: []*cli.Command{
			{
				Name: "subCmd3",
				Subcommands: []*cli.Command{
					{Name: "subSubCmd3"},
				},
			},
		},
	}

	applyGlobalFlags([]cli.Flag{&cli.StringFlag{Name: "test"}}, cmd1)
	assert.Len(t, cmd1.Flags, 1, "command should have 1 flag")

	applyGlobalFlags([]cli.Flag{&cli.StringFlag{Name: "test"}}, cmd2)
	assert.Len(t, cmd2.Flags, 1, "command should have 1 flag")
	for _, subcommand := range cmd2.Subcommands {
		assert.Len(t, subcommand.Flags, 1, "subcommand should have 1 flag")
	}

	applyGlobalFlags([]cli.Flag{&cli.StringFlag{Name: "test"}}, cmd3)
	assert.Len(t, cmd3.Flags, 1, "command should have 1 flag")
	for _, subcommand := range cmd3.Subcommands {
		assert.Len(t, subcommand.Flags, 1, "subcommand should have 1 flag")
		for _, subsubcommand := range subcommand.Subcommands {
			assert.Len(t, subsubcommand.Flags, 0, "subsubcommand should not have any flags")
		}
	}
}
