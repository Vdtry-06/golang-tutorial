package main

import (
	"fmt"
	"time"
)

func selectFunc() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- 2
	}()

	for i := 0; i < 2; i++ {
		select {
		case v1 := <-ch1:
			fmt.Println("Select: ", v1)
		case v2 := <-ch2:
			fmt.Println("Select: ", v2)
		}
	}
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // send
			x, y = y, x+y
		case <-quit: // receive
			fmt.Println("quit")
			return
		}
	}
}

func fiboProcess() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Fibo ",i + 1, ": ", <-c) // receive
		}
		quit <- 0 // send end loop
	}()
	fibonacci(c, quit)
}

func main() {
	selectFunc()
	fiboProcess()
}
