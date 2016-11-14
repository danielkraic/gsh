package main

import "github.com/fatih/color"

// conect to server operation
func connect(matched []Server, printOnly bool) {
	if len(matched) == 0 {
		color.Cyan("No server match patterns")
	} else if len(matched) == 1 {
		color.Green("%s", matched[0].getConnectionString())
		if !printOnly {
			matched[0].connect()
		}
	} else {
		color.Cyan("Multiple servers match patterns:")
		for _, s := range matched {
			color.White(s.getConnectionString())
		}
	}
}

// upload file to server operation
func upload(src string, dest string, matched []Server, printOnly bool) {
	if len(matched) == 0 {
		color.Cyan("No server match patterns")
	} else if len(matched) == 1 {
		color.Green("%s", matched[0].getUploadString(src, dest))
		if !printOnly {
			matched[0].upload(src, dest)
		}
	} else {
		color.Cyan("Multiple servers match patterns:")
		for _, s := range matched {
			color.White(s.getUploadString(src, dest))
		}
	}
}

// download file from server operation
func download(src string, dest string, matched []Server, printOnly bool) {
	if len(matched) == 0 {
		color.Cyan("No server match patterns")
	} else if len(matched) == 1 {
		color.Green("%s", matched[0].getDownloadString(src, dest))
		if !printOnly {
			matched[0].download(src, dest)
		}
	} else {
		color.Cyan("Multiple servers match patterns:")
		for _, s := range matched {
			color.White(s.getDownloadString(src, dest))
		}
	}
}
