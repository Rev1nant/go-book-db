package handler

import (
	"net/http"

	"github.com/Rev1nant/go-book-db/internal/model"
	"github.com/Rev1nant/go-book-db/internal/service"
	"github.com/Rev1nant/go-book-db/pkg/req"
	"github.com/Rev1nant/go-book-db/pkg/res"
)

type AuthorHandlers struct {
	AuthorService service.Author
}

func NewAuthorHandler(router *http.ServeMux, deps AuthorHandlers) {
	handler := &AuthorHandlers{
		AuthorService: deps.AuthorService,
	}

	router.Handle("GET /author", handler.GetAllAuthor())
	router.Handle("GET /author/{id}", handler.GetOneAuthor())
	router.Handle("POST /author", handler.AddAuthor())
	router.Handle("PUT /author/{id}", handler.UpdateAuthor())
	router.Handle("DELETE /author/{id}", handler.DeleteAuthor())
}

func (handler *AuthorHandlers) GetAllAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authors, err := handler.AuthorService.GetAllAuthor()
		if err != nil {
			res.Json(w, "Error 500", http.StatusInternalServerError)
		}

		res.Json(w, authors, http.StatusOK)
	}
}
func (handler *AuthorHandlers) GetOneAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := req.GetId(&w, r)
		if err != nil {
			return
		}

		author, err := handler.AuthorService.GetOneAuthor(id)
		if err != nil {
			res.Json(w, "Error 404", http.StatusNotFound)
			return
		}

		res.Json(w, author, http.StatusOK)
	}
}

func (handler *AuthorHandlers) AddAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.Body[model.Author](&w, r)
		if err != nil {
			return
		}
		err = handler.AuthorService.AddAuthor(*body)
		if err != nil {
			return
		}

		res.Json(w, "Ok", http.StatusOK)
	}
}

func (handler *AuthorHandlers) UpdateAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := req.GetId(&w, r)
		if err != nil {
			return
		}

		body, err := req.Body[model.Author](&w, r)
		if err != nil {
			return
		}

		err = handler.AuthorService.UpdateAuthor(id, *body)
		if err != nil {
			res.Json(w, "Error 500", http.StatusInternalServerError)
			return
		}

		res.Json(w, "Ok", http.StatusOK)
	}
}

func (handler *AuthorHandlers) DeleteAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := req.GetId(&w, r)
		if err != nil {
			return
		}

		err = handler.AuthorService.DeleteAuthor(id)
		if err != nil {
			res.Json(w, "Error 500", http.StatusInternalServerError)
			return
		}

		res.Json(w, "Ok", http.StatusOK)
	}
}
