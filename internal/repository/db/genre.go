package db

import (
	"log"

	"github.com/Rev1nant/go-book-db/internal/domain"
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

func (r *GenreRepo) GetAllGenre() ([]domain.Genre, error) {
	rows, err := r.db.DB.Query(`SELECT * FROM genre`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	genres := []domain.Genre{}
	for rows.Next() {
		var genre domain.Genre
		err = rows.Scan(&genre.GenreID, &genre.NameGenre)
		if err != nil {
			log.Println(err)
		}
		genres = append(genres, genre)
	}

	return genres, nil
}

func (r *GenreRepo) GetOneGenre(id int) (domain.Genre, error) {
	rows, err := r.db.DB.Query(`SELECT * FROM genre WHERE genre_id = $1`, id)
	if err != nil {
		return domain.Genre{}, err
	}
	defer rows.Close()

	genre := domain.Genre{}
	rows.Next()
	err = rows.Scan(&genre.GenreID, &genre.NameGenre)
	if err != nil {
		log.Println(err)
	}

	return genre, nil
}

func (r *GenreRepo) AddGenre(genreName string) error {
	_, err := r.db.DB.Exec(`INSERT INTO genre (name_genre) VALUES ($1);`, genreName)
	if err != nil {
		return err
	}

	return nil
}

func (r *GenreRepo) UpdateGenre(id int, genreNameNew string) error {
	_, err := r.db.DB.Exec(`UPDATE genre SET name_genre = $1 WHERE genre_id = $2`, genreNameNew, id)
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
