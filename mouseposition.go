package main

import (
	"github.com/go-vgo/robotgo"
)

func mousePosition() <-chan []int{
	r := make(chan []int)
	go func() {
		mleft := robotgo.AddEvent("mleft")
		if mleft {
			store := make([]int, 0)
			x, y := robotgo.GetMousePos()	
			store = append(store, x, y)
			r <- store
		}
	}()
	return r
}
