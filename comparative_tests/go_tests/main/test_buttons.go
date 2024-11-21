package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"log"
	"os"
)

func ButtonsTest() {
	go func() {
		// Create a new window
		window := new(app.Window)
		window.Option(app.Title("Test buttons"))
		window.Option(app.Size(unit.Dp(800), unit.Dp(600)))

		if err := draw(window); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}

func draw(window *app.Window) error {
	// ops are the operations that the window will execute
	var ops op.Ops

	// Create a new button
	var button1, button2, button3, button4 widget.Clickable

	// theme is the default theme for the window
	theme := material.NewTheme()

	// Listen for events in the window
	for {
		// Grab first event
		event := window.Event()

		type Context = layout.Context
		type Dimensions = layout.Dimensions

		// Check the type of event
		switch typ := event.(type) {

		// This is sent when the application should re-render
		case app.FrameEvent:
			gtx := app.NewContext(&ops, typ)

			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx Context) Dimensions {
						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Left:   unit.Dp(35),
							Right:  unit.Dp(35),
						}

						return margins.Layout(gtx,
							func(gtx Context) Dimensions {
								return layout.Flex{
									Axis:    layout.Horizontal,
									Spacing: layout.SpaceBetween,
								}.Layout(gtx,
									layout.Rigid(
										func(gtx Context) Dimensions {
											btn := material.Button(theme, &button1, "Button 1")
											if button1.Clicked(gtx) {
												fmt.Println("Button 1 clicked")
											}
											return btn.Layout(gtx)
										},
									),
									layout.Rigid(
										func(gtx Context) Dimensions {
											btn := material.Button(theme, &button2, "Button 2")
											if button2.Clicked(gtx) {
												fmt.Println("Button 2 clicked")
											}
											return btn.Layout(gtx)
										},
									),
									layout.Rigid(
										func(gtx Context) Dimensions {
											btn := material.Button(theme, &button3, "Button 3")
											if button3.Clicked(gtx) {
												fmt.Println("Button 3 clicked")
											}
											return btn.Layout(gtx)
										},
									),
									layout.Rigid(
										func(gtx Context) Dimensions {
											btn := material.Button(theme, &button4, "Button 4")
											if button4.Clicked(gtx) {
												fmt.Println("Button 4 clicked")
											}
											return btn.Layout(gtx)
										},
									),
								)
							},
						)
					},
				),

				layout.Rigid(
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
			)
			typ.Frame(gtx.Ops)

		// This is sent when the application should close
		case app.DestroyEvent:
			os.Exit(0)
		}
	}
}

func main() {
	fmt.Println("\nTest buttons:")
	ButtonsTest()
}
