package repository

import (
	"github.com/Rev1nant/go-book-db/internal/domain"
	"github.com/Rev1nant/go-book-db/internal/repository/db"
	"github.com/Rev1nant/go-book-db/pkg/datebase"
)

type Author interface {
	GetAllAuthor() ([]domain.Author, error)
	GetOneAuthor(id int) (domain.Author, error)
	AddAuthor(firstName, lastName string) error
	UpdateAuthor(id int, firstNameNew, lastNameNew string) error
	DeleteAuthor(id int) error
}

type Genre interface {
	GetAllGenre() ([]domain.Genre, error)
	GetOneGenre(id int) (domain.Genre, error)
	AddGenre(genreName string) error
	UpdateGenre(id int, genreNameNew string) error
	DeleteGenre(id int) error
}

type Book interface {
	GetAllBook() ([]domain.Book, error)
	GetOneBook(id int) (domain.Book, error)
	AddBook(title string, authorID int, price float64, amount int) error
	AddBookGenre(bookID, genreID int) error
	UpdateBook(id int, title string, authorID int, price float64, amount int) error
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
