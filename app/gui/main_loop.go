package gui

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/explorer"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func MainLoop(w *app.Window) error {
	var openButton, saveButton widget.Clickable
	var blurButton, grayscaleButton, contrastButton, rotateButton, resizeButton widget.Clickable
	var blurSlider, grayscaleSlider, contrastSlider widget.Float
	var applyBlurButton, applyGrayscaleButton, applyContrastButton widget.Clickable

	showBlurOptions := false
	showGrayscaleOptions := false
	showContrastOptions := false

	openButtonTheme := material.NewTheme()
	openButtonTheme.Palette.ContrastBg = color.NRGBA{R: 31, G: 206, B: 145, A: 255}

	saveButtonTheme := material.NewTheme()
	saveButtonTheme.Palette.ContrastBg = color.NRGBA{R: 165, G: 178, B: 173, A: 255}

	actionButtonsTheme := material.NewTheme()
	actionButtonsTheme.Palette.ContrastBg = color.NRGBA{R: 0, G: 255, B: 166, A: 255}
	actionButtonsTheme.Palette.ContrastFg = color.NRGBA{R: 0, G: 0, B: 0, A: 255}

	var ops op.Ops

	expl := explorer.NewExplorer(w)
	fileChan := make(chan FileResult)
	saveChan := make(chan error)

	var loadedImage image.Image

	for {
		e := w.Event()
		switch typ := e.(type) {
		case app.DestroyEvent:
			return typ.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, typ)

			HandleOpenButtonClick(expl, fileChan, &openButton, gtx)
			HandleSaveButtonClick(expl, saveChan, &saveButton, FileResult{}, gtx)

			select {
			case fileResult := <-fileChan:
				if fileResult.Error == nil {
					file, err := os.Open(fileResult.Name)
					if err == nil {
						defer file.Close()
						img, err := jpeg.Decode(file)
						if err == nil {
							loadedImage = img
						}
					}
				}
			default:
			}

			if blurButton.Clicked(gtx) {
				showBlurOptions = !showBlurOptions
			}

			if grayscaleButton.Clicked(gtx) {
				showGrayscaleOptions = !showGrayscaleOptions
			}

			if contrastButton.Clicked(gtx) {
				showContrastOptions = !showContrastOptions
			}

			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceBetween,
			}.Layout(gtx,
				// Top row: Open and Save buttons
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Left:   unit.Dp(25),
						Right:  unit.Dp(25),
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Horizontal,
							Spacing: layout.SpaceBetween,
						}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return material.Button(openButtonTheme, &openButton, "Open image").Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return material.Button(saveButtonTheme, &saveButton, "Save").Layout(gtx)
							}),
						)
					})
				}),
				// Middle: Display loaded image
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Left:   unit.Dp(25),
						Right:  unit.Dp(25),
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						if loadedImage != nil {
							imgOp := paint.NewImageOp(loadedImage)
							return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return Image{Src: imgOp}.Layout(gtx)
							})
						}
						return layout.Dimensions{}
					})
				}),
				// Bottom row: Action buttons
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Left:   unit.Dp(25),
						Right:  unit.Dp(25),
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Horizontal,
							Spacing: layout.SpaceBetween,
						}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return material.Button(actionButtonsTheme, &blurButton, "Blur").Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return material.Button(actionButtonsTheme, &grayscaleButton, "Grayscale").Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return material.Button(actionButtonsTheme, &contrastButton, "Contrast").Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return material.Button(actionButtonsTheme, &rotateButton, "Rotate").Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return material.Button(actionButtonsTheme, &resizeButton, "Resize").Layout(gtx)
							}),
						)
					})
				}),
				// Show blur options
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if showBlurOptions {
						return layout.Flex{
							Axis:    layout.Vertical,
							Spacing: layout.SpaceEvenly,
						}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Inset{
									Top:    unit.Dp(10),
									Bottom: unit.Dp(10),
									Left:   unit.Dp(25),
									Right:  unit.Dp(25),
								}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return layout.Flex{
										Axis: layout.Horizontal,
									}.Layout(gtx,
										layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
											return material.Slider(actionButtonsTheme, &blurSlider).Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Inset{
												Left: unit.Dp(10),
											}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
												return material.Label(actionButtonsTheme, unit.Sp(16), fmt.Sprintf("%.0f", blurSlider.Value*100)).Layout(gtx)
											})
										}),
									)
								})
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Inset{
									Top:    unit.Dp(10),
									Bottom: unit.Dp(10),
									Left:   unit.Dp(25),
									Right:  unit.Dp(25),
								}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return material.Button(actionButtonsTheme, &applyBlurButton, "Apply Blur").Layout(gtx)
								})
							}),
						)
					}
					return layout.Dimensions{}
				}),

				// Show grayscale options
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if showGrayscaleOptions {
						return layout.Flex{
							Axis:    layout.Vertical,
							Spacing: layout.SpaceEvenly,
						}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Inset{
									Top:    unit.Dp(10),
									Bottom: unit.Dp(10),
									Left:   unit.Dp(25),
									Right:  unit.Dp(25),
								}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return layout.Flex{
										Axis: layout.Horizontal,
									}.Layout(gtx,
										layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
											return material.Slider(actionButtonsTheme, &grayscaleSlider).Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Inset{
												Left: unit.Dp(10),
											}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
												return material.Label(actionButtonsTheme, unit.Sp(16), fmt.Sprintf("%.0f", grayscaleSlider.Value*100)).Layout(gtx)
											})
										}),
									)
								})
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Inset{
									Top:    unit.Dp(10),
									Bottom: unit.Dp(10),
									Left:   unit.Dp(25),
									Right:  unit.Dp(25),
								}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return material.Button(actionButtonsTheme, &applyGrayscaleButton, "Apply grayscale").Layout(gtx)
								})
							}),
						)
					}
					return layout.Dimensions{}
				}),

				// Show contrast options
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if showContrastOptions {
						return layout.Flex{
							Axis:    layout.Vertical,
							Spacing: layout.SpaceEvenly,
						}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Inset{
									Top:    unit.Dp(10),
									Bottom: unit.Dp(10),
									Left:   unit.Dp(25),
									Right:  unit.Dp(25),
								}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return layout.Flex{
										Axis: layout.Horizontal,
									}.Layout(gtx,
										layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
											return material.Slider(actionButtonsTheme, &contrastSlider).Layout(gtx)
										}),
										layout.Rigid(func(gtx layout.Context) layout.Dimensions {
											return layout.Inset{
												Left: unit.Dp(10),
											}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
												return material.Label(actionButtonsTheme, unit.Sp(16), fmt.Sprintf("%.0f", contrastSlider.Value*100)).Layout(gtx)
											})
										}),
									)
								})
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Inset{
									Top:    unit.Dp(10),
									Bottom: unit.Dp(10),
									Left:   unit.Dp(25),
									Right:  unit.Dp(25),
								}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return material.Button(actionButtonsTheme, &applyContrastButton, "Apply contrast").Layout(gtx)
								})
							}),
						)
					}
					return layout.Dimensions{}
				}),
			)

			typ.Frame(gtx.Ops)
		}
	}
}