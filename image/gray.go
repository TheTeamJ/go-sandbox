package main

import (
	"image"
	"image/color"
	// "image/png"
)

func toGray(img image.Image) image.Image {
	bounds := img.Bounds()
	dest := image.NewGray(bounds)
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y))
			gray, _ := c.(color.Gray)
			dest.Set(x, y, gray)
		}
	}
	return dest
}
