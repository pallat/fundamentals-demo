package main

import (
	"fmt"
	"strconv"
)

func emptyInterface() {
	var a interface{}

	a = 10
	a = "ten"
	a = true
	a = Book{}

	fmt.Println(a)
}

func signature() {
	var s Stringer
	// s = 1 // can not

	s = Int(1)

	s = Book{
		Name:   "The Lord of the Rings",
		Author: "J. R. R. Tolkien",
	}

	fmt.Println(s)
}

type Stringer interface {
	String() string
}

type String string

func (s String) String() string {
	return string(s)
}

type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

type Book struct {
	Name   string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("%s by %s", b.Name, b.Author)
}
