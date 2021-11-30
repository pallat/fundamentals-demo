package main

import (
	"fmt"

	"github.com/pallat/bookstore/book"
)

func main() {
	bk := book.NewBook(
		"The Lord of the Rings",
		"J. R. R. Tolkien",
	)

	fmt.Printf("%#v\n", bk)
}
