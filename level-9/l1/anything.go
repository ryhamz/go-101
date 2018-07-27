package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("here's the main goroutine")
	wg.Add(2)
	go func() {
		for i := 0; i < 15; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	go func() {
		for i := 'a'; i < 'g'; i++ {
			fmt.Println(string(i))
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("NUM CPUs:", runtime.NumCPU())

}
