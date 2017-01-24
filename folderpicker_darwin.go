package folderpicker

import "os/exec"

func pickFolder(msg string) *exec.Cmd {
	return exec.Command("osascript",
		"-e", "tell application \"Finder\"",
		"-e", "activate",
		"-e", "set myfolder to choose folder with prompt \""+msg+"\"",
		"-e", "end tell",
		"-e", "return (posix path of myfolder)",
	)
}
