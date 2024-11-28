package image_manipulation

import (
	"image"
	"image/color"
	"math"
)

func ContrastFilter(img image.Image, contrastFactor float64) image.Image {
	bounds := img.Bounds()
	contrastAdjustedImage := image.NewRGBA(bounds)

	contrast := (259 * (contrastFactor + 255)) / (255 * (259 - contrastFactor))

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)

			r := truncate(contrast*(float64(originalColor.R)-128) + 128)
			g := truncate(contrast*(float64(originalColor.G)-128) + 128)
			b := truncate(contrast*(float64(originalColor.B)-128) + 128)

			contrastAdjustedImage.Set(x, y, color.RGBA{R: r, G: g, B: b, A: originalColor.A})
		}
	}

	return contrastAdjustedImage
}

func truncate(value float64) uint8 {
	return uint8(math.Min(math.Max(value, 0), 255))
}
