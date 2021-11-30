package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for {

	// }

	fmt.Println(factorial(5))
}

func factorial(n int) int {
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	return sum
}
