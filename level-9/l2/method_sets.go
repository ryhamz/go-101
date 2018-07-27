package main

import (
"fmt"
)

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hi")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first: "Elliot",
	}
	saySomething(&p)
}