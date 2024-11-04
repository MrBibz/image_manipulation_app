package image_manipulation_functions

import (
	"fmt"
	"image/jpeg"
	"os"
)

func GetImageDimensions(imagePath string) (width, height int) {
	// Open the image file and check for errors
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("Error opening file : ", err)
		return
	}
	// Close the file when the function returns
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	// Decode the image file and check for errors
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("Error decoding file : ", err)
		return
	}

	// Get the dimensions of the image and return
	width = img.Bounds().Dx()
	height = img.Bounds().Dy()

	return
}
