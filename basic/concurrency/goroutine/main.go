package main

import (
	"fmt"
	"runtime"
	"time"
)

// Goroutine
func RoutineFunc() {
	go func() {
		for i := 0; i < 50; i++ {
			fmt.Println("I'm trying goroutine 1: ", i)
			runtime.Gosched()
		}
	}()

	go func() {
		for i := 0; i < 50; i++ {
			fmt.Println("I'm trying goroutine 2: ", i)
			runtime.Gosched()
		}
	}()

	time.Sleep(time.Second)
}

func main() {
	fmt.Println("goroutine")
	RoutineFunc()
}
