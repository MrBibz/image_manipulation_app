package image_manipulation_functions

import (
	"fmt"
	bft "go_tests/basic_functions"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"
)

func ResizeImage(img image.Image, newWidth, newHeight int, outputsPath string) {
	// Ensure the output directory exists
	outputDir := filepath.Dir(outputsPath)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating output directory:", err)
			return
		}
	}

	// Original image dimensions
	originalWidth, originalHeight := bft.GetImageDimensions(img)

	resizedImage := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	// Resizing with bilinear interpolation
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			// Calculate original pixel coordinates
			originalX := float64(x) * float64(originalWidth) / float64(newWidth)
			originalY := float64(y) * float64(originalHeight) / float64(newHeight)

			// Bilinear interpolation
			x1, y1 := int(originalX), int(originalY)
			x2, y2 := x1+1, y1+1

			// Limit the coordinates to the original image bounds
			if x2 >= originalWidth {
				x2 = originalWidth - 1
			}
			if y2 >= originalHeight {
				y2 = originalHeight - 1
			}

			// Calculate the weights for the interpolation
			wx := originalX - float64(x1)
			wy := originalY - float64(y1)

			// Get the neighboring pixel colors
			c11 := color.RGBAModel.Convert(img.At(x1, y1)).(color.RGBA)
			c12 := color.RGBAModel.Convert(img.At(x1, y2)).(color.RGBA)
			c21 := color.RGBAModel.Convert(img.At(x2, y1)).(color.RGBA)
			c22 := color.RGBAModel.Convert(img.At(x2, y2)).(color.RGBA)

			// Interpolate the pixel color
			r := uint8((1-wx)*(1-wy)*float64(c11.R) + wx*(1-wy)*float64(c21.R) + (1-wx)*wy*float64(c12.R) + wx*wy*float64(c22.R))
			g := uint8((1-wx)*(1-wy)*float64(c11.G) + wx*(1-wy)*float64(c21.G) + (1-wx)*wy*float64(c12.G) + wx*wy*float64(c22.G))
			b := uint8((1-wx)*(1-wy)*float64(c11.B) + wx*(1-wy)*float64(c21.B) + (1-wx)*wy*float64(c12.B) + wx*wy*float64(c22.B))
			a := uint8((1-wx)*(1-wy)*float64(c11.A) + wx*(1-wy)*float64(c21.A) + (1-wx)*wy*float64(c12.A) + wx*wy*float64(c22.A))

			// Apply the new pixel color
			resizedImage.Set(x, y, color.RGBA{R: r, G: g, B: b, A: a})
		}
	}

	// Create the output file
	outputFile, err := os.Create(outputsPath)
	if err != nil {
		fmt.Println("Error creating output file : ", err)
		return
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {

		}
	}(outputFile)

	// Encode the resized image as JPEG and save it to the output file
	err = jpeg.Encode(outputFile, resizedImage, nil)
	if err != nil {
		fmt.Println("Error encoding output file : ", err)
		return
	}

	fmt.Println("Resized image saved to : ", outputsPath)
}
