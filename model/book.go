package model

import (
	"errors"
	"fmt"
)

type Book struct {
	ID     int
	Title  string
	Author string
}

func NewBook(ID int, title string, author string) (Book, error) {
	if ID <= 0 {
		return Book{}, errors.New("ID must be greater 0")
	}
	if title == "" {
		return Book{}, errors.New("title cannot be empty")
	}
	if author == "" {
		return Book{}, errors.New("author cannot be empty")
	}
	return Book{
		ID:     ID,
		Title:  title,
		Author: author,
	}, nil
}

func (b Book) String() string {
	return fmt.Sprintf("ID: %d, Название: %s, Автор: %s", b.ID, b.Title, b.Author)
}
