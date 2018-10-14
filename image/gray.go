package main

import (
	"image"
	"image/color"
	// "image/png"
)

// グレー化
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

// 2値値
const threshold = 128

func binarize(img image.Image) image.Image {
	bounds := img.Bounds()
	dest := image.NewGray(bounds)
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y))
			gray, _ := c.(color.Gray)
			if gray.Y > threshold {
				gray.Y = 255
			} else {
				gray.Y = 0
			}
			dest.Set(x, y, gray)
		}
	}
	return dest
}
