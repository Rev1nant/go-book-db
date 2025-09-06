package service

import (
	"github.com/Rev1nant/go-book-db/internal/model"
	"github.com/Rev1nant/go-book-db/internal/repository"
)

type AuthorService struct {
	repo repository.Repositories
}

func NewAuthorService(repo *repository.Repositories) *AuthorService {
	return &AuthorService{
		repo: *repo,
	}
}

func (s *AuthorService) GetAllAuthor() ([]model.Author, error) {
	return s.repo.Author.FindAll()
}

func (s *AuthorService) GetOneAuthor(id int) (model.Author, error) {
	return s.repo.Author.FindByID(id)
}

func (s *AuthorService) AddAuthor(author model.Author) error {
	return s.repo.Author.Create(author)
}

func (s *AuthorService) UpdateAuthor(id int, author model.Author) error {
	return s.repo.Author.Update(id, author)
}

func (s *AuthorService) DeleteAuthor(id int) error {
	return s.repo.Author.Delete(id)
}

func (s *AuthorService) GetAuthorID(author model.Author) (int, error) {
	return s.repo.Author.GetAuthorID(author)
}
