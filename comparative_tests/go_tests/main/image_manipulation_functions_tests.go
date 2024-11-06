package main

import (
	"fmt"
	bf "go_tests/basic_functions"
	imf "go_tests/image_manipulation_functions"
)

func main() {
	const IMAGEPATH string = "../images/bread.jpg"
	const OUTPUTSPATH string = "./go_outputs/"

	// Test BlurFilter
	fmt.Println("\nTest BlurFilter : ")
	img := bf.ReadImage(IMAGEPATH)
	imf.BlurFilter(img, 3, OUTPUTSPATH+"blurred_bread.jpg")

	// Test GrayscaleFilter
	fmt.Println("\nTest GrayscaleFilter : ")
	img = bf.ReadImage(IMAGEPATH)
	imf.GrayscaleFilter(img, 100, OUTPUTSPATH+"grayscale_bread.jpg")

	// Test ResizeImage
	fmt.Println("\nTest ResizeImage : ")
	img = bf.ReadImage(IMAGEPATH)
	imf.ResizeImage(img, 800, 400, OUTPUTSPATH+"resized_bread.jpg")
}
