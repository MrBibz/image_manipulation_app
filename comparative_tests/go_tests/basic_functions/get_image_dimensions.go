package basic_functions

import (
	"image"
)

func GetImageDimensions(img image.Image) (width, height int) {
	// Get the dimensions of the image and return
	width = img.Bounds().Dx()
	height = img.Bounds().Dy()

	return
}
