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

    // time.Sleep(30 * time.Second)
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
    exitScreen := make(chan bool)
    for true {
        screenshotGrab(r1, r2, "screenshot.png")
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
            if whiteScreenCt > 0 || blackScreenCt > 1 {
                go func () {
                    time.Sleep(1100 * time.Millisecond)
                    screenshotGrab(r1, r2, "temp.png")
                    exitCourseCrop("temp.png")
                    redScreen(exitScreen)
                    // fmt.Printf("%v\n",redScreen(exitScreen))
                }()
            }
            blackScreenCt ++ 
            if blackScreenCt == 1 && whiteScreenCt == 0{
                startTime = time.Now()
            } else if whiteScreenCt == 0 && blackScreenCt < 3 {
                time.Sleep(1 * time.Second)
            } else if whiteScreenCt == 2 && blackScreenCt == 1 {
                endTime = time.Now()
                whiteBool = true
                time.Sleep(5 * time.Second)
                screenshotGrab(r1, r2, "screenshot.png")
            } else if blackScreenCt == 3 && whiteScreenCt == 0 {
                endTime = time.Now()
                blackBool = true
                time.Sleep(5 * time.Second)
                screenshotGrab(r1, r2, "screenshot.png")
            }
            exitBool := <- exitScreen
            if exitBool {
                fmt.Println("exited course\n")
                whiteScreenCt = 0
                blackScreenCt = 0
                whiteBool = false
                blackBool = false
            }  
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
                if float64(hamming) < minDist {
                    minDist = float64(hamming)
                    currentLevel = key
                }
            }
            if currentLevel == "bowser3" {
                fmt.Printf("Full Time: %v\n", endTime.Sub(timeVal.Add(-25 * time.Second)))
            }
            fmt.Printf("%v: %v\n", currentLevel, endTime.Sub(startTime))
        }           
    }
}