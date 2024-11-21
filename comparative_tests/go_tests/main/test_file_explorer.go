package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"gioui.org/font/gofont"
	"gioui.org/x/explorer"
)

func main() {
	go func() {
		w := new(app.Window)
		if err := fileSelectorLoop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

type (
	C = layout.Context
	D = layout.Dimensions
)

type FileResult struct {
	Error error
	Name  string
}

func fileSelectorLoop(w *app.Window) error {
	expl := explorer.NewExplorer(w)
	var openBtn, saveBtn widget.Clickable
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
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
	for {
		select {
		case fileResult = <-fileChan:
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
						// Assuming we have some image to save
						img := image.NewRGBA(image.Rect(0, 0, 100, 100))
						if err := jpeg.Encode(file, img, nil); err != nil {
							saveChan <- fmt.Errorf("failed encoding image file: %w", err)
							return
						}
					}(fileResult)
				}
				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(material.Button(th, &openBtn, "Open Image").Layout),
					layout.Rigid(func(gtx C) D {
						if fileResult.Error != nil {
							return material.H6(th, fileResult.Error.Error()).Layout(gtx)
						}
						if fileResult.Name != "" {
							return material.H6(th, fileResult.Name).Layout(gtx)
						}
						return D{}
					}),
					layout.Rigid(func(gtx C) D {
						if fileResult.Name == "" {
							gtx = gtx.Disabled()
						}
						return material.Button(th, &saveBtn, "Save Image").Layout(gtx)
					}),
					layout.Rigid(func(gtx C) D {
						if saveErr == nil {
							return D{}
						}
						return material.H6(th, saveErr.Error()).Layout(gtx)
					}),
				)
				e.Frame(gtx.Ops)
			}
			acks <- struct{}{}
		}
	}
}
