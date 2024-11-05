package basic_functions_tests

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func ReadImage(filePath string) (img image.Image) {
	// Open the image file and check for errors
	file, err := os.Open(filePath)
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
	img, err = jpeg.Decode(file)
	if err != nil {
		fmt.Println("Error decoding file : ", err)
		return
	}

	return
}
