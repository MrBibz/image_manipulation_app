package gui

import (
	im "app/image_manipulation"
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

func ApplyAllManipulations(originalImage image.Image, manipulations []im.Manipulation) image.Image {
	img := originalImage
	for _, manipulation := range manipulations {
		switch manipulation.Type {
		case "blur":
			img = im.BlurFilter(img, manipulation.Intensity)
		case "grayscale":
			img = im.GrayscaleFilter(img, manipulation.Intensity)
		case "contrast":
			img = im.ContrastFilter(img, manipulation.Factor)
		case "rotate":
			img = im.RotateImage(img, manipulation.Angle)
		case "resize":
			img = im.ResizeImage(img, manipulation.NewWidth, manipulation.NewHeight)
		}
	}
	return img
}

func AddOrReplaceManipulation(manipulations []im.Manipulation, newManipulation im.Manipulation) []im.Manipulation {
	for i, manipulation := range manipulations {
		if manipulation.Type == newManipulation.Type {
			manipulations[i] = newManipulation
			return manipulations
		}
	}
	return append(manipulations, newManipulation)
}
