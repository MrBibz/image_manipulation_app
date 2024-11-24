package gui

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/x/explorer"
	"image"
	"image/jpeg"
	"os"
)

func HandleOpenButtonClick(expl *explorer.Explorer, fileChan chan FileResult, openBtn *widget.Clickable, gtx layout.Context) {
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
}

func HandleSaveButtonClick(expl *explorer.Explorer, saveChan chan error, saveBtn *widget.Clickable, fileResult FileResult, gtx layout.Context) {
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
}
