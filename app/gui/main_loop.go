package gui

import (
	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/explorer"
	"image"
	"os"
)

func MainLoop(w *app.Window) error {
	expl := explorer.NewExplorer(w)
	var openBtn, saveBtn widget.Clickable
	var btn1, btn2, btn3, btn4, btn5 widget.Clickable
	th := material.NewTheme()
	fileChan := make(chan FileResult)
	saveChan := make(chan error)

	events := make(chan event.Event)
	acks := make(chan struct{})

	go func() {
		for {
			ev := w.Event()
			events <- ev
			<-acks
			if _, ok := ev.(app.DestroyEvent); ok {
				return
			}
		}
	}()
	var fileResult FileResult
	var ops op.Ops
	var imgOp paint.ImageOp
	for {
		select {
		case fileResult = <-fileChan:
			if fileResult.Error == nil {
				file, err := os.Open(fileResult.Name)
				if err == nil {
					img, _, err := image.Decode(file)
					if err == nil {
						imgOp = paint.NewImageOp(img)
					}
					file.Close()
				}
			}
			w.Invalidate()
		case <-saveChan:
			w.Invalidate()
		case e := <-events:
			expl.ListenEvents(e)
			switch e := e.(type) {
			case app.DestroyEvent:
				acks <- struct{}{}
				return e.Err
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				HandleOpenButtonClick(expl, fileChan, &openBtn, gtx)
				HandleSaveButtonClick(expl, saveChan, &saveBtn, fileResult, gtx)

				openBtnStyle, saveBtnStyle := CreateButtonStyles(th, &openBtn, &saveBtn)

				layout.Flex{
					Axis: layout.Vertical,
				}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Left:   unit.Dp(25),
							Right:  unit.Dp(25),
						}
						return margins.Layout(gtx,
							func(gt layout.Context) layout.Dimensions {
								return layout.Flex{
									Axis:    layout.Horizontal,
									Spacing: layout.SpaceBetween,
								}.Layout(gtx,
									layout.Rigid(openBtnStyle.Layout),
									layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
										return layout.Dimensions{}
									}),
									layout.Rigid(saveBtnStyle.Layout),
								)
							})

					}),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return layout.Dimensions{}
					}),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						if imgOp.Size().X > 0 && imgOp.Size().Y > 0 {
							imageWidget := Image{
								Src:      imgOp,
								Fit:      Contain,
								Position: layout.Center,
								Scale:    1,
							}
							return imageWidget.Layout(gtx)
						}
						return layout.Dimensions{}
					}),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return layout.Dimensions{}
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Left:   unit.Dp(25),
							Right:  unit.Dp(25),
						}
						return margins.Layout(gtx,
							func(gt layout.Context) layout.Dimensions {
								return layout.Flex{
									Axis:    layout.Horizontal,
									Spacing: layout.SpaceEvenly,
								}.Layout(gtx,
									layout.Rigid(material.Button(th, &btn1, "Button 1").Layout),
									layout.Rigid(material.Button(th, &btn2, "Button 2").Layout),
									layout.Rigid(material.Button(th, &btn3, "Button 3").Layout),
									layout.Rigid(material.Button(th, &btn4, "Button 4").Layout),
									layout.Rigid(material.Button(th, &btn5, "Button 5").Layout),
								)
							})
					}),
				)
				e.Frame(gtx.Ops)
			}
			acks <- struct{}{}
		}
	}
}
