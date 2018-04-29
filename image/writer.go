package main

import (
	"image"
	"image/png"
	"os"

	"golang.org/x/image/bmp"
)

func saveAsPng(img image.Image, path string) {
	file, _ := os.Create(path)
	defer file.Close()
	png.Encode(file, img)
}

func saveAsBmp(img image.Image, path string) {
	file, _ := os.Create(path)
	defer file.Close()
	bmp.Encode(file, img)
}
