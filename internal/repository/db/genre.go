package db

import (
	"log"

	"github.com/Rev1nant/go-book-db/internal/model"
	"github.com/Rev1nant/go-book-db/pkg/datebase"
)

type GenreRepo struct {
	db datebase.DB
}

func NewGenreRepo(db *datebase.DB) *GenreRepo {
	return &GenreRepo{
		db: *db,
	}
}

func (r *GenreRepo) GetAllGenre() ([]model.Genre, error) {
	rows, err := r.db.DB.Query(`SELECT name_genre FROM genre`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	genres := []model.Genre{}
	for rows.Next() {
		var genre model.Genre
		err = rows.Scan(&genre.NameGenre)
		if err != nil {
			log.Println(err)
		}
		genres = append(genres, genre)
	}

	return genres, nil
}

func (r *GenreRepo) GetOneGenre(id int) (model.Genre, error) {
	rows, err := r.db.DB.Query(`SELECT name_genre FROM genre WHERE genre_id = $1`, id)
	if err != nil {
		return model.Genre{}, err
	}
	defer rows.Close()

	genre := model.Genre{}
	rows.Next()
	err = rows.Scan(&genre.NameGenre)
	if err != nil {
		log.Println(err)
	}

	return genre, nil
}

func (r *GenreRepo) AddGenre(genre model.Genre) error {
	_, err := r.db.DB.Exec(`INSERT INTO genre (name_genre) VALUES ($1);`, genre.NameGenre)
	if err != nil {
		return err
	}

	return nil
}

func (r *GenreRepo) UpdateGenre(id int, genre model.Genre) error {
	_, err := r.db.DB.Exec(`UPDATE genre SET name_genre = $1 WHERE genre_id = $2`, genre.NameGenre, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *GenreRepo) DeleteGenre(id int) error {
	_, err := r.db.DB.Exec(`DELETE FROM genre WHERE genre_id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
