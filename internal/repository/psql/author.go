package psql

import (
	"log"

	"github.com/Rev1nant/go-book-db/internal/model"
	"github.com/Rev1nant/go-book-db/pkg/db"
)

type AuthorRepo struct {
	db db.DB
}

func NewAuthorRepo(db *db.DB) *AuthorRepo {
	return &AuthorRepo{
		db: *db,
	}
}

func (r *AuthorRepo) FindAll() ([]model.Author, error) {
	rows, err := r.db.DB.Query(`SELECT firstname_author, lastname_author FROM author`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	authors := []model.Author{}
	for rows.Next() {
		author := model.Author{}
		err = rows.Scan(&author.FirstName, &author.LastName)
		if err != nil {
			log.Println(err)
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (r *AuthorRepo) FindByID(id int) (model.Author, error) {
	rows, err := r.db.DB.Query(`SELECT firstname_author, lastname_author FROM author WHERE author_id = $1`, id)
	if err != nil {
		return model.Author{}, err
	}
	defer rows.Close()

	author := model.Author{}
	rows.Next()
	err = rows.Scan(&author.FirstName, &author.LastName)
	if err != nil {
		return author, err
	}

	return author, nil
}

func (r *AuthorRepo) Create(author model.Author) error {
	_, err := r.db.DB.Exec(`INSERT INTO author (firstname_author, lastname_author) VALUES ($1, $2)`, author.FirstName, author.LastName)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthorRepo) Update(id int, author model.Author) error {
	_, err := r.db.DB.Exec(`UPDATE author SET firstname_author = $1, lastname_author = $2 WHERE author_id = $3`, author.FirstName, author.LastName, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthorRepo) Delete(id int) error {
	_, err := r.db.DB.Exec(`DELETE FROM author WHERE author_id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthorRepo) GetAuthorID(author model.Author) (int, error) {
	row, err := r.db.DB.Query(`SELECT author_id FROM author WHERE firstname_author = $1 AND lastname_author = $2`, author.FirstName, author.LastName)
	if err != nil {
		return 0, err
	}

	defer row.Close()

	var authorID int
	row.Next()
	err = row.Scan(&authorID)
	if err != nil {
		return 0, err
	}

	return authorID, nil
}
