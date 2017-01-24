package folderpicker

import (
	"fmt"
	"runtime"
	"testing"
)

func TestExample(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Linux is not implemented yet")
	}
	folder, err := Prompt("Select a folder")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Selected folder:", folder)
}
