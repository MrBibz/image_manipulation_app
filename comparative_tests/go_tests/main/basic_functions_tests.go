package main

import (
	"fmt"
	bf "go_tests/basic_functions"
)

func main() {
	const IMAGEPATH string = "../images/bread.jpg"

	// Test ReadImage
	fmt.Println("\nTest ReadImage : ")
	img := bf.ReadImage(IMAGEPATH)
	fmt.Println("Image : ", img)

	// Test GetImageDimensions
	fmt.Println("\nTest GetImageDimensions : ")
	width, height := bf.GetImageDimensions(img)
	fmt.Println("Image dimensions: ", width, "x", height)

	// Test ExtractPixels
	fmt.Println("\nTest ExtractPixels : ")
	bf.ExtractPixels(IMAGEPATH)
}
