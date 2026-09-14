package main

import (
	"fmt"

	"yourgo/repository"
	"yourgo/worker"
)

func main() {
	// Внедрение зависимостей (Dependency Injection):
	// создаем репозиторий и передаем его воркеру снаружи.
	repo := repository.NewBookInMemoryRepository()
	wrkr, err := worker.New(repo)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := wrkr.CreateBooks(); err != nil {
		fmt.Println(err)
		return
	}
	if err := wrkr.PrintBook(5); err != nil {
		fmt.Println(err)
		return
	}

	// Ключевой момент: воркер работает через интерфейс,
	// поэтому тот же самый worker можно использовать
	// и с другим типом репозитория.
	repo2 := repository.NewSecondMemoryRepository()
	wrkr2, err := worker.New(repo2)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := wrkr2.CreateBooks(); err != nil {
		fmt.Println(err)
		return
	}
	if err := wrkr2.PrintBook(10); err != nil {
		fmt.Println(err)
		return
	}
}
