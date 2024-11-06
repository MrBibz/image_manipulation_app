package image_manipulation_functions

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
)

// EnsureOutputDir ensures the output directory exists
func EnsureOutputDir(outputsPath string) error {
	outputDir := filepath.Dir(outputsPath)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating output directory: %w", err)
		}
	}
	return nil
}

// CreateOutputFile creates the output file
func CreateOutputFile(outputsPath string) (*os.File, error) {
	outputFile, err := os.Create(outputsPath)
	if err != nil {
		return nil, fmt.Errorf("error creating output file: %w", err)
	}
	return outputFile, nil
}

// SaveImageAsJPEG encodes the image as JPEG and saves it to the output file
func SaveImageAsJPEG(outputFile *os.File, img image.Image) error {
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			fmt.Println("Error closing output file:", err)
		}
	}(outputFile)

	err := jpeg.Encode(outputFile, img, nil)
	if err != nil {
		return fmt.Errorf("error encoding output file: %w", err)
	}
	return nil
}
