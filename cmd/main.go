package main

import (
	"net/http"

	"github.com/Rev1nant/go-book-db/configs"
	"github.com/Rev1nant/go-book-db/internal/handler"
	"github.com/Rev1nant/go-book-db/internal/repository"
	"github.com/Rev1nant/go-book-db/internal/server"
	"github.com/Rev1nant/go-book-db/internal/service"
	"github.com/Rev1nant/go-book-db/pkg/db"
)

func App() http.Handler {
	conf := configs.LoadConfig("../.env")
	dsn := conf.DB.Dsn
	db := db.NewDB(dsn)

	repo := repository.NewRepositories(db)
	serv := service.NewServices(repo)
	router := http.NewServeMux()
	handler.NewAuthorHandler(router, handler.AuthorHandlers{
		AuthorService: serv.Author,
	})

	handler.NewGenreHendler(router, handler.GenreHandlers{
		GenreService: serv.Genre,
	})

	return router
}

func main() {
	server := server.Server{}
	server.Run("8080", App)
}
