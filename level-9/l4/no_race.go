package main

import (
"fmt"
"sync")

var value int
var wg sync.WaitGroup
var m sync.Mutex

func increment() {
	
	m.Lock()
	curr_val := value
	curr_val += 1

	value = curr_val
	m.Unlock()
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