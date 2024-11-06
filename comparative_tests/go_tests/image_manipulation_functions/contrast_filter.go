package image_manipulation_functions

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"time"
)

func ContrastFilter(img image.Image, contrastFactor float64, outputsPath string) {
	start := time.Now()

	// Ensure the output directory exists
	if err := EnsureOutputDir(outputsPath); err != nil {
		fmt.Println(err)
		return
	}

	bounds := img.Bounds()
	contrastAdjustedImage := image.NewRGBA(bounds)

	// Calculate the contrast adjustment
	contrast := (259 * (contrastFactor + 255)) / (255 * (259 - contrastFactor))

	// Apply the contrast filter
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Get the original pixel color
			originalColor := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)

			// Adjust the contrast
			r := truncate(contrast*(float64(originalColor.R)-128) + 128)
			g := truncate(contrast*(float64(originalColor.G)-128) + 128)
			b := truncate(contrast*(float64(originalColor.B)-128) + 128)

			// Set the new pixel color
			contrastAdjustedImage.Set(x, y, color.RGBA{R: r, G: g, B: b, A: originalColor.A})
		}
	}

	// Create the output file
	outputFile, err := CreateOutputFile(outputsPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Encode the resized image as JPEG and save it to the output file
	if err := SaveImageAsJPEG(outputFile, contrastAdjustedImage); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Contrast adjusted image saved to:", outputsPath)
	fmt.Println("Execution time:", time.Since(start))
}

func truncate(value float64) uint8 {
	return uint8(math.Min(math.Max(value, 0), 255))
}
