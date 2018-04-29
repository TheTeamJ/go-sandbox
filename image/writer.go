package main

import (
	"image"
	"image/png"
	"os"
)

func saveAsPng(img image.Image, path string) {
	file, _ := os.Create(path)
	defer file.Close()
	png.Encode(file, img)
}
