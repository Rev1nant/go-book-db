package service

import "github.com/Rev1nant/go-book-db/internal/repository"

type BookService struct {
	repo repository.Repositories
}

func NewBookService(repo *repository.Repositories) *BookService {
	return &BookService{
		repo: *repo,
	}
}
