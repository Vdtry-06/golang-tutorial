package main

import (
	"fmt"
	"time"
)

// Channel
func ChanFunc() {
	myChan := make(chan int)
	go func() {
		myChan <- 1 // send data
	}()

	res := <-myChan // receive data
	fmt.Println("Received ChanFunc: ", res)
}

// Send and Receive
// Case 1
func ReceiveAndSend(c chan int) {
	fmt.Printf("Received: %d\n", <-c)
	fmt.Printf("Sending 2...\n")
	c <- 2
}

func ReceiveAndSendProcess() {
	myChan := make(chan int)
	go ReceiveAndSend(myChan)
	myChan <- 1
	fmt.Printf("Value from received: %d\n", <-myChan)
}

// Receive
func ReceiveOnly(c <-chan int) {
	fmt.Printf("Received: %d\n", <-c)
}

// Send
func SendOnly(c chan<- int) {
	fmt.Printf("Sending 2...\n")
	c <- 2
}

func ReceiveAndSendProcess2() {
	myChan := make(chan int)

	go ReceiveOnly(myChan)
	myChan <- 1 // send

	go SendOnly(myChan)
	res := <-myChan // receive

	fmt.Printf("Value from received: %d\n", res)
}

// streaming
func StreamingFunc() {
	myChan := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			myChan <- i
			time.Sleep(time.Second / 2)
		}
	}()

	for i := 0; i < 10; i++ {
		res := <-myChan
		fmt.Println(res)
	}
}

func main() {
	fmt.Println("Channel")
	ChanFunc()
	ReceiveAndSendProcess()
	ReceiveAndSendProcess2()
	StreamingFunc()
}
