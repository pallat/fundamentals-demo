package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("start")
	go doSomeThing()
	time.Sleep(time.Second)
	fmt.Println("end")
}

func doSomeThing() {
	fmt.Println("do something")
}

func sewing() {
	ch := make(chan struct{})
	go func() {
		for range ch {
			fmt.Print("-")
		}
	}()

	go func() {
		t1 := time.Tick(time.Millisecond * 100)
		t2 := time.Tick(time.Millisecond * 200)
		for {
			select {
			case <-t1:
				fmt.Print("x")
			case <-t2:
				ch <- struct{}{}
			}
		}
	}()

	time.Sleep(5 * time.Second)
}

func train() {
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Print("-")
		}
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Print("*")
		}
	}()

	time.Sleep(10 * time.Second)
}

func mainfibo() {
	times := 15
	limit := make(chan struct{}, times)

	for i := 0; i < times; i++ {
		limit <- struct{}{}
	}
	close(limit)

	ch := make(chan int)

	go fibonacci(ch, limit)

	for i := range ch {
		fmt.Println(i)
	}
}

func fibonacci(ch chan int, limit chan struct{}) {
	a, b := 0, 1
	for range limit {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func pingpong() {
	ch := make(chan int)
	finish := make(chan struct{})
	go pong(ch, finish)
	go ping(ch)

	<-finish
}

func ping(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Fprintf(os.Stdout, "ping(send): %d\n", i)
		os.Stdout.WriteString("your message there")
	}
	close(ch)
}

func pong(ch chan int, finish chan struct{}) {
	for i := range ch {
		fmt.Fprintf(os.Stderr, "%30d :pong(receive)\n", i)
		os.Stderr.WriteString("your message here")
	}
	finish <- struct{}{}
}
