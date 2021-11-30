package main

import (
	"fmt"
)

func main() {
	repo := newFakeRepository()
	handler(repo)

	repo = NewRepository()
	handler(repo)
}

type repository interface {
	QueryAuthor(id uint) (*Author, error)
}

func handler(repo repository) {
	book, _ := repo.QueryAuthor(1)
	fmt.Println(book)
}

type fakeRepository struct{}

func (repo fakeRepository) QueryAuthor(uint) (*Author, error) {
	return &Author{
		ID:   99,
		Name: "Pallat",
	}, nil
}

func newFakeRepository() repository {
	return &fakeRepository{}
}

type Author struct {
	ID   uint
	Name string
}

func NewRepository() repository {
	return &Repository{db: store}
}

var store = map[uint]Author{
	1: {Name: "J. K. Rowling"},
	2: {Name: "J. R. R. Tolkien"},
	3: {Name: "George R. R. Martin"},
}

type Repository struct {
	db map[uint]Author
}

func (repo *Repository) QueryAuthor(id uint) (*Author, error) {
	if a, ok := repo.db[id]; ok {
		a.ID = id
		return &a, nil
	} else {
		return nil, fmt.Errorf("not found")
	}
}
