package main

import (
	"fmt"
	"strings"
)

func slice() {
	var fiveNum [5]int
	var nums []int

	nums = fiveNum[1:3]

	_ = len(nums)
	_ = cap(nums)

	{
		nums := []int{
			1,
			2,
			3,
		}
		nums = append(nums, 4, 5)
	}

	_ = nums
}

func main() {
	fmt.Printf("%q\n", variadic())
}

// func demo() {
// 	slice := make([]int, 0)
// 	slice := []int{}
// }

func variadic(v ...string) string {
	fmt.Printf("%T %#v\n", v, v)

	return strings.Join(v, ",")
}

func array() {
	var fourNum [4]int
	fourNum[0] = 1
	fourNum[2] = 3
}

func makeSlice() {
	var nums []int

	nums = make([]int, 4)

	nums[0] = 1
	nums[2] = 2

	nums = append(nums, 20)
}
