package repository

import (
	"github.com/Rev1nant/go-book-db/internal/model"
	"github.com/Rev1nant/go-book-db/internal/repository/psql"
	"github.com/Rev1nant/go-book-db/pkg/db"
)

type Author interface {
	FindAll() ([]model.Author, error)
	FindByID(id int) (model.Author, error)
	Create(author model.Author) error
	Update(id int, author model.Author) error
	Delete(id int) error
	GetAuthorID(author model.Author) (int, error)
}

type Genre interface {
	FindAll() ([]model.Genre, error)
	FindByID(id int) (model.Genre, error)
	Create(genre model.Genre) error
	Update(id int, genre model.Genre) error
	Delete(id int) error
	GetGenreID(genreName string) (int, error)
}

type Book interface {
	FindAll() ([]model.Book, error)
	FindByID(id int) (model.Book, error)
	Create(book model.Book) error
	AddGenre(bookID, genreID int) error
	Update(book model.Book, authorID, bookID int) error
	UpdateGenre(bookGenreID, genreID int) error
	Delete(id int) error
	DeleteGenre(bookID, genreID int) error
	GetBookID(title string, author model.Author) (int, error)
}

type Repositories struct {
	Author Author
	Genre  Genre
	Book   Book
}

func NewRepositories(db *db.DB) *Repositories {
	return &Repositories{
		Author: psql.NewAuthorRepo(db),
		Genre:  psql.NewGenreRepo(db),
		Book:   psql.NewBookRepo(db),
	}
}
