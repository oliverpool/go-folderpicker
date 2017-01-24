package folderpicker

import "testing"
import "fmt"

func TestExample(t *testing.T) {
	folder, err := Prompt("Select a folder")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Selected folder:", folder)
}
