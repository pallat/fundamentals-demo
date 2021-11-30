package main

import "fmt"

func main() {
	rangeOfChannel()
	rangeOfMap()
	rangeIndexValueOfArray()
	rangeIndexOfArray()
	rangeIndexOfSlice()
	rangeIndexValueOfSlice()
}

func rangeOfChannel() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func rangeOfMap() {
	m := map[string]string{
		"a": "apple",
		"b": "banana",
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func rangeIndexValueOfArray() {
	a := [...]int{1, 3, 5, 7, 9}
	for i, v := range a {
		fmt.Println(i, v)
	}
}

func rangeIndexOfArray() {
	a := [...]int{1, 3, 5, 7, 9}
	for i := range a {
		fmt.Println(a[i])
	}
}

func rangeIndexOfSlice() {
	s := []int{1, 3, 5, 7, 9}
	for i := range s {
		fmt.Println(s[i])
	}
}
func rangeIndexValueOfSlice() {
	s := []int{1, 3, 5, 7, 9}
	for i, v := range s {
		fmt.Println(i, v)
	}
}
