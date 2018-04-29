package main

import (
	"fmt"
	"image"
	_ "image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	file, err := os.Open("./data/go40.png")
	defer file.Close()
	if err != nil {
		return
	}
	img, imgFormat, _ := image.Decode(file)
	fmt.Println(imgFormat)
	// 画像情報
	fmt.Println(img.At(0, 0))
	fmt.Printf("%d, %d\n", img.Bounds().Min.X, img.Bounds().Min.Y)
	fmt.Printf("%d, %d\n", img.Bounds().Max.X, img.Bounds().Max.Y)
}
