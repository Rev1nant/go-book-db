package main

import (
	"github.com/Rev1nant/go-book-db/configs"
	"github.com/Rev1nant/go-book-db/internal/repository"
	"github.com/Rev1nant/go-book-db/internal/server"
	"github.com/Rev1nant/go-book-db/internal/service"
	"github.com/Rev1nant/go-book-db/pkg/db"
)

func main() {
	conf := configs.LoadConfig("../.env")
	dsn := conf.DB.Dsn
	db := db.NewDB(dsn)
	defer db.DB.Close()

	repo := repository.NewRepositories(db)
	serv := service.NewServices(repo)

	server := server.Server{}
	server.Run("8080")
}
