package main

import (
	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"os"
)

func main() {
	go func() {
		// Create a new window
		window := new(app.Window)
		window.Option(app.Title("Test buttons"))
		window.Option(app.Size(unit.Dp(800), unit.Dp(600)))

		// ops are the operations that the window will execute
		var ops op.Ops

		// Create a new button
		var button widget.Clickable

		// theme is the default theme for the window
		theme := material.NewTheme()

		// Listen for events in the window
		for {
			// Grab first event
			event := window.Event()

			// Check the type of event
			switch typ := event.(type) {

			// This is sent when the application should re-render
			case app.FrameEvent:
				gtx := app.NewContext(&ops, typ)
				btn := material.Button(theme, &button, "Button")
				btn.Layout(gtx)
				typ.Frame(gtx.Ops)

			// This is sent when the application should close
			case app.DestroyEvent:
				os.Exit(0)
			}
		}
	}()

	app.Main()
}
