package main

import (
	"fmt"
)

func main() {
	ch := demoOneWay()
	fmt.Println(<-ch)
}

func demoOneWay() <-chan string {
	ch := make(chan string)
	go demoWriteChannel(ch, "John")
	return ch
}

func demoWriteChannel(w chan<- string, name string) {
	w <- "Hello, " + name
}

func demo2ways() {
	rw := make(chan string)
	go greeting(rw)

	rw <- "John"
	fmt.Println(<-rw)
}

func greeting(rw chan string) {
	s := <-rw
	rw <- "Hello, " + s
}
