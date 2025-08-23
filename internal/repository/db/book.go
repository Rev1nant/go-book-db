package db

import (
	"log"

	"github.com/Rev1nant/go-book-db/internal/domain"
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

func (r *BookRepo) GetAllBook() ([]domain.Book, error) {
	rows, err := r.db.DB.Query(`SELECT book_id, title, author_id, firstname_author, lastname_author, price, amount FROM book INNER JOIN author USING(author_id)`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []domain.Book{}
	for rows.Next() {
		book := domain.Book{}
		err = rows.Scan(&book.BookID, &book.Title, &book.Author.AuthorID, &book.Author.FirstName, &book.Author.LastName, &book.Price, &book.Amount)
		if err != nil {
			log.Println(err)
		}
		row, err := r.db.DB.Query(`SELECT genre_id, name_genre FROM genre INNER JOIN book_genre USING(genre_id) WHERE book_id = $1;`, book.BookID)
		if err != nil {
			return nil, err
		}
		defer row.Close()

		genres := []domain.Genre{}
		for row.Next() {
			genre := domain.Genre{}
			err = row.Scan(&genre.GenreID, &genre.NameGenre)
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

func (r *BookRepo) GetOneBook(id int) (domain.Book, error) {
	rows, err := r.db.DB.Query(`SELECT book_id, title, author_id, firstname_author, lastname_author, price, amount FROM book INNER JOIN author USING(author_id) WHERE book_id = $1`, id)
	if err != nil {
		return domain.Book{}, err
	}
	defer rows.Close()

	book := domain.Book{}
	rows.Next()
	err = rows.Scan(&book.BookID, &book.Title, &book.Author.AuthorID, &book.Author.FirstName, &book.Author.LastName, &book.Price, &book.Amount)
	if err != nil {
		log.Println(err)
	}
	row, err := r.db.DB.Query(`SELECT name_genre FROM genre INNER JOIN book_genre USING(genre_id) WHERE book_id = $1;`, id)
	if err != nil {
		return domain.Book{}, err
	}

	genres := []domain.Genre{}
	for row.Next() {
		genre := domain.Genre{}
		err = row.Scan(&genre.GenreID, &genre.NameGenre)
		if err != nil {
			log.Println(err)
		}
		genres = append(genres, genre)
	}

	book.Genre = genres
	return book, nil
}

func (r *BookRepo) AddBook(title string, authorID int, price float64, amount int) error {
	_, err := r.db.DB.Exec(`INSERT INTO book (title, author_id, price, amount) VALUES ($1, $2, $3, $4);`, title, authorID, price, amount)
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

func (r *BookRepo) UpdateBook(id int, title string, authorID int, price float64, amount int) error {
	_, err := r.db.DB.Exec(`UPDATE book SET title = $1, author_id = $2, price = $3, amount = $4 WHERE book_id = $5`, title, authorID, price, amount, id)
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
