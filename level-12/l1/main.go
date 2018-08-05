package main

import (
	"fmt"
	"go-101/level-12/l1/dog"
)

func main() {
	fmt.Println("let's talk about a dog.")
	age := dog.Years(1)
	fmt.Println(age)
}
