package main

import "fmt"

func main() {
	var calculate = add

	fmt.Println(calculate(1, 2))

	m := Math{add: add}
	fmt.Println(m.add(1, 2))
}

type Math struct {
	add func(int, int) int
}

func add(a, b int) int {
	return a + b
}
