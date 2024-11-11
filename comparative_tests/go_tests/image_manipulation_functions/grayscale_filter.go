package image_manipulation_functions

import (
	"fmt"
	bft "go_tests/basic_functions"
	"image"
	"image/color"
	"time"
)

func GrayscaleFilter(img image.Image, intensity int, outputsPath string) {
	start := time.Now()

	// Ensure the output directory exists
	if err := EnsureOutputDir(outputsPath); err != nil {
		fmt.Println(err)
		return
	}

	// Get the image dimensions
	bounds := img.Bounds()
	width, height := bft.GetImageDimensions(img)

	grayedImage := image.NewRGBA(bounds)

	// Apply the grayscale filter
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			originalColor := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)

			// Calculate the luminosity of the pixel
			grayValue := uint8(0.3*float64(originalColor.R) + 0.59*float64(originalColor.G) + 0.11*float64(originalColor.B))

			// Apply the grayscale filter with the intensity (0 = no change, 100 = full grayscale)
			blendFactor := float64(intensity) / 100
			r := uint8((1-blendFactor)*float64(originalColor.R) + blendFactor*float64(grayValue))
			g := uint8((1-blendFactor)*float64(originalColor.G) + blendFactor*float64(grayValue))
			b := uint8((1-blendFactor)*float64(originalColor.B) + blendFactor*float64(grayValue))

			// Set the new pixel color
			grayedImage.Set(x, y, color.RGBA{
				R: r,
				G: g,
				B: b,
				A: originalColor.A,
			})
		}
	}

	// Create the output file
	outputFile, err := CreateOutputFile(outputsPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Encode the grayed image as JPEG and save it to the output file
	if err := SaveImageAsJPEG(outputFile, grayedImage); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Grayed image saved to : ", outputsPath)
	fmt.Println("Execution time : ", time.Since(start))
}