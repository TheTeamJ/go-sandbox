package main

import (
	"encoding/json"
	"golang.org/x/image/bmp"
	"image"
	"image/color"
	"io/ioutil"
	"os"
)

func generateBmpImage(path string, outputPath string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	imageMat := ImageMat{}
	json.Unmarshal(file, &imageMat)

	img := coloring(imageMat.PixelColors)

	outFile, _ := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE, 0600)
	defer outFile.Close()
	bmp.Encode(outFile, img)
}

func coloring(pxColors [][][]int) *image.RGBA {
	height := len(pxColors)
	width := len(pxColors[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			rgba := pxColors[h][w]
			img.Set(w, h, color.RGBA{
				uint8(rgba[0]), uint8(rgba[1]), uint8(rgba[2]), uint8(rgba[3])})
		}
	}
	return img
}
