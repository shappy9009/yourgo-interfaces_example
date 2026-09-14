package repository

import (
	"fmt"

	"yourgo/model"
)

// SecondMemoryRepository - имитация другого репозитория:
// делает то же самое, но, условимся, сохраняет данные по-своему, в другое место.
type SecondMemoryRepository struct {
	books map[int]model.Book
}

func NewSecondMemoryRepository() *SecondMemoryRepository {
	return &SecondMemoryRepository{
		books: make(map[int]model.Book),
	}
}

func (r SecondMemoryRepository) ByID(ID int) (model.Book, error) {
	book, exists := r.books[ID]
	if !exists {
		return model.Book{}, fmt.Errorf("book with id %d not found", ID)
	}
	return book, nil
}

func (r SecondMemoryRepository) Add(b model.Book) error {
	if _, exists := r.books[b.ID]; exists {
		return fmt.Errorf("book with id %d already exists", b.ID)
	}
	r.books[b.ID] = b
	return nil
}

func (r SecondMemoryRepository) Delete(ID int) error {
	if _, exists := r.books[ID]; !exists {
		return fmt.Errorf("book with id %d not found", ID)
	}
	delete(r.books, ID)
	return nil
}
