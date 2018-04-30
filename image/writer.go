package main

import (
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	// "strings"

	// "github.com/skelterjohn/go.matrix"
	"golang.org/x/image/bmp"
)

type ImageMat struct {
	PixelColors [][][]int
}

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

func saveAsMat(img image.Image, path string) {
	bound := img.Bounds().Max
	pixels := make([][][]int, bound.Y)

	for y := 0; y < bound.Y; y++ {
		pixels[y] = make([][]int, bound.X)
		for x := 0; x < bound.X; x++ {
			c := color.RGBAModel.Convert(img.At(x, y))
			rgba := c.(color.RGBA)
			pixels[y][x] = []int{int(rgba.R), int(rgba.G), int(rgba.B), int(rgba.A)}
		}
	}

	imgMat := ImageMat{}
	imgMat.PixelColors = pixels

	outputJson, _ := json.Marshal(imgMat)
	_ = ioutil.WriteFile(path, outputJson, 0644)
}
