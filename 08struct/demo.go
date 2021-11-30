package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Book struct {
	Name   string
	Author string
}

func NewBook(jsonString string) Book {
	r := strings.NewReader(jsonString)
	var book Book
	if err := json.NewDecoder(r).Decode(&book); err != nil {
		log.Fatal(err)
	}
	return book
}

func (book Book) Json() (string, error) {
	b, err := json.Marshal(&book)
	return string(b), err
}

func main() {
	jsonString := `{
		"name": "Harry Potter and the Sorcerer's Stone",
		"author": "J. K. Rowling"
	}`

	book := NewBook(jsonString)
	fmt.Println(book)
	fmt.Println(book.Json())
}
