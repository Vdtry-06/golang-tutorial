package main

import (
	"fmt"
	"sync"
	"time"
)

// Topic: 10000 url web, need crawl html return, maximum 5 worker, (= concurrency 5 page crawl)
// FIFO

func main() {
	n := 10000
	maxWorker := 5

	wg := new(sync.WaitGroup)

	queueCh := make(chan int, n)

	for i := 1; i <= n; i++ {
		queueCh <- i // send
	}

	close(queueCh)

	for i := 1; i <= maxWorker; i++ {
		wg.Add(1) // ++
		go func(count int) {
			for v := range queueCh {
				time.Sleep(time.Millisecond * 5)
				// fmt.Printf("Worker %d is crawling web url %d\n", i, <-queueCh)
				// i is not execute. Running, i = 1 -> 5 then i made incorrectly. Resolve: add variable: count
				fmt.Printf("Worker %d is crawling web url %d\n", count, v)
			}
			wg.Done() // -- = 0 
		}(i)
	}

	wg.Wait() // wait when all wg.Done
}
