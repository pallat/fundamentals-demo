package main

import (
	"fmt"
	"strconv"
)

func main() {
	greeting("")
	greeting("Pallat")

	if n, err := strconv.Atoi("5s"); err != nil {
		fmt.Println("NaN:", err)
	} else {
		fmt.Println(n)
	}
}

func greeting(name string) {
	if name != "" {
		fmt.Printf("Hello %s\n", name)
	} else {
		fmt.Println("Hello friend")
	}
}
