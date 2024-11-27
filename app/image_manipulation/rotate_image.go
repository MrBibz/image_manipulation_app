package image_manipulation

import (
	"fmt"
	"image"
	"time"
)

func RotateImage(img image.Image, angle int, outputsPath string) {
	start := time.Now()

	// Ensure the output directory exists
	if err := EnsureOutputDir(outputsPath); err != nil {
		fmt.Println(err)
		return
	}

	// Get the original image dimensions
	originalWidth, originalHeight := GetImageDimensions(img)

	// Initialisation of the rotated image
	var rotatedImage *image.RGBA
	switch angle {
	case 90, -270:
		rotatedImage = image.NewRGBA(image.Rect(0, 0, originalHeight, originalWidth))
		for y := 0; y < originalHeight; y++ {
			for x := 0; x < originalWidth; x++ {
				rotatedImage.Set(originalHeight-y-1, x, img.At(x, y))
			}
		}
	case -90, 270:
		rotatedImage = image.NewRGBA(image.Rect(0, 0, originalHeight, originalWidth))
		for y := 0; y < originalHeight; y++ {
			for x := 0; x < originalWidth; x++ {
				rotatedImage.Set(y, originalWidth-x-1, img.At(x, y))
			}
		}
	case 180, -180:
		rotatedImage = image.NewRGBA(image.Rect(0, 0, originalWidth, originalHeight))
		for y := 0; y < originalHeight; y++ {
			for x := 0; x < originalWidth; x++ {
				rotatedImage.Set(originalWidth-x-1, originalHeight-y-1, img.At(x, y))
			}
		}
	default:
		fmt.Println("Unsupported rotation angle. Please use 90, -90, 180 or -180 degrees.")
		return
	}

	// Create the output file
	outputFile, err := CreateOutputFile(outputsPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Encode the blurred image as JPEG and save it to the output file
	if err := SaveImageAsJPEG(outputFile, rotatedImage); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Rotated image saved to:", outputsPath)
	fmt.Println("Execution time:", time.Since(start))
}
