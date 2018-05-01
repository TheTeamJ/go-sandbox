package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"

	"github.com/disintegration/imaging"
)

func main2() {
	file, err := os.Open("./out/comic1.bmp")
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

	// 二値化
	// img = binarize(img)

	saveAsMat(img, "./out/img.json")
	// 画像を生成する
	// img = generateBmpImage("./out/img.json", "./out/out.bmp")
}

func main3() {
	img := stratum_encode("./raw/tonkatsu", 22)
	saveAsBmp(img, "./out/tonkatsu.bmp")
	saveAsPng(img, "./out/tonkatsu.png")
}

func main() {
	stratum_decode("./out/tonkatsu.png", "./out/tonkatsu")
}

func main4() {
	for p := 0; p < 22; p++ {
		file, _ := os.Open("./raw/tonkatsu/" + strconv.Itoa(p) + ".png")
		defer file.Close()
		img, _, _ := image.Decode(file)
		img = binarize(img)
		saveAsPng(img, "./out/tonkatsuBin/"+strconv.Itoa(p)+".png")
	}
}
