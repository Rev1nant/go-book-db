package db

import (
	"log"

	"github.com/Rev1nant/go-book-db/internal/domain"
	"github.com/Rev1nant/go-book-db/pkg/datebase"
)

type AuthorRepo struct {
	db datebase.DB
}

func NewAuthorRepo(db *datebase.DB) *AuthorRepo {
	return &AuthorRepo{
		db: *db,
	}
}

func (r *AuthorRepo) GetAllAuthor() ([]domain.Author, error) {
	rows, err := r.db.DB.Query(`SELECT * FROM author`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	authors := []domain.Author{}
	for rows.Next() {
		author := domain.Author{}
		err = rows.Scan(&author.AuthorID, &author.FirstName, &author.LastName)
		if err != nil {
			log.Println(err)
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (r *AuthorRepo) GetOneAuthor(id int) (domain.Author, error) {
	row, err := r.db.DB.Query(`SELECT * FROM author WHERE author_id = $1`, id)
	if err != nil {
		return domain.Author{}, err
	}
	defer row.Close()

	author := domain.Author{}
	row.Next()
	err = row.Scan(&author.AuthorID, &author.FirstName, &author.LastName)
	if err != nil {
		log.Println(err)
	}

	return author, nil
}

func (r *AuthorRepo) AddAuthor(firstName, lastName string) error {
	_, err := r.db.DB.Exec(`INSERT INTO author (firstname_author, lastname_author) VALUES ($1, $2)`, firstName, lastName)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthorRepo) UpdateAuthor(id int, firstNameNew, lastNameNew string) error {
	_, err := r.db.DB.Exec(`UPDATE author SET firstname_author = $1, lastname_author = $2 WHERE author_id = $3`, firstNameNew, lastNameNew, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthorRepo) DeleteAuthor(id int) error {
	_, err := r.db.DB.Exec(`DELETE FROM author WHERE author_id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
