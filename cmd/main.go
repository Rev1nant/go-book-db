package main

import (
	"net/http"

	"github.com/Rev1nant/go-book-db/configs"
	"github.com/Rev1nant/go-book-db/internal/repository"
	"github.com/Rev1nant/go-book-db/internal/server"
	"github.com/Rev1nant/go-book-db/internal/service"
	"github.com/Rev1nant/go-book-db/pkg/datebase"
)

func main() {
	conf := configs.LoadConfig("../.env")
	dsn := conf.DB.Dsn
	db := datebase.NewDB(dsn)
	defer db.DB.Close()

	repo := repository.NewRepositories(db)
	serv := service.NewServices(repo)

	http.HandleFunc("/author", serv.Author)
	http.HandleFunc("/author", serv.Author)
	http.HandleFunc("/author", serv.Author)
	http.HandleFunc("/author", serv.Author)
	http.HandleFunc("/author", serv.Author)

	http.HandleFunc("/genre", serv.Genre)
	http.HandleFunc("/genre", serv.Genre)
	http.HandleFunc("/genre", serv.Genre)
	http.HandleFunc("/genre", serv.Genre)
	http.HandleFunc("/genre", serv.Genre)

	http.HandleFunc("/book", serv.Book)
	http.HandleFunc("/book", serv.Book)
	http.HandleFunc("/book", serv.Book)
	http.HandleFunc("/book", serv.Book)
	http.HandleFunc("/book", serv.Book)
	http.HandleFunc("/book", serv.Book)
	http.HandleFunc("/book", serv.Book)
	http.HandleFunc("/book", serv.Book)

	server := server.Server{}
	server.Run("8080")
}
