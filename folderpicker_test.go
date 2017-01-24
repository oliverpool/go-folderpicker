package folderpicker

import "testing"
import "runtime"

func TestCleanFolder(t *testing.T) {
	if runtime.GOOS == "windows" {
		windowsCleanFolder(t)
	} else {
		unixCleanFolder(t)
	}
}

func unixCleanFolder(t *testing.T) {
	stringSet := []struct {
		in  string
		out string
	}{
		{},
		{
			in:  `/`,
			out: `/`,
		},
		{
			in:  `/unix/style`,
			out: `/unix/style`,
		},
		{
			in:  `/unix/style/trailing/`,
			out: `/unix/style/trailing`,
		},
		{
			in:  ` /unix/style/with spaces at the end/ `,
			out: `/unix/style/with spaces at the end`,
		},
	}

	for _, s := range stringSet {
		out := cleanFolder([]byte(s.in))
		if s.out != out {
			t.Error("Got:", out, "Expected:", s.out)
		}
	}
}

func windowsCleanFolder(t *testing.T) {
	stringSet := []struct {
		in  string
		out string
	}{
		{},
		{
			in:  `\`, // Powershell selected nothing
			out: ``,
		},
		{
			in:  ` C:\Windows\Style  `, // with spaces
			out: `C:\Windows\Style`,
		},
		{
			in:  `C:\`,
			out: `C:\`,
		},
		{
			in:  `C:\Windows\Style`,
			out: `C:\Windows\Style`,
		},
		{
			in:  `C:\Windows\Style\trailing\`,
			out: `C:\Windows\Style\trailing`,
		},
		{
			in:  `\\WindowsNetworkShare\Name`,
			out: `\\WindowsNetworkShare\Name`,
		},
		{
			in:  `\\WindowsNetworkShare\Name\trailing\`,
			out: `\\WindowsNetworkShare\Name\trailing`,
		},
	}

	for _, s := range stringSet {
		out := cleanFolder([]byte(s.in))
		if s.out != out {
			t.Error("Got:", out, "Expected:", s.out)
		}
	}
}
