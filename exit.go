package main

import (
	"fmt"
	"os"
)

// Exitapp ... Exits the application safely.
func Exitapp() {
	fmt.Println("Exitapp called")
	os.Exit(0)
}
