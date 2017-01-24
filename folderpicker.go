package folderpicker

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// ErrEmptyFolderSelected indicates that the selected path is empty
var ErrEmptyFolderSelected = errors.New("No folder selected")

// PromptCmd returns a command to prompt the user to pick a folder
func PromptCmd(msg string) *exec.Cmd {
	return pickFolder(msg)
}

// Prompt let the user pick a folder and returns a clean result
func Prompt(msg string) (folder string, err error) {
	cmd := PromptCmd(msg)
	out, err := cmd.CombinedOutput()
	folder = cleanFolder(out)
	if err == nil && folder == "" {
		err = ErrEmptyFolderSelected
	}
	return
}

func cleanFolder(b []byte) string {
	fmt.Println("raw")
	fmt.Println(string(b))
	s := string(b)
	s = strings.TrimSpace(s)
	s = filepath.Clean(s)
	fmt.Println("cleaned")
	fmt.Println(s)
	if s == "." || s == `\` {
		return ""
	}
	return s
}
