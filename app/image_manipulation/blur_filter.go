package image_manipulation

import (
	"image"
	"image/color"
	"math"
)

func BlurFilter(img image.Image, intensity int) image.Image {
	bounds := img.Bounds()
	width, height := GetImageDimensions(img)

	blurredImage := image.NewRGBA(bounds)

	kernelSize := 2*intensity + 1
	kernel := make([][]float64, kernelSize)
	for i := range kernel {
		kernel[i] = make([]float64, kernelSize)
		for j := range kernel[i] {
			kernel[i][j] = 1.0 / float64(kernelSize*kernelSize)
		}
	}

	for y := intensity; y < height-intensity; y++ {
		for x := intensity; x < width-intensity; x++ {
			var rSum, gSum, bSum float64

			for ky := -intensity; ky <= intensity; ky++ {
				for kx := -intensity; kx <= intensity; kx++ {
					pxColor := color.RGBAModel.Convert(img.At(x+kx, y+ky)).(color.RGBA)
					weight := kernel[ky+intensity][kx+intensity]

					rSum += float64(pxColor.R) * weight
					gSum += float64(pxColor.G) * weight
					bSum += float64(pxColor.B) * weight
				}
			}

			blurredImage.Set(x, y, color.RGBA{
				R: uint8(math.Min(math.Max(rSum, 0), 255)),
				G: uint8(math.Min(math.Max(gSum, 0), 255)),
				B: uint8(math.Min(math.Max(bSum, 0), 255)),
				A: 255,
			})
		}
	}

	return blurredImage
}
