package main

import "fmt"

func main() {
	fmt.Println(HOFGreeting(newNameFunc("John")))
}

func newNameFunc(name string) nameFunc {
	return func() string { return name }
}

func parameters() {
	nameFn := func() string { return "John" }
	s := HOFGreeting(nameFn)

	fmt.Println(s)
}

type nameFunc func() string

func HOFGreeting(nameFn nameFunc) string {
	return fmt.Sprintf("hello, %s", nameFn())
}
