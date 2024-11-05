package image_manipulation_functions_tests

import (
	bft "go_tests/basic_functions_tests"
	"image"
	"image/color"
)

func BlurFilter(img image.Image) *image.RGBA {
	// Get the image dimensions
	bounds := img.Bounds()
	width, height := bft.GetImageDimensions(img)

	blurredImage := image.NewRGBA(bounds)

	// Create the blur kernel
	kernel := [3][3]float64{
		{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
		{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
		{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
	}

	// Apply the blur filter
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			var rSum, gSum, bSum float64

			// Apply the kernel
			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					// Get the neighboring pixel
					pxColor := color.RGBAModel.Convert(img.At(x+kx, y+ky)).(color.RGBA)
					weight := kernel[ky+1][kx+1]

					// Add the weighted color to the sum
					rSum += float64(pxColor.R) * weight
					gSum += float64(pxColor.G) * weight
					bSum += float64(pxColor.B) * weight
				}
			}

			// Set the new pixel color
			blurredImage.Set(x, y, color.RGBA{
				R: uint8(rSum),
				G: uint8(gSum),
				B: uint8(bSum),
				A: 255,
			})
		}
	}

	return blurredImage
}
