package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				log.Println("recovered ", err)
			} else {
				log.Printf("recovered %#v\n", r)
			}
		}
	}()

	n, err := strconv.Atoi("a")
	if err != nil {
		log.Panic("main ", err)
	}

	fmt.Println("number", n)
}
