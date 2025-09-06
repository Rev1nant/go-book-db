package handler

import (
	"net/http"

	"github.com/Rev1nant/go-book-db/internal/model"
	"github.com/Rev1nant/go-book-db/internal/service"
	"github.com/Rev1nant/go-book-db/pkg/req"
	"github.com/Rev1nant/go-book-db/pkg/res"
)

type BookHandlers struct {
	BookService  service.Book
	GenreService service.Genre
}

func NewBookHandler(router *http.ServeMux, deps BookHandlers) {
	handler := &BookHandlers{
		BookService:  deps.BookService,
		GenreService: deps.GenreService,
	}

	router.Handle("GET /book", handler.GetAllBook())
	router.Handle("GET /book/{id}", handler.GetOneBook())
	router.Handle("POST /book", handler.AddBook())
	router.Handle("POST /book/genre/{id}", handler.AddBookGenre())
	router.Handle("PUT /book/{id}", handler.UpdateBook())
	router.Handle("PUT /book/genre/{id}", handler.UpdateBookGenre())
	router.Handle("DELETE /book/{id}", handler.DeleteBook())
	router.Handle("DELETE /book/genre/{id}", handler.DeleteBookGenre())
}

func (handler *BookHandlers) GetAllBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := handler.BookService.GetAllBook()
		if err != nil {
			res.Json(w, "Error 500", http.StatusInternalServerError)
			return
		}

		res.Json(w, books, http.StatusOK)
	}
}

func (handler *BookHandlers) GetOneBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := req.GetId(&w, r)
		if err != nil {
			return
		}

		book, err := handler.BookService.GetOneBook(id)
		if err != nil {
			res.Json(w, "Error 500", http.StatusInternalServerError)
			return
		}

		res.Json(w, book, http.StatusOK)
	}
}

func (handler *BookHandlers) AddBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.Body[model.Book](&w, r)
		if err != nil {
			return
		}

		err = handler.BookService.AddBook(*body)
		if err != nil {
			res.Json(w, "Error 500", http.StatusInternalServerError)
			return
		}

		res.Json(w, "Ok", http.StatusOK)
	}
}

func (handler *BookHandlers) AddBookGenre() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookID, err := req.GetId(&w, r)
		if err != nil {
			return
		}

		body, err := req.Body[model.Genre](&w, r)
		if err != nil {
			return
		}

		genreID, err := handler.GenreService.GetGenreID(body.NameGenre)
		if err != nil {
			res.Json(w, "Error 500", http.StatusInternalServerError)
			return
		}

		err = handler.BookService.AddBookGenre(bookID, genreID)
		if err != nil {
			res.Json(w, "Error 500", http.StatusInternalServerError)
			return
		}

		res.Json(w, "ok", http.StatusOK)
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
