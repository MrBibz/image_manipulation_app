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

	// Test RotateImage
	fmt.Println("\nTest RotateImage : ")
	img = bf.ReadImage(IMAGEPATH)
	imf.RotateImage(img, 90, OUTPUTSPATH+"rotated_bread.jpg")

	// Test ContrastFilter
	fmt.Println("\nTest ContrastFilter : ")
	img = bf.ReadImage(IMAGEPATH)
	imf.ContrastFilter(img, 100, OUTPUTSPATH+"contrast_bread.jpg")
}
