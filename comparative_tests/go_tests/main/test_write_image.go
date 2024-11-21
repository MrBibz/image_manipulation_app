package main

import (
	"fmt"
	bf "go_tests/basic_functions"
	imf "go_tests/image_manipulation_functions"
	"image"
	"image/color"
	"time"
)

func InvertColors(img image.Image, outputsPath string) {
	start := time.Now()

	// Ensure the output directory exists
	if err := imf.EnsureOutputDir(outputsPath); err != nil {
		fmt.Println(err)
		return
	}

	// Get the image dimensions
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	invertedImage := image.NewRGBA(bounds)

	// Invert the colors
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Get the pixel color
			pxColor := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)

			// Invert the color
			invertedImage.Set(x, y, color.RGBA{
				R: 255 - pxColor.R,
				G: 255 - pxColor.G,
				B: 255 - pxColor.B,
				A: pxColor.A,
			})
		}
	}

	// Create the output file
	outputFile, err := imf.CreateOutputFile(outputsPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Encode the inverted image as JPEG and save it to the output file
	if err := imf.SaveImageAsJPEG(outputFile, invertedImage); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Inverted color image saved to:", outputsPath)
	fmt.Println("Execution time:", time.Since(start))
}

func main() {
	imagePath := "../images/bread.jpg"
	outputsPath := "./outputs/test_write_image_output.jpg"

	img := bf.ReadImage(imagePath)
	InvertColors(img, outputsPath)
}
