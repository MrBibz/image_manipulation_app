package gui

import (
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image/color"
)

func CreateButtonStyles(th *material.Theme, openBtn, saveBtn *widget.Clickable) (material.ButtonStyle, material.ButtonStyle) {
	openBtnStyle := material.Button(th, openBtn, "Open image")
	openBtnStyle.Background = color.NRGBA{R: 40, G: 95, B: 140, A: 255} // Green background
	openBtnStyle.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}    // White text

	saveBtnStyle := material.Button(th, saveBtn, "Save image")
	saveBtnStyle.Background = color.NRGBA{R: 145, G: 155, B: 157, A: 255} // Red background
	saveBtnStyle.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}      // White text

	return openBtnStyle, saveBtnStyle
}
