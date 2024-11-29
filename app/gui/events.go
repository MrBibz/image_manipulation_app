package gui

import (
	"app/image_manipulation"
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

func HandleSaveButtonClick(expl *explorer.Explorer, saveChan chan error, saveBtn *widget.Clickable, fileResult FileResult, modifiedImage image.Image, gtx layout.Context) {
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
			if err := jpeg.Encode(file, modifiedImage, nil); err != nil {
				saveChan <- fmt.Errorf("failed encoding image file: %w", err)
				return
			}
		}(fileResult)
	}
}

func ApplyFilters(img image.Image, blurIntensity int, grayscaleIntensity int, contrastFactor float64) image.Image {
	if blurIntensity > 0 {
		img = image_manipulation.BlurFilter(img, blurIntensity)
	}
	if grayscaleIntensity > 0 {
		img = image_manipulation.GrayscaleFilter(img, grayscaleIntensity)
	}
	if contrastFactor != 0 {
		img = image_manipulation.ContrastFilter(img, contrastFactor)
	}
	return img
}
