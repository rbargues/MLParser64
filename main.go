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
    fmt.Printf("%v\n", timeVal)

    var r1 [2]int
    r1[0], r1[1] = 65, 80
    var r2 [2]int
    r2[0], r2[1] = 1080, 800
    jsonFile := readJSON("levels.json")
    bowserFile := readJSON("bowsers.json")
    whiteScreenCt := 0
    blackScreenCt := 0

    var startTime time.Time
    var endTime time.Time
    whiteBool := false
    blackBool := false
    for true {
        screenshotGrab(r1, r2)
        if whiteScreen() {
            if whiteScreenCt == 1 {
                startTime = time.Now()
                time.Sleep(1000 * time.Millisecond)
            } else if whiteScreenCt > 2 {
                whiteScreenCt = 0
            }
            blackScreenCt = 0
            whiteScreenCt ++
            time.Sleep(400 * time.Millisecond)
        } else if blackScreen() {
            if blackScreenCt == 0 && whiteScreenCt == 0{
                startTime = time.Now()
            }
            if whiteScreenCt == 0 && blackScreenCt < 2 {
                time.Sleep(1 * time.Second)
            }
            blackScreenCt ++   
        }
        if whiteScreenCt == 2 && blackScreenCt == 1 {
            endTime = time.Now()
            whiteBool = true
            time.Sleep(5 * time.Second)
            screenshotGrab(r1, r2)
        } else if blackScreenCt == 3 && whiteScreenCt == 0 {
            endTime = time.Now()
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
            currentLevel := ""
            for key, value := range jsonFile {
                hamming := hammingDistance(phashVal, value)
                if float64(hamming) < minDist {
                    minDist = float64(hamming)
                    currentLevel = key
                }
            }
            fmt.Printf("%v: %v\n", currentLevel, endTime.Sub(startTime))

        } else if blackBool == true {
            blackBool = false
            whiteScreenCt = 0
            blackScreenCt = 0
            blackImgCrop("screenshot.png")
            dctArr := obtainDCT("cropped.png")
            phashVal := phash(dctArr)
            minDist := math.Inf(0)
            currentLevel := ""
            for key, value := range bowserFile {
                hamming := hammingDistance(phashVal, value)
                fmt.Printf("%v, %v\n", key, hamming)
                if hamming >= 200 {
                    continue
                }
                if float64(hamming) < minDist {
                    minDist = float64(hamming)
                    currentLevel = key
                }
            }
            fmt.Printf("%v: %v\n", currentLevel, endTime.Sub(startTime))
        }           
    }
}