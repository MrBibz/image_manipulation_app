package gui

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func MainLoop(w *app.Window) error {
	var leftButton, rightButton widget.Clickable
	var button1, button2, button3, button4, button5 widget.Clickable

	th := material.NewTheme()
	var ops op.Ops

	for {
		e := w.Event()
		switch typ := e.(type) {
		case app.DestroyEvent:
			return typ.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, typ)

			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceBetween,
			}.Layout(gtx,
				// Top row: Two buttons with margins
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    unit.Dp(10),
						Bottom: unit.Dp(10),
						Left:   unit.Dp(10),
						Right:  unit.Dp(10),
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Horizontal,
							Spacing: layout.SpaceBetween,
						}.Layout(gtx,
							// Left button with margin
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Inset{
									Right: unit.Dp(10),
								}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return material.Button(th, &leftButton, "Left").Layout(gtx)
								})
							}),
							// Space between
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								return layout.Dimensions{}
							}),
							// Right button with margin
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Inset{
									Left: unit.Dp(10),
								}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return material.Button(th, &rightButton, "Right").Layout(gtx)
								})
							}),
						)
					})
				}),
				// Middle image area with margins
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    unit.Dp(10),
						Bottom: unit.Dp(10),
						Left:   unit.Dp(20),
						Right:  unit.Dp(20),
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						// Placeholder for the image
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Dimensions{
								Size: gtx.Constraints.Min,
							}
						})
					})
				}),
				// Bottom row: Five buttons evenly spaced with margins
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    unit.Dp(10),
						Bottom: unit.Dp(10),
						Left:   unit.Dp(10),
						Right:  unit.Dp(10),
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Horizontal,
							Spacing: layout.SpaceBetween,
						}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return material.Button(th, &button1, "Button 1").Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return material.Button(th, &button2, "Button 2").Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return material.Button(th, &button3, "Button 3").Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return material.Button(th, &button4, "Button 4").Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return material.Button(th, &button5, "Button 5").Layout(gtx)
							}),
						)
					})
				}),
			)

			typ.Frame(gtx.Ops)
		}
	}
}
