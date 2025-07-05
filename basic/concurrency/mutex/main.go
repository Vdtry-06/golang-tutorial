package main

import (
	"fmt"
	"sync"
	"time"
)

func dataRacing() {
	lock := new(sync.Mutex)
	cnt := 0
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				lock.Lock()
				cnt++ // 3 step: get cnt, increase cnt, save cnt => conflict. resolve: use Mutex
				fmt.Println(cnt)
				lock.Unlock()
			}
		}()
	}

	time.Sleep(time.Second)
}

func main() {
	dataRacing()
}
