package main

import (
	"fmt"
	bf "go_tests/basic_functions"
	imf "go_tests/image_manipulation_functions"
)

func main() {
	const IMAGEPATH string = "../images/bread.jpg"
	const OUTPUTSPATH string = "./go_outputs/"

	// Test BlurImage
	fmt.Println("\nTest BlurImage : ")
	img := bf.ReadImage(IMAGEPATH)
	imf.BlurImage(img, 0, OUTPUTSPATH+"blurred_bread.jpg")
}
