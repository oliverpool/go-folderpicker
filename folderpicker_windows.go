package folderpicker

import "os/exec"

func pickFolder(msg string) *exec.Cmd {
	return exec.Command("powershell",
		"-windowstyle",
		"hidden",
		"(new-object -COM 'Shell.Application').BrowseForFolder(0,'"+msg+"',0x1+0x100,0).self.path+'\\'",
	)
}
