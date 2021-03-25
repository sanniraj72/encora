package main

import (
	"log"
	"sync"
	"time"
)

var mutex sync.Mutex

func main() {

	for v := 0; v < 10; v++ {
		go func(v int) {
			doublev := callDouble(v)
			log.Printf("Thread %d returned: %d", v, doublev)
		}(v)
	}

	time.Sleep(time.Second * 10)
}

func callDouble(v int) int {
	// Adjust code to call double only up to 5 times concurrently
	throttle := time.Tick(time.Second * 5)
	for {
		<-throttle
		mutex.Lock()
		dVal := double(v)
		mutex.Unlock()
		return dVal
	}
}

func double(v int) int {
	time.Sleep(time.Second)
	return v * 2
}
