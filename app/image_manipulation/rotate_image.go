package image_manipulation

import (
	"image"
	"image/draw"
)

func RotateImage(img image.Image, angle int) image.Image {
	bounds := img.Bounds()
	var rotatedImage *image.RGBA

	switch angle {
	case 90, -270:
		rotatedImage = image.NewRGBA(image.Rect(0, 0, bounds.Dy(), bounds.Dx()))
		for y := 0; y < bounds.Dy(); y++ {
			for x := 0; x < bounds.Dx(); x++ {
				rotatedImage.Set(bounds.Dy()-y-1, x, img.At(x, y))
			}
		}
	case -90, 270:
		rotatedImage = image.NewRGBA(image.Rect(0, 0, bounds.Dy(), bounds.Dx()))
		for y := 0; y < bounds.Dy(); y++ {
			for x := 0; x < bounds.Dx(); x++ {
				rotatedImage.Set(y, bounds.Dx()-x-1, img.At(x, y))
			}
		}
	case 180, -180:
		rotatedImage = image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
		for y := 0; y < bounds.Dy(); y++ {
			for x := 0; x < bounds.Dx(); x++ {
				rotatedImage.Set(bounds.Dx()-x-1, bounds.Dy()-y-1, img.At(x, y))
			}
		}
	default:
		rotatedImage = image.NewRGBA(bounds)
		draw.Draw(rotatedImage, bounds, img, image.Point{}, draw.Src)
	}

	return rotatedImage
}
