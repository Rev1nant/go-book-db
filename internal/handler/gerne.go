package handler

import (
	"net/http"

	"github.com/Rev1nant/go-book-db/internal/model"
	"github.com/Rev1nant/go-book-db/internal/service"
	"github.com/Rev1nant/go-book-db/pkg/req"
	"github.com/Rev1nant/go-book-db/pkg/res"
)

type GenreHandlers struct {
	GenreService service.Genre
}

func NewGenreHendler(router *http.ServeMux, deps GenreHandlers) {
	handler := &GenreHandlers{
		GenreService: deps.GenreService,
	}

	router.Handle("GET /genre", handler.GetAllGenre())
	router.Handle("GET /genre/{id}", handler.GetOneGenre())
	router.Handle("POST /genre", handler.AddGenre())
	router.Handle("PUT /genre/{id}", handler.UpdateGenre())
	router.Handle("DELETE /genre/{id}", handler.DeleteGenre())
}

func (handler *GenreHandlers) GetAllGenre() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		genres, err := handler.GenreService.GetAllGenre()
		if err != nil {
			res.Json(w, "500", http.StatusInternalServerError)
		}

		res.Json(w, genres, http.StatusOK)
	}
}

func (handler *GenreHandlers) GetOneGenre() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := req.GetId(&w, r)
		if err != nil {
			return
		}

		genre, err := handler.GenreService.GetOneGenre(id)
		if err != nil {
			res.Json(w, "Error 500", http.StatusInternalServerError)
			return
		}

		res.Json(w, genre, http.StatusOK)
	}
}

func (handler *GenreHandlers) AddGenre() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.Body[model.Genre](&w, r)
		if err != nil {
			return
		}
		err = handler.GenreService.AddGenre(*body)
		if err != nil {
			return
		}

		res.Json(w, "Ok", http.StatusOK)
	}
}

func (handler *GenreHandlers) UpdateGenre() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := req.GetId(&w, r)
		if err != nil {
			return
		}

		body, err := req.Body[model.Genre](&w, r)
		if err != nil {
			return
		}

		err = handler.GenreService.UpdateGenre(id, *body)
		if err != nil {
			res.Json(w, "Error 500", http.StatusInternalServerError)
		}

		res.Json(w, "Ok", http.StatusOK)
	}
}

func (handler *GenreHandlers) DeleteGenre() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := req.GetId(&w, r)
		if err != nil {
			return
		}

		err = handler.GenreService.DeleteGenre(id)
		if err != nil {
			res.Json(w, "Error 500", http.StatusInternalServerError)
			return
		}

		res.Json(w, "Ok", http.StatusOK)
	}
}
