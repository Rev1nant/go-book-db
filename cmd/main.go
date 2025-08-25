package main

import (
	"github.com/Rev1nant/go-book-db/configs"
	"github.com/Rev1nant/go-book-db/internal/server"
	"github.com/Rev1nant/go-book-db/pkg/datebase"
)

func main() {
	conf := configs.LoadConfig("../.env")
	dsn := conf.DB.Dsn
	db := datebase.NewDB(dsn)
	defer db.DB.Close()

	// repo := repository.NewRepositories(db)

	// http.HandleFunc("/book", service)

	server := server.Server{}
	server.Run("8080")
}
