package folderpicker

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

func TestExample(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Linux is not implemented yet")
	}
	//return
	folder, err := Prompt("Select a folder")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Selected folder:", folder, filepath.ToSlash(folder))
}
