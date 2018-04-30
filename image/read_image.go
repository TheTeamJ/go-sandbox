package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	file, err := os.Open("./raw/go40.png")
	defer file.Close()
	if err != nil {
		return
	}
	img, imgFormat, _ := image.Decode(file)
	fmt.Println(imgFormat)

	// 画像情報
	fmt.Println(img.At(0, 0))
	bounds := img.Bounds()
	fmt.Printf("%d, %d\n", bounds.Min.X, bounds.Min.Y)
	fmt.Printf("%d, %d\n", bounds.Max.X, bounds.Max.Y)

	background := imaging.New(bounds.Max.X, bounds.Max.Y, color.Gray{222})
	fmt.Println(background.At(1, 1))

	// グレー化
	toGray(img)

	// saveAsBmp(img, "./out/img.bmp")
	// saveAsPng(img, "./out/img.png")
	saveAsMat(img, "./out/img.json")

	// 画像を生成する
	generateBmpImage("./out/img.json", "./out/out.bmp")
	// fmt.Println(img)
}
