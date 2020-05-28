package main

import (
	"os"
	"fmt"
	"image"
	"image/png"
	"image/color"
)

func whiteScreen(fname string) bool {
	imageFile, err := os.Open(fname)
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
func blackScreen(fname string) bool{
	imageFile, err := os.Open(fname)
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
			} else if rgba.B != 0 {
				m.Set(x, y, imgColor)
			} else {
				m.Set(x, y, rgba)
			}
		}
	}
	newImg, _ := os.Create(output)
	defer newImg.Close()
	png.Encode(newImg, m)
}
func compareLifeCounts (img1 string, img2 string, img3 string) {
	dct1 := obtainDCT(img1)
	dct2 := obtainDCT(img2)
	dct3 := obtainDCT(img3)

	phashVal1 := phash(dct1)
	phashVal2 := phash(dct2)
	phashVal3 := phash(dct3)

	hamming1 := hammingDistance(phashVal2, phashVal1)
	hamming2 := hammingDistance(phashVal2, phashVal3)

	if hamming1 < hamming2 {
        fmt.Printf("more similar to first %v %v\n", hamming1, hamming2)
		fmt.Printf("%v\n", "life count changed")
	} else {
        fmt.Printf("more similar to second %v %v\n", hamming1, hamming2)
		fmt.Printf("%v\n", "life count is the same")
	}
}
func blackScreenLife(fname string) bool{
	imageFile, err := os.Open(fname)
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
	for x := bounds.Min.X; x < 5; x++ {
		for y := bounds.Min.Y; y < 5; y++ {
			rgba:= imageInfo.At(x,y).(color.RGBA)
			if rgba.R < 60 && rgba.B < 60 && rgba.G < 60 {
				blackPixelCount ++
			}
		}
	}
	if float64(blackPixelCount) / float64(25) > 0.99 {
		return true
	} else {
		return false
	}
}