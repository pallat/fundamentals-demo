package book

const defaultAuthor = "Skooldio"

func NewBook(name, author string) Book {
	return Book{
		Name:   name,
		Author: author,
	}
}

type Book struct {
	Name   string
	Author string
}
