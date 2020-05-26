package main

import (
	"image"
	"image/png"
	"os"
	"github.com/kbinani/screenshot"
)

func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func screenshotGrab(r1 [2]int, r2 [2]int, filename string) {
	captureRectangle := image.Rect(r1[0], r1[1], r2[0], r2[1])
	img, err := screenshot.CaptureRect(captureRectangle)
	if err != nil {
		panic(err)
	}
	// filepath := fmt.Sprintf("screenshot.png")
	save(img, filename)
}