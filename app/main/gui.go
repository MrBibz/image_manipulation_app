package main

import (
	"app/gui"
	"log"
	"os"

	"gioui.org/app"
)

func main() {
	go func() {
		w := new(app.Window)
		if err := gui.MainLoop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
