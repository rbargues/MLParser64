package main

import (
	"os"
	// "fmt"
	"image/png"
	"image/color"
)

func whiteScreen() bool {
	imageFile, err := os.Open("./screenshot.png")
	if err != nil {
		panic(err)
	}
	defer imageFile.Close()
	imageInfo, err := png.Decode(imageFile)
	if err != nil {
		panic(err)
	}
	bounds := imageInfo.Bounds()

	whitePixelCount := 0
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			rgba:= imageInfo.At(x,y).(color.RGBA)
			if rgba.R > 250 && rgba.B > 250 && rgba.G > 250 {
				whitePixelCount ++
			}
		}
	}
	if float64(whitePixelCount) / float64(bounds.Max.X * bounds.Max.Y) > 0.99 {
		return true
	} else {
		return false
	}
}
func blackScreen() bool{
	imageFile, err := os.Open("./screenshot.png")
	if err != nil {
		panic(err)
	}
	defer imageFile.Close()
	imageInfo, err := png.Decode(imageFile)
	if err != nil {
		panic(err)
	}
	bounds := imageInfo.Bounds()

	blackPixelCount := 0
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			rgba:= imageInfo.At(x,y).(color.RGBA)
			if rgba.R < 5 && rgba.B < 5 && rgba.G < 5 {
				blackPixelCount ++
			}
		}
	}
	if float64(blackPixelCount) / float64(bounds.Max.X * bounds.Max.Y) > 0.99 {
		return true
	} else {
		return false
	}
}
func redScreen(exitScreen chan bool) {
	imageFile, err := os.Open("./temp-cropped.png")
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
			// fmt.Printf("%v %v %v\n", rgba.R, rgba.G, rgba.B)
			if (rgba.R >= 185 && rgba.R <= 255) && (rgba.G >= 105 && rgba.G <= 115) && (rgba.B >= 105 && rgba.B <= 115) {
				redCt ++
			}
		}
	}
	if float64(redCt) / float64(bounds.Max.X * bounds.Max.Y) > 0.90 {
		exitScreen <- true
	} else {
		exitScreen <- false
	}
}