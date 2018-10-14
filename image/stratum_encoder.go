package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"
	"strconv"
)

// ひとまずべた書きする

type Stratum struct {
	PixelColors [][][]string
	width       int
	height      int
}

func stratum_encode(dirPath string, pageNums int) *image.RGBA {
	res := initStratum(dirPath)
	for p := 0; p < pageNums; p++ {
		res = putColors(res, dirPath, p)
	}
	res = padding(res)

	img := bin2dec(res)
	// fmt.Println(img)
	return img
}

// 最初の画像を読み込んで、画像の縦横長を確定する
func initStratum(dirPath string) Stratum {
	res := Stratum{}

	// 最初の画像のファイル名は0である
	filePath := dirPath + "/0.png"
	file, _ := os.Open(filePath)
	defer file.Close()
	img, _, _ := image.Decode(file)
	bounds := img.Bounds()

	res.width = bounds.Max.X
	res.height = bounds.Max.Y

	// 行列ライブラリを使うとこの辺も綺麗に初期化できるかも
	pixels := make([][][]string, res.height)
	for y := 0; y < res.height; y++ {
		pixels[y] = make([][]string, res.width)
		for x := 0; x < res.width; x++ {
			pixels[y][x] = []string{"", "", ""}
		}
	}
	res.PixelColors = pixels

	return res
}

func putColors(res Stratum, dirPath string, pageNum int) Stratum {
	filePath := dirPath + "/" + strconv.Itoa(pageNum) + ".png"
	fmt.Println(filePath)
	file, _ := os.Open(filePath)
	defer file.Close()
	img, _, _ := image.Decode(file)

	// 二値化する
	img = binarize(img)

	// 更新するPixelColorsのidx
	idx := 0
	if pageNum <= 7 {
		idx = 0
	} else if pageNum <= 15 {
		idx = 1
	} else if pageNum <= 23 {
		idx = 2
	}

	for y := 0; y < res.height; y++ {
		for x := 0; x < res.width; x++ {
			nowColor := res.PixelColors[y][x][idx]
			c := color.GrayModel.Convert(img.At(x, y))
			gray := c.(color.Gray).Y
			if gray == 0 {
				// 黒色
				res.PixelColors[y][x][idx] = nowColor + "0"
			} else {
				// 白色
				res.PixelColors[y][x][idx] = nowColor + "1"
			}
		}
	}
	return res
}

// 空白詰めする
// 値が入っていないところを0で埋める
func padding(res Stratum) Stratum {
	for y := 0; y < res.height; y++ {
		for x := 0; x < res.width; x++ {
			for idx := 0; idx < 3; idx++ {
				nowColor := res.PixelColors[y][x][idx]
				if len(nowColor) < 8 {
					res.PixelColors[y][x][idx] = fmt.Sprintf("%08s", nowColor)
				}
			}
		}
	}
	return res
}

// 2進数文字列のStratum画像を10進数のimage.RGBAに変換する
func bin2dec(res Stratum) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, res.width, res.height))

	for y := 0; y < res.height; y++ {
		for x := 0; x < res.width; x++ {
			rgba := []int64{0, 0, 0, 255}
			for idx := 0; idx < 3; idx++ {
				rgba[idx], _ = strconv.ParseInt(res.PixelColors[y][x][idx], 2, 64)
			}
			img.Set(x, y, color.RGBA{
				uint8(rgba[0]), uint8(rgba[1]), uint8(rgba[2]), uint8(rgba[3])})
			// fmt.Println(rgba)
		}
	}
	return img
}
