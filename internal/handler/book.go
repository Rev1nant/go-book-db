package handler

import (
	"net/http"

	"github.com/Rev1nant/go-book-db/internal/service"
)

type BookHandlers struct {
	BookService service.BookService
}

func NewBookHandler(router *http.ServeMux, deps BookHandlers) {
	handler := &BookHandlers{
		BookService: deps.BookService,
	}

	router.Handle("GET /book", handler.GetAllBook())
	router.Handle("GET /book/{id}", handler.GetOneBook())
	router.Handle("POST /book", handler.AddBook())
	router.Handle("POST /book/genre", handler.AddBookGenre())
	router.Handle("PUT /book/{id}", handler.UpdateBook())
	router.Handle("PUT /book/genre/{id}", handler.UpdateBookGenre())
	router.Handle("DELETE /book/{id}", handler.DeleteBook())
	router.Handle("DELETE /book/genre/{id}", handler.DeleteBookGenre())
}

func (handler *BookHandlers) GetAllBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *BookHandlers) GetOneBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *BookHandlers) AddBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *BookHandlers) AddBookGenre() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *BookHandlers) UpdateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *BookHandlers) UpdateBookGenre() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *BookHandlers) DeleteBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *BookHandlers) DeleteBookGenre() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
