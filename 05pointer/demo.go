package main

import "fmt"

func main() {
	demo()
}

func demo() {
	var s string
	var p *string

	p = &s
	s = "hello"

	fmt.Printf("s = %q\n", s)

	*p += " world"

	fmt.Printf("s = %q\n", s)

}
