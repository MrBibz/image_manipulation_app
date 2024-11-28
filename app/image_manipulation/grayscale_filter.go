package image_manipulation

import (
	"image"
	"image/color"
)

func GrayscaleFilter(img image.Image, intensity int) image.Image {
	bounds := img.Bounds()
	width, height := GetImageDimensions(img)

	grayedImage := image.NewRGBA(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			originalColor := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)

			grayValue := uint8(0.3*float64(originalColor.R) + 0.59*float64(originalColor.G) + 0.11*float64(originalColor.B))

			blendFactor := float64(intensity) / 100
			r := uint8((1-blendFactor)*float64(originalColor.R) + blendFactor*float64(grayValue))
			g := uint8((1-blendFactor)*float64(originalColor.G) + blendFactor*float64(grayValue))
			b := uint8((1-blendFactor)*float64(originalColor.B) + blendFactor*float64(grayValue))

			grayedImage.Set(x, y, color.RGBA{
				R: r,
				G: g,
				B: b,
				A: originalColor.A,
			})
		}
	}

	return grayedImage
}
