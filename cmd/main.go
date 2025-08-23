package main

import (
	"fmt"

	"github.com/Rev1nant/go-book-db/configs"
	"github.com/Rev1nant/go-book-db/internal/repository"
	"github.com/Rev1nant/go-book-db/pkg/datebase"
)

func main() {
	conf := configs.LoadConfig("../.env")
	dsn := conf.DB.Dsn
	db := datebase.NewDB(dsn)
	defer db.DB.Close()

	repo := repository.NewRepositories(db)

	fmt.Println(repo.Genre.GetAllGenre())
}
