package main

import (
	"os"
	"os/user"
)

// Checkforfile ... Checks if a file/folder exists.
// I stole this from a random CLI launcher on GitHub sorry Sir-Boops
func Checkforfile(Path string) bool {
	// True if found
	// False if not
	ans := false
	if _, err := os.Stat(Path); err == nil {
		// File/Folder found!
		ans = true
	}
	return ans
}

// Getmaindir ... Gets the storage directory for modpacks, etc.
func Getmaindir() string {
	// TODO: platform seperate config folders
	// windows should be in appdata, linux in .config, that stuff
	// find out how to set based on platform?
	usr, _ := user.Current()
	return usr.HomeDir + "/.justlaunch/"
}
