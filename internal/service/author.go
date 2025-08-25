package service

import "github.com/Rev1nant/go-book-db/internal/repository"

type AuthorService struct {
	repo repository.Repositories
}

func NewAuthorService(repo *repository.Repositories) *AuthorService {
	return &AuthorService{
		repo: *repo,
	}
}
