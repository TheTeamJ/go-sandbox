package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"
	"strconv"
)

func stratum_decode(srcPath string, outDirPath string, high_quality bool) {
	file, _ := os.Open(srcPath)
	defer file.Close()
	img, _, _ := image.Decode(file)

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	fmt.Println(width)
	fmt.Println(height)

	// 行列計算したほうが速いけど、いまは気にしない
	for layer := 0; layer < 24; layer++ {
		outImg := image.NewRGBA(image.Rect(0, 0, width, height))

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				c := color.RGBAModel.Convert(img.At(x, y))
				gray := getBlackOrWhite(c, layer)
				if gray == 0 {
					// 黒色
					outImg.Set(x, y, color.Gray{uint8(0)})
				} else {
					// 白色
					outImg.Set(x, y, color.Gray{uint8(255)})
				}
			}
		}

		outImgPathName := outDirPath + "/" + strconv.Itoa(layer)
		fmt.Println(outImgPathName)
		if high_quality {
			saveAsBmp(outImg, outImgPathName+".bmp")
		} else {
			saveAsPng(outImg, outImgPathName+".png")
		}
	}
}

func getIdx(layer int) int {
	if layer < 7 {
		return 0
	}
	if layer < 15 {
		return 1
	}
	return 2
}

func getBlackOrWhite(c color.Color, layer int) int {
	rgba := c.(color.RGBA)
	c10 := uint8(0)
	if layer <= 7 {
		// 10進数のカラー値 (0~255)
		c10 = rgba.R
	} else if layer <= 15 {
		c10 = rgba.G
	} else {
		c10 = rgba.B
	}

	// 2進数に変換
	c2 := fmt.Sprintf("%b", c10)
	c2 = fmt.Sprintf("%08s", c2)
	ci2, _ := strconv.Atoi(string(c2[layer%8]))
	return ci2 // 0, 1
}
