package image_manipulation_test_functions

import (
	"fmt"
	"image/color"
)

func ExtractPixels(imagePath string) {
	// Get the dimensions of the image
	img := ReadImage(imagePath)
	width, height := GetImageDimensions(img)

	// Create a matrix to stock the pixels
	pixels := make([][]color.RGBA, height)
	for y := 0; y < height; y++ {
		pixels[y] = make([]color.RGBA, width)
		for x := 0; x < width; x++ {
			// Get the color of the pixel at (x, y)
			rgba := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			pixels[y][x] = rgba
		}
	}

	// Print the five first pixels of the first five rows
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			fmt.Printf("Pixel [%d, %d]: R=%d G=%d B=%d A=%d\n",
				y, x, pixels[y][x].R, pixels[y][x].G, pixels[y][x].B, pixels[y][x].A)
		}
	}
}
