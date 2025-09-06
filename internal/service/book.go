package service

import (
	"github.com/Rev1nant/go-book-db/internal/model"
	"github.com/Rev1nant/go-book-db/internal/repository"
)

type BookService struct {
	repo repository.Repositories
}

func NewBookService(repo *repository.Repositories) *BookService {
	return &BookService{
		repo: *repo,
	}
}

func (s *BookService) GetAllBook() ([]model.Book, error) {
	return s.repo.Book.FindAll()
}

func (s *BookService) GetOneBook(id int) (model.Book, error) {
	return s.repo.Book.FindByID(id)
}

func (s *BookService) AddBook(book model.Book) error {
	return s.repo.Book.Create(book)
}

func (s *BookService) AddBookGenre(bookID, genreID int) error {
	return s.repo.Book.AddGenre(bookID, genreID)
}

func (s *BookService) UpdateBook(book model.Book, authorID, bookID int) error {
	return s.repo.Book.Update(book, authorID, bookID)
}

func (s *BookService) UpdateBookGenre(bookGenreID, genreID int) error {
	return s.repo.Book.UpdateGenre(bookGenreID, genreID)
}

func (s *BookService) DeleteBook(id int) error {
	return s.repo.Book.Delete(id)
}

func (s *BookService) DeleteBookGenre(bookID, genreID int) error {
	return s.repo.Book.DeleteGenre(bookID, genreID)
}

func (s *BookService) GetBookID(title string, author model.Author) (int, error) {
	return s.repo.Book.GetBookID(title, author)
}
