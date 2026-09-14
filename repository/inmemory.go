package repository

import (
	"fmt"

	"yourgo/model"
)

type BookInMemoryRepository struct {
	books map[int]model.Book
}

func NewBookInMemoryRepository() *BookInMemoryRepository {
	return &BookInMemoryRepository{
		books: make(map[int]model.Book),
	}
}

func (r BookInMemoryRepository) ByID(ID int) (model.Book, error) {
	book, exists := r.books[ID]
	if !exists {
		return model.Book{}, fmt.Errorf("book with id %d not found", ID)
	}
	return book, nil
}

func (r BookInMemoryRepository) Add(b model.Book) error {
	if _, exists := r.books[b.ID]; exists {
		return fmt.Errorf("book with id %d already exists", b.ID)
	}
	r.books[b.ID] = b
	return nil
}

func (r BookInMemoryRepository) Delete(ID int) error {
	if _, exists := r.books[ID]; !exists {
		return fmt.Errorf("book with id %d not found", ID)
	}
	delete(r.books, ID)
	return nil
}
