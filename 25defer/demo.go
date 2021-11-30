package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	tricky()
}

func usecase(r io.ReadCloser) {
	defer r.Close()
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(b))
}

func tricky() {
	for i := 0; i < 3; i++ {
		defer func() { fmt.Println(i) }()
	}
}

func dontForget(n int) {
	defer fmt.Println(n)
	defer func() {
		fmt.Println(n)
	}()
	n += n
}
