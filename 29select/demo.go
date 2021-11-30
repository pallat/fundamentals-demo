package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go slow(ch1, 100*time.Millisecond)
	go slow(ch2, 100*time.Millisecond)

	for i := 0; i < 10; i++ {
		select {
		case i := <-ch1:
			fmt.Println("channel #1 got ", i)
		case i := <-ch2:
			fmt.Println("channel #2 got ", i)
		}
	}
}

func slow(ch chan int, d time.Duration) {
	t := time.Tick(d)
	for range t {
		ch <- int(d)
	}
}
