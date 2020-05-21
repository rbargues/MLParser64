package main

import (
	"os"
	"image/png"
	"image"
	"image/draw"
	"math"
 )
 func imgCrop(fname string){
	cropCopy, _ := os.Open(fname)
	//630, 470
	defer cropCopy.Close()
	cropInfo, _ := png.Decode(cropCopy)

	bounds := cropInfo.Bounds()
	imgWidth := int((math.Round(float64(bounds.Max.X) * (float64(570.0) / float64(1015.0)))))
	imgHeight := int((math.Round(float64(bounds.Max.Y) * (float64(60) / float64(720)))))


	// used to start cropping
	newX := int(math.Round(float64(bounds.Max.X) * float64(230.0 / 1015.0)))
	newY := int(math.Round(float64(bounds.Max.Y) * float64(220.0 / 720.0)))

	m := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	draw.Draw(m, image.Rect(0, 0, imgWidth, imgHeight), cropInfo, image.Point{newX, newY}, draw.Src)

	newFilePath := "cropped.png"
	newImg, _ := os.Create(newFilePath)
	defer newImg.Close()

	png.Encode(newImg, m)
}
