package folderpicker

import (
	"errors"
	"path/filepath"
	"strings"
)

// ErrEmptyFolderSelected indicates that the selected path is empty
var ErrEmptyFolderSelected = errors.New("No folder selected")

// Prompt let the user pick a folder and returns a clean result
func Prompt(msg string) (folder string, err error) {
	folder, err = pickFolder(msg)
	folder = cleanFolder(folder)
	if err == nil && folder == "" {
		err = ErrEmptyFolderSelected
	}
	return
}

func cleanFolder(s string) string {
	s = strings.TrimSpace(s)
	s = filepath.Clean(s)
	if s == "." || s == `\` {
		return ""
	}
	return s
}

type Prompter interface {
	Prompt(msg string) (string, error)
}
