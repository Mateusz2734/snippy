package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func ReadSnippets() map[string]string {
	var snippets map[string]string

	homeDir, _ := os.UserHomeDir()
	path := filepath.Join(homeDir, ".snippy")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Create(path)
	}

	data, err := os.ReadFile(path)

	if err != nil {
		cli.Exit("Cannot read snippets", 1)
	}

	err = json.Unmarshal(data, &snippets)

	if err != nil {
		cli.Exit("Cannot parse snippets", 1)
	}

	return snippets
}

func WriteSnippets(data map[string]string) {
	homeDir, _ := os.UserHomeDir()
	path := filepath.Join(homeDir, ".snippy")

	jsonData, err := json.Marshal(data)

	if err != nil {
		cli.Exit("Cannot parse snippets", 1)
	}

	err = os.WriteFile(path, jsonData, 0644)

	if err != nil {
		cli.Exit("Cannot write snippets", 1)
	}
}
