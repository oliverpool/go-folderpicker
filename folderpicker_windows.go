package folderpicker

import (
	"errors"
	"fmt"
	"os/exec"
	"syscall"

	"github.com/lxn/win"
)

func oldpickFolder(msg string) *exec.Cmd {
	return exec.Command("powershell",
		"-windowstyle",
		"hidden",
		"(new-object -COM 'Shell.Application').BrowseForFolder(0,'"+msg+"',0x1+0x100,0).self.path+'\\'",
	)
}

func pickFolder(msg string) (folder string, err error) {
	// Calling OleInitialize (or similar) is required for BIF_NEWDIALOGSTYLE.
	if hr := win.OleInitialize(); hr != win.S_OK && hr != win.S_FALSE {
		err = errors.New(fmt.Sprint("OleInitialize Error: ", hr))
		return
	}
	defer win.OleUninitialize()

	pathFromPIDL := func(pidl uintptr) (string, error) {
		var path [win.MAX_PATH]uint16
		if !win.SHGetPathFromIDList(pidl, &path[0]) {
			return "", errors.New("SHGetPathFromIDList failed")
		}

		return syscall.UTF16ToString(path[:]), nil
	}

	// We use this callback to disable the OK button in case of "invalid"
	// selections.
	callback := func(hwnd win.HWND, msg uint32, lp, wp uintptr) uintptr {
		const BFFM_SELCHANGED = 2
		if msg == BFFM_SELCHANGED {
			_, err := pathFromPIDL(lp)
			var enabled uintptr
			if err == nil {
				enabled = 1
			}

			const BFFM_ENABLEOK = win.WM_USER + 101

			win.SendMessage(hwnd, BFFM_ENABLEOK, 0, enabled)
		}

		return 0
	}

	var ownerHwnd win.HWND

	// We need to put the initial path into a buffer of at least MAX_LENGTH
	// length, or we may get random crashes.
	var buf [win.MAX_PATH]uint16
	copy(buf[:], syscall.StringToUTF16(""))

	const BIF_NEWDIALOGSTYLE = 0x00000040
	const BIF_SHAREABLE = 0x00008000

	bi := win.BROWSEINFO{
		HwndOwner: ownerHwnd,
		LpszTitle: syscall.StringToUTF16Ptr(msg),
		UlFlags:   BIF_NEWDIALOGSTYLE + BIF_SHAREABLE,
		Lpfn:      syscall.NewCallback(callback),
	}

	//win.SHParseDisplayName(&buf[0], 0, &bi.PidlRoot, 0, nil)

	pidl := win.SHBrowseForFolder(&bi)
	if pidl == 0 {
		return
	}
	defer win.CoTaskMemFree(pidl)

	folder, err = pathFromPIDL(pidl)
	return
}
