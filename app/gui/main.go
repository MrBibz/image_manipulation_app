package gui

import (
	"fmt"
	"gioui.org/op/paint"
	"image"
	"image/jpeg"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/explorer"
)

func main() {
	go func() {
		w := new(app.Window)
		if err := mainLoop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func mainLoop(w *app.Window) error {
	expl := explorer.NewExplorer(w)
	var openBtn, saveBtn widget.Clickable
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
	var saveErr error
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
		case saveErr = <-saveChan:
			w.Invalidate()
		case e := <-events:
			expl.ListenEvents(e)
			switch e := e.(type) {
			case app.DestroyEvent:
				acks <- struct{}{}
				return e.Err
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				if openBtn.Clicked(gtx) {
					go func() {
						file, err := expl.ChooseFile("jpg")
						if err != nil {
							err = fmt.Errorf("failed opening image file: %w", err)
							fileChan <- FileResult{Error: err}
							return
						}
						defer file.Close()

						if f, ok := file.(*os.File); ok {
							fileChan <- FileResult{Name: f.Name()}
						} else {
							fileChan <- FileResult{Error: fmt.Errorf("failed to cast file to *os.File")}
						}
					}()
				}
				if saveBtn.Clicked(gtx) {
					go func(fileResult FileResult) {
						if fileResult.Error != nil {
							saveChan <- fmt.Errorf("no file loaded, cannot save")
							return
						}
						file, err := expl.CreateFile("file.jpg")
						if err != nil {
							saveChan <- fmt.Errorf("failed exporting image file: %w", err)
							return
						}
						defer func() {
							saveChan <- file.Close()
						}()
						img := image.NewRGBA(image.Rect(0, 0, 100, 100))
						if err := jpeg.Encode(file, img, nil); err != nil {
							saveChan <- fmt.Errorf("failed encoding image file: %w", err)
							return
						}
					}(fileResult)
				}
				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(material.Button(th, &openBtn, "Open Image").Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if fileResult.Error != nil {
							return material.H6(th, fileResult.Error.Error()).Layout(gtx)
						}
						if fileResult.Name != "" {
							return material.H6(th, fileResult.Name).Layout(gtx)
						}
						return layout.Dimensions{}
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if fileResult.Name == "" {
							gtx = gtx.Disabled()
						}
						return material.Button(th, &saveBtn, "Save Image").Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if saveErr == nil {
							return layout.Dimensions{}
						}
						return material.H6(th, saveErr.Error()).Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
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
				)
				e.Frame(gtx.Ops)
			}
			acks <- struct{}{}
		}
	}
}
