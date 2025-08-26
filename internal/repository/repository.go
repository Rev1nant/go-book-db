package repository

import (
	"github.com/Rev1nant/go-book-db/internal/model"
	"github.com/Rev1nant/go-book-db/internal/repository/db"
	"github.com/Rev1nant/go-book-db/pkg/datebase"
)

type Author interface {
	FindAll() ([]model.Author, error)
	FindByID(id int) (model.Author, error)
	Create(author model.Author) error
	Update(id int, author model.Author) error
	Delete(id int) error
}

type Genre interface {
	FindAll() ([]model.Genre, error)
	FindByID(id int) (model.Genre, error)
	Create(genre model.Genre) error
	Update(id int, genre model.Genre) error
	Delete(id int) error
}

type Book interface {
	FindAll() ([]model.Book, error)
	FindByID(id int) (model.Book, error)
	Create(book model.Book, authorID int) error
	AddGenre(bookID, genreID int) error
	Update(book model.Book, authorID, bookID int) error
	UpdateGenre(bookGenreID, genreID int) error
	Delete(id int) error
	DeleteGenre(bookGenreID int) error
}

type Repositories struct {
	Author Author
	Genre  Genre
	Book   Book
}

func NewRepositories(database *datebase.DB) *Repositories {
	return &Repositories{
		Author: db.NewAuthorRepo(database),
		Genre:  db.NewGenreRepo(database),
		Book:   db.NewBookRepo(database),
	}
}
