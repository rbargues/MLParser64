package main

import (
	"fmt"
    "time"
    "math"
)

//takes 25 seconds from last bowser star until final black screen
func main() {
    //starts the timer
    timer := make(chan time.Time)
    go startgame(timer)
    timeVal := <- timer

    var r1 [2]int
    r1[0], r1[1] = 65, 80
    var r2 [2]int
    r2[0], r2[1] = 1080, 800
    jsonFile := readJSON()
    whiteScreenCt := 0
    blackScreenCt := 0

    var startTime time.Time
    whiteBool := false
    blackBool := false
    for true {
        screenshotGrab(r1, r2)
        if whiteScreen() {
            blackScreenCt = 0
            whiteScreenCt ++
        } else if blackScreen() {
            blackScreenCt ++
        }
        if whiteScreenCt == 1 && blackScreenCt == 1 {
            startTime = time.Now()
            whiteBool = true
            time.Sleep(5 * time.Second)
            screenshotGrab(r1, r2)
        } else if blackScreenCt === 3 && whiteScreenCt == 0 {
            startTime = time.Now()
            blackBool = true
            time.Sleep(5 * time.Second)
            screenshotGrab(r1, r2)
        }
        if whiteBool == true {
            whiteBool = false
            whiteScreenCt = 0
            blackScreenCt = 0
            whiteImgCrop("screenshot.png")
            dctArr := obtainDCT("cropped.png")
            phashVal := phash(dctArr)
            minDist := math.Inf(0)
            for key, value := range jsonFile {
                hamming := hammingDistance(phashVal, value)
                if float64(hamming) < minDist {
                    minDist = float64(hamming)
                    currentLevel = key
                }
            }

        } else if blackBool == true {
            blackBool = false
            whiteScreenCt = 0
            blackScreenCt = 0
            blackImgCrop("screenshot.png")
            dctArr := obtainDCT("cropped.png")
            phashVal := phash(dctArr)
            minDist := math.Inf(0)
            for key, value := range jsonFile {
                hamming := hammingDistance(phashVal, value)
                if float64(hamming) < minDist {
                    minDist = float64(hamming)
                    currentLevel = key
                }
            }
        }



            dctArr := obtainDCT("cropped.png")
            phashVal := phash(dctArr)
            minDist := math.Inf(0)
            currentLevel := ""
            for key, value := range jsonFile {
                hamming := hammingDistance(phashVal, value)
                if float64(hamming) < minDist {
                    minDist = float64(hamming)
                    currentLevel = key
                }
            }
 
            fmt.Printf("%v\n", currentLevel)
        } else if blackScreenBool == true {
            blackScreenCt ++
        }
    }
}