package main

import (
	"os"
	"image/png"
	"image"
	"image/draw"
	"math"
	"fmt"
	"image/color"
 )
 func imgCrop(fname string){
	 //450-350 	480-360
	cropCopy, _ := os.Open(fname)
	//630, 470
	defer cropCopy.Close()
	cropInfo, _ := png.Decode(cropCopy)

	bounds := cropInfo.Bounds()
	imgWidth := int((math.Round(float64(bounds.Max.X) * (float64(100.0) / float64(1015.0)))))
	imgHeight := int((math.Round(float64(bounds.Max.Y) * (float64(60.0) / float64(720)))))


	// used to start cropping
	newX := int(math.Round(float64(bounds.Max.X) * float64(170.0 / 1015.0)))
	newY := int(math.Round(float64(bounds.Max.Y) * float64(20.0 / 720.0)))

	m := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	draw.Draw(m, image.Rect(0, 0, imgWidth, imgHeight), cropInfo, image.Point{newX, newY}, draw.Src)

	newFilePath := "cropped.png"
	newImg, _ := os.Create(newFilePath)
	defer newImg.Close()

	png.Encode(newImg, m)
}
func redScreen() bool{
	imageFile, err := os.Open("./cropped.png")
	if err != nil {
		panic(err)
	}
	defer imageFile.Close()
	imageInfo, err := png.Decode(imageFile)
	if err != nil {
		panic(err)
	}
	bounds := imageInfo.Bounds()

	redCt := 0
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			rgba:= imageInfo.At(x,y).(color.RGBA)
			fmt.Printf("%v %v %v\n", rgba.R, rgba.G, rgba.B)
			if (rgba.R >= 245 && rgba.R <= 255) && (rgba.G >= 105 && rgba.G <= 115) && (rgba.B >= 105 && rgba.B <= 115) {
				redCt ++
			}
		}
	}
	if float64(redCt) / float64(bounds.Max.X * bounds.Max.Y) > 0.99 {
		return true
	} else {
		return false
	}
}
func removeNonRed(input string, output string) {
	copyImg, _ := os.Open(input)
	defer copyImg.Close()
	copyInfo, _ := png.Decode(copyImg)
	bounds := copyInfo.Bounds()

	m := image.NewRGBA(image.Rect(0, 0, bounds.Max.X, bounds.Max.Y))
	imgColor := color.RGBA{0,0,0,255}

	for x:= 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			rgba := copyInfo.At(x, y).(color.RGBA)
			if rgba.R < 240 {
				m.Set(x, y, imgColor)
			} else {
				fmt.Printf("%v\n", rgba)
				m.Set(x, y, rgba)
			}
		}
	}
	newImg, _ := os.Create(output)
	defer newImg.Close()
	png.Encode(newImg, m)
}
func main() {
	imgCrop("temp.png")
	// fmt.Printf("%v",redScreen())
	removeNonRed()
}