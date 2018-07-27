package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var value int64
var wg sync.WaitGroup

func increment() {
	atomic.AddInt64(&value, 1)
	wg.Done()
}

// Fixed the race condition.
func main() {
	wg.Add(3)
	go increment()
	go increment()
	go increment()
	wg.Wait()
	fmt.Println(value)

}
