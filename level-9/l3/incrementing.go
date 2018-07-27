package main

import (
"fmt"
"sync"
"runtime"
)

var value int
var wg sync.WaitGroup

func increment() {
	curr_val := value
	curr_val += 1
	runtime.Gosched()

	value = curr_val
	wg.Done()
}

// Intentionally creating a race condition.
func main() {
	wg.Add(3)
	go increment()
	go increment()
	go increment()
	wg.Wait()
	fmt.Println(value)

}