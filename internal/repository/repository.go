package repository

import (
	"github.com/Rev1nant/go-book-db/internal/model"
	"github.com/Rev1nant/go-book-db/internal/repository/db"
	"github.com/Rev1nant/go-book-db/pkg/datebase"
)

type Author interface {
	GetAllAuthor() ([]model.Author, error)
	GetOneAuthor(id int) (model.Author, error)
	AddAuthor(author model.Author) error
	UpdateAuthor(id int, author model.Author) error
	DeleteAuthor(id int) error
}

type Genre interface {
	GetAllGenre() ([]model.Genre, error)
	GetOneGenre(id int) (model.Genre, error)
	AddGenre(genre model.Genre) error
	UpdateGenre(id int, genre model.Genre) error
	DeleteGenre(id int) error
}

type Book interface {
	GetAllBook() ([]model.Book, error)
	GetOneBook(id int) (model.Book, error)
	AddBook(book model.Book, authorID int) error
	AddBookGenre(bookID, genreID int) error
	UpdateBook(book model.Book, authorID, bookID int) error
	UpdateBookGenre(bookGenreID, genreID int) error
	DeleteBook(id int) error
	DeleteBookGenre(bookGenreID int) error
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
