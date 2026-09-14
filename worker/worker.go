package worker

import (
	"errors"
	"fmt"

	"yourgo/model"
)

// Интерфейс прописываем там, где он используется.
// Метода Delete здесь нет: worker его не использует, значит он интерфейсу не нужен.
// Имя с маленькой буквы - интерфейс неэкспортируем, вне этого пакета он не нужен.
type repo interface {
	ByID(ID int) (model.Book, error)
	Add(b model.Book) error
}

type Worker struct {
	repo repo
}

func New(repo repo) (*Worker, error) {
	if repo == nil {
		return nil, errors.New("repo is nil")
	}
	return &Worker{
		repo: repo,
	}, nil
}

func (w Worker) CreateBooks() error {
	b1, err := model.NewBook(5, "Война и мир", "Лев Толстой")
	if err != nil {
		return fmt.Errorf("create new book: %w", err)
	}
	if err := w.repo.Add(b1); err != nil {
		return fmt.Errorf("add book with id %d in repository: %w", b1.ID, err)
	}

	b2, err := model.NewBook(10, "Преступление и наказание", "Фёдор Достоевский")
	if err != nil {
		return fmt.Errorf("create new book: %w", err)
	}
	if err := w.repo.Add(b2); err != nil {
		return fmt.Errorf("add book with id %d in repository: %w", b2.ID, err)
	}

	return nil
}

func (w Worker) PrintBook(ID int) error {
	b, err := w.repo.ByID(ID)
	if err != nil {
		return fmt.Errorf("get book by id %d from repository: %w", ID, err)
	}
	fmt.Printf("%+v\n", b)
	return nil
}
