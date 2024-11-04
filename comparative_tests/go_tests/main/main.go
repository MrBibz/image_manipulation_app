package main

import (
	"fmt"
	imf "go_tests/image_manipulation_functions"
)

func main() {
	// Test GetImageDimensions
	fmt.Println("\nTest GetImageDimensions : ")
	testGetImageDimensions()
}

func testGetImageDimensions() {
	const IMAGEPATH string = "../images/bread.jpg"

	width, height := imf.GetImageDimensions(IMAGEPATH)

	fmt.Println("Image dimensions: ", width, "x", height)
}
