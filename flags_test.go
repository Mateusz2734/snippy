package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestWithClipboard(t *testing.T) {
	state := &State{}
	buf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)

	app := cli.App{
		Flags:     []cli.Flag{WithClipboard(state)},
		Writer:    buf,
		ErrWriter: errBuf,
	}

	t.Setenv("SNIPPY_USE_CLIPBOARD", "")

	app.Run([]string{"app"})
	assert.False(t, state.UseClipboard, "default clipboard value should be false")

	app.Run([]string{"app", "--clipboard"})
	assert.True(t, state.UseClipboard, "clipboard value should be true")

	app.Run([]string{"app", "-c"})
	assert.True(t, state.UseClipboard, "clipboard value should be true")

	t.Setenv("SNIPPY_USE_CLIPBOARD", "true")

	app.Run([]string{"app"})
	assert.True(t, state.UseClipboard, "clipboard value should be true")

	t.Setenv("SNIPPY_USE_CLIPBOARD", "false")

	app.Run([]string{"app"})
	assert.False(t, state.UseClipboard, "clipboard value should be false")

	t.Setenv("SNIPPY_USE_CLIPBOARD", "invalid")

	err := app.Run([]string{"app"})
	assert.NotNil(t, err, "error expected")

	t.Setenv("SNIPPY_USE_CLIPBOARD", "false")

	app.Run([]string{"app", "--clipboard"})
	assert.True(t, state.UseClipboard, "clipboard value should be true")
}

func TestWithPage(t *testing.T) {
	state := &State{}
	buf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)

	app := cli.App{
		Flags:     []cli.Flag{WithPage(state)},
		Writer:    buf,
		ErrWriter: errBuf,
	}

	err := app.Run([]string{"app"})
	assert.Equal(t, 1, state.CurrentPage, "default page should be 1")
	assert.Nil(t, err, "no error expected")

	err = app.Run([]string{"app", "--page", "2"})
	assert.Equal(t, 2, state.CurrentPage, "page should be 2")
	assert.Nil(t, err, "no error expected")

	err = app.Run([]string{"app", "-p", "3"})
	assert.Equal(t, 3, state.CurrentPage, "page should be 3")
	assert.Nil(t, err, "no error expected")

	err = app.Run([]string{"app", "-p", "a"})
	assert.NotNil(t, err, "error expected")

	err = app.Run([]string{"app", "-p", "0"})
	assert.NotNil(t, err, "error expected")
}

func TestWithPageSize(t *testing.T) {
	state := &State{}
	buf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)

	app := cli.App{
		Flags:     []cli.Flag{WithPageSize(state)},
		Writer:    buf,
		ErrWriter: errBuf,
	}

	t.Setenv("SNIPPY_PAGE_SIZE", "")
	app.Run([]string{"app"})

	assert.Equal(t, 20, state.PageSize, "default page size should be 20")
	assert.Empty(t, errBuf.String(), "no error expected")

	errBuf.Reset()
	buf.Reset()

	t.Setenv("SNIPPY_PAGE_SIZE", "30")
	app.Run([]string{"app"})

	assert.Equal(t, 30, state.PageSize, "page size should be 30")
	assert.Empty(t, errBuf.String(), "no error expected")
}

func TestMetadata(t *testing.T) {
	state := &State{}
	buf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)

	app := cli.App{
		Writer:    buf,
		ErrWriter: errBuf,
		Flags: []cli.Flag{
			WithMetadata(state),
			WithNoMetadata(state),
		},
	}

	t.Setenv("SNIPPY_NO_METADATA", "")

	app.Run([]string{"app"})
	assert.False(t, state.NoMetadata, "default noMetadata value should be false")

	app.Run([]string{"app", "--no-metadata"})
	assert.True(t, state.NoMetadata, "noMetadata value should be true")

	app.Run([]string{"app", "--metadata"})
	assert.False(t, state.NoMetadata, "noMetadata value should be false")

	app.Run([]string{"app", "--metadata", "--no-metadata"})
	assert.False(t, state.NoMetadata, "noMetadata value should be true")

	app.Run([]string{"app", "--no-metadata", "--metadata"})
	assert.False(t, state.NoMetadata, "noMetadata value should be false")

	t.Setenv("SNIPPY_NO_METADATA", "true")

	app.Run([]string{"app"})
	assert.True(t, state.NoMetadata, "noMetadata value should be true")

	app.Run([]string{"app", "--metadata"})
	assert.False(t, state.NoMetadata, "noMetadata value should be false")

	t.Setenv("SNIPPY_NO_METADATA", "false")

	app.Run([]string{"app"})
	assert.False(t, state.NoMetadata, "noMetadata value should be false")

	t.Setenv("SNIPPY_NO_METADATA", "invalid")

	err := app.Run([]string{"app"})
	assert.NotNil(t, err, "error expected")
}
