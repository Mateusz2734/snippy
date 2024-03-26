package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli/v2"
)

func ReadSnippets() (map[string]*Snippet, map[string]*Snippet) {
	var globalSnippets = make(map[string]*Snippet)

	homeDir, _ := os.UserHomeDir()
	path := filepath.Join(homeDir, ".snippy")

	globalSnippets = readSnippetsFromFile(path)

	localExists, localPath := localSnippyExists()
	if localExists {
		localSnippets := readSnippetsFromFile(localPath)
		return globalSnippets, localSnippets
	}

	return globalSnippets, nil
}

func WriteSnippets(data map[string]*Snippet, global bool) {
	path := ""

	if global {
		homeDir, _ := os.UserHomeDir()
		path = filepath.Join(homeDir, ".snippy")
	} else {
		localExists, localPath := localSnippyExists()
		if localExists {
			path = localPath
		} else {
			cli.Exit("Local snippy not found", 1)
		}
	}

	writeSnippetsToFile(data, path)
}

func CreateBackupFile(data map[string]*Snippet, dir string) error {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return fmt.Errorf("Cannot parse snippets")
	}

	stat, err := os.Stat(dir)

	if err != nil && os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	stat, err = os.Stat(dir)

	if err != nil {
		return fmt.Errorf("Cannot stat backup directory")
	}

	if !stat.IsDir() {
		return fmt.Errorf("Backup directory is a file")
	}

	backupFile := filepath.Join(dir, fmt.Sprintf(".snippy_backup_%d", time.Now().Unix()))

	err = os.WriteFile(backupFile, jsonData, 0644)

	if err != nil {
		return fmt.Errorf("Cannot write snippets")
	}

	return nil
}

func writeSnippetsToFile(data map[string]*Snippet, path string) {
	jsonData, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		cli.Exit("Cannot parse snippets", 1)
	}

	err = os.WriteFile(path, jsonData, 0644)

	if err != nil {
		cli.Exit("Cannot write snippets", 1)
	}
}

func localSnippyExists() (bool, string) {
	path, _ := os.Getwd()
	homeDir, _ := os.UserHomeDir()

	for path != filepath.Dir(path) {
		if path == homeDir {
			return false, ""
		}

		snippyPath := filepath.Join(path, ".snippy")
		if stat, err := os.Stat(snippyPath); err == nil && !stat.IsDir() {
			return true, snippyPath
		}

		path = filepath.Dir(path)
	}

	return false, ""
}

func readSnippetsFromFile(path string) map[string]*Snippet {
	snippets := make(map[string]*Snippet)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Create(path)
	}

	data, err := os.ReadFile(path)

	if err != nil {
		cli.Exit("Cannot read snippets", 1)
	}

	if len(data) == 0 {
		return snippets
	}

	err = json.Unmarshal(data, &snippets)

	if err != nil {
		cli.Exit("Cannot parse snippets", 1)
	}

	return snippets
}
