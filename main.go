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
    var grabLevelName bool
    var startTime time.Time
    for true {
        screenshotGrab(r1, r2)
        if whiteScreen() {
            whiteScreenCt ++
        } else if blackScreen() {
            blackScreenCt ++
        }
        if whiteScreenCt == 1 || (blackScreenCt == 1 && whiteScreenCt == 0) {
            startTime = time.Now()
        } else if (whiteScreenCt == 1 && blackScreenCt == 1) || (blackScreenCt == 3 && whiteScreenCt == 0) {
            // will need to handle levels where there are only 2 black screens
            // princess secret slide etc
            // maybe have dct on static portion of level
            whiteScreenCt = 0
            blackScreenCt = 0
            endTime := time.Since(startTime)
            fmt.Printf("%v\n", endTime)
            time.Sleep(5 * time.Second)
        }

        // blackScreenBool := blackScreen()
        // if blackScreenBool == true && whiteScreenCt == 1 {
        //     // should be no instance where we have black then white again
        //     whiteScreenCt = 0 
        //     time.Sleep(5 * time.Second)
        //     screenshotGrab(r1, r2)
        //     imgCrop("screenshot.png")
        //     dctArr := obtainDCT("cropped.png")
        //     phashVal := phash(dctArr)
        //     minDist := math.Inf(0)
        //     currentLevel := ""
        //     for key, value := range jsonFile {
        //         hamming := hammingDistance(phashVal, value)
        //         if float64(hamming) < minDist {
        //             minDist = float64(hamming)
        //             currentLevel = key
        //         }
        //     }
 
        //     fmt.Printf("%v\n", currentLevel)
        // } else if blackScreenBool == true {
        //     blackScreenCt ++
        // }
    }
}