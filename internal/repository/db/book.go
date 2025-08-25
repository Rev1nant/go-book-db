package db

import (
	"log"

	"github.com/Rev1nant/go-book-db/internal/model"
	"github.com/Rev1nant/go-book-db/pkg/datebase"
)

type BookRepo struct {
	db datebase.DB
}

func NewBookRepo(db *datebase.DB) *BookRepo {
	return &BookRepo{
		db: *db,
	}
}

func (r *BookRepo) GetAllBook() ([]model.Book, error) {
	rows, err := r.db.DB.Query(`SELECT title, firstname_author, lastname_author, price, amount FROM book INNER JOIN author USING(author_id)`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []model.Book{}
	for rows.Next() {
		book := model.Book{}
		err = rows.Scan(&book.Title, &book.Author.FirstName, &book.Author.LastName, &book.Price, &book.Amount)
		if err != nil {
			log.Println(err)
		}
		row, err := r.db.DB.Query(`SELECT name_genre FROM genre INNER JOIN book_genre USING(genre_id) INNER JOIN book USING(book_id) INNER JOIN author USING(author_id) WHERE title = $1, firstname_author = $2, lastname_author = $3 ;`, book.Title, book.Author.FirstName, book.Author.LastName)
		if err != nil {
			return nil, err
		}
		defer row.Close()

		genres := []model.Genre{}
		for row.Next() {
			genre := model.Genre{}
			err = row.Scan(&genre.NameGenre)
			if err != nil {
				log.Println(err)
			}
			genres = append(genres, genre)
		}

		book.Genre = genres
		books = append(books, book)
	}

	return books, nil
}

func (r *BookRepo) GetOneBook(id int) (model.Book, error) {
	rows, err := r.db.DB.Query(`SELECT title, firstname_author, lastname_author, price, amount FROM book INNER JOIN author USING(author_id) WHERE book_id = $1`, id)
	if err != nil {
		return model.Book{}, err
	}
	defer rows.Close()

	book := model.Book{}
	rows.Next()
	err = rows.Scan(&book.Title, &book.Author.FirstName, &book.Author.LastName, &book.Price, &book.Amount)
	if err != nil {
		log.Println(err)
	}
	row, err := r.db.DB.Query(`SELECT name_genre FROM genre INNER JOIN book_genre USING(genre_id) WHERE book_id = $1;`, id)
	if err != nil {
		return model.Book{}, err
	}

	genres := []model.Genre{}
	for row.Next() {
		genre := model.Genre{}
		err = row.Scan(&genre.NameGenre)
		if err != nil {
			log.Println(err)
		}
		genres = append(genres, genre)
	}

	book.Genre = genres
	return book, nil
}

func (r *BookRepo) AddBook(book model.Book, authorID int) error {
	_, err := r.db.DB.Exec(`INSERT INTO book (title, author_id, price, amount) VALUES ($1, $2, $3, $4);`, book.Title, authorID, book.Price, book.Amount)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepo) AddBookGenre(bookID, genreID int) error {
	_, err := r.db.DB.Exec(`INSERT INTO book_genre (book_id, genre_id) VALUES ($1, $2);`, bookID, genreID)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepo) UpdateBook(book model.Book, authorID, bookID int) error {
	_, err := r.db.DB.Exec(`UPDATE book SET title = $1, author_id = $2, price = $3, amount = $4 WHERE book_id = $5`, book.Title, authorID, book.Price, book.Amount, bookID)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepo) UpdateBookGenre(bookGenreID, genreID int) error {
	_, err := r.db.DB.Exec(`UPDATE book_genre SET genre_id = $1 WHERE book_genre_id = $2`, genreID, bookGenreID)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepo) DeleteBook(id int) error {
	_, err := r.db.DB.Exec(`DELETE from book WHERE book_id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepo) DeleteBookGenre(bookGenreID int) error {
	_, err := r.db.DB.Exec(`DELETE FROM book_genre WHERE book_genre_id = $1`, bookGenreID)
	if err != nil {
		return err
	}

	return nil
}
