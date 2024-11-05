package main

import (
	"fmt"
	imf "go_tests/basic_test_functions"
)

func main() {
	const IMAGEPATH string = "../images/bread.jpg"

	// Test ReadImage
	fmt.Println("\nTest ReadImage : ")
	img := imf.ReadImage(IMAGEPATH)
	fmt.Println("Image : ", img)

	// Test GetImageDimensions
	fmt.Println("\nTest GetImageDimensions : ")
	width, height := imf.GetImageDimensions(img)
	fmt.Println("Image dimensions: ", width, "x", height)

	// Test ExtractPixels
	fmt.Println("\nTest ExtractPixels : ")
	imf.ExtractPixels(IMAGEPATH)
}
