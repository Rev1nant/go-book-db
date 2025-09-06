package service

import (
	"github.com/Rev1nant/go-book-db/internal/model"
	"github.com/Rev1nant/go-book-db/internal/repository"
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
	GetGenreID(genreName string) (int, error)
}

type Book interface {
	GetAllBook() ([]model.Book, error)
	GetOneBook(id int) (model.Book, error)
	AddBook(book model.Book) error
	AddBookGenre(bookID, genreID int) error
	UpdateBook(book model.Book, authorID, bookID int) error
	UpdateBookGenre(bookGenreID, genreID int) error
	DeleteBook(id int) error
	DeleteBookGenre(bookGenreID int) error
	GetBookID(title string, author model.Author) (int, error)
}

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
