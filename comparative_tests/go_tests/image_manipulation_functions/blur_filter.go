package image_manipulation_functions

import (
	"fmt"
	bft "go_tests/basic_functions"
	"image"
	"image/color"
	"math"
)

func BlurFilter(img image.Image, blurIntensity int, outputsPath string) {
	// Ensure the output directory exists
	if err := EnsureOutputDir(outputsPath); err != nil {
		fmt.Println(err)
		return
	}

	// Get the image dimensions
	bounds := img.Bounds()
	width, height := bft.GetImageDimensions(img)

	blurredImage := image.NewRGBA(bounds)

	// Create the blur kernel
	kernelSize := 2*blurIntensity + 1
	kernel := make([][]float64, kernelSize)
	for i := range kernel {
		kernel[i] = make([]float64, kernelSize)
		for j := range kernel[i] {
			kernel[i][j] = 1.0 / float64(kernelSize*kernelSize)
		}
	}

	// Apply the blur filter
	for y := blurIntensity; y < height-blurIntensity; y++ {
		for x := blurIntensity; x < width-blurIntensity; x++ {
			var rSum, gSum, bSum float64

			// Apply the kernel
			for ky := -blurIntensity; ky <= blurIntensity; ky++ {
				for kx := -blurIntensity; kx <= blurIntensity; kx++ {
					// Get the neighboring pixel
					pxColor := color.RGBAModel.Convert(img.At(x+kx, y+ky)).(color.RGBA)
					weight := kernel[ky+blurIntensity][kx+blurIntensity]

					// Add the weighted color to the sum
					rSum += float64(pxColor.R) * weight
					gSum += float64(pxColor.G) * weight
					bSum += float64(pxColor.B) * weight
				}
			}

			// Set the new pixel color
			blurredImage.Set(x, y, color.RGBA{
				R: uint8(math.Min(math.Max(rSum, 0), 255)),
				G: uint8(math.Min(math.Max(gSum, 0), 255)),
				B: uint8(math.Min(math.Max(bSum, 0), 255)),
				A: 255,
			})
		}
	}

	// Create the output file
	outputFile, err := CreateOutputFile(outputsPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Encode the blurred image as JPEG and save it to the output file
	if err := SaveImageAsJPEG(outputFile, blurredImage); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Blurred image saved to:", outputsPath)
}
