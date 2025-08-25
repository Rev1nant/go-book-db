package service

import (
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

// func (s *GenreService) GetAllGenre(res http.Response, req *http.Request) ([]domain.Genre, error) {
// 	genre
// }
