package main

import (
	"fmt"
    "time"
    "github.com/go-vgo/robotgo"
)

func startgame(timer chan time.Time) {
    buttonPress := robotgo.AddEvent("rshift")
    if buttonPress {
        timer <- time.Now()
    }
}