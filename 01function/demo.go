package main

import "fmt"

func main() {
	greeting("Pallat", "Anchaleechamaikorn")

	result := add(4, 10)
	fmt.Println(result)

	a, b := 11, 22
	a, b = swap(a, b)
	fmt.Println(a, b)
}

func greeting(firstName, lasstName string) {
	fmt.Println("Hello", firstName, lasstName)
}

func add(a, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}
