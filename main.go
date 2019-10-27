//go:generate go run github.com/UnnoTed/fileb0x b0x.yml
package main

import (
	"fmt"
	"os"

	"github.com/ImVexed/muon"

	"justlaunch/webfiles"
	"net/http"
)

func main() {
	fileHandler := http.FileServer(webfiles.HTTP)

	cfg := &muon.Config{
		Title:      "Just Launch The Game",
		Width:      500,
		Height:     500,
		Titled:     true,
		Resizeable: true,
	}

	instance := muon.New(cfg, fileHandler)

	// Note: this was previously on a native check but
	// it's impossible to show a dialog in the frontend
	// without some serious bullshit on that front.
	// It's way better to just call it from the frontend.
	instance.Bind("checkJava", Checkjava)
	instance.Bind("die", Exitapp)

	if !Checkforfile(Getmaindir()) {
		os.MkdirAll(Getmaindir(), os.ModePerm)
	}

	if err := instance.Start(); err != nil {
		fmt.Println("PLEASE REPORT: https://github.com/tilda/justlaunch/issues")
		panic(err)
	}
}
