package main

import (
	"fmt"
	"os/exec"
)

// Checkjava ... Checks if Java is installed.
func Checkjava() (installed bool) {
	installed = false
	fmt.Println("Checking Java...")
	java, err := exec.Command("java -version").Output()
	if err == nil {
		installed = true
		fmt.Printf("Java version: %s\n", java)
	}
	return installed
}
