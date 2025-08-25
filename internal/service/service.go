package service

import (
	"github.com/Rev1nant/go-book-db/internal/repository"
)

type Author interface{}

type Genre interface{}

type Book interface{}

type Services struct {
	Author Author
	Genre  Genre
	Book   Book
}

func NewServices(repo *repository.Repositories) *Services {
	return &Services{
		Author: NewAuthorService(repo),
		Genre:  NewGenreService(repo),
		Book:   NewBookService(repo),
	}
}
