package service

import (
	"github.com/Rev1nant/go-book-db/internal/model"
	"github.com/Rev1nant/go-book-db/internal/repository"
)

type GenreService struct {
	repo repository.Repositories
}

func NewGenreService(repo *repository.Repositories) *GenreService {
	return &GenreService{
		repo: *repo,
	}
}

func (s *GenreService) GetAllGenre() ([]model.Genre, error) {
	return s.repo.Genre.FindAll()
}

func (s *GenreService) GetOneGenre(id int) (model.Genre, error) {
	return s.repo.Genre.FindByID(id)
}

func (s *GenreService) AddGenre(genre model.Genre) error {
	return s.repo.Genre.Create(genre)
}

func (s *GenreService) UpdateGenre(id int, genre model.Genre) error {
	return s.repo.Genre.Update(id, genre)
}

func (s *GenreService) DeleteGenre(id int) error {
	return s.repo.Genre.Delete(id)
}
