package main

import "fmt"

func main() {
	// m := make(map[string]string)
	m := map[string]string{}

	m["a"] = "apple"

	fmt.Println(m)

	m = map[string]string{
		"a": "apple",
		"b": "banana",
	}
	fmt.Println(m)
}
