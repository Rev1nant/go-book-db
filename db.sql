/*Create database tables*/
CREATE TABLE genre (
    genre_id SERIAL PRIMARY KEY,
    name_genre VARCHAR(50) NOT NULL
);

CREATE TABLE author (
    author_id SERIAL PRIMARY KEY,
    firstname_author VARCHAR(50) NOT NULL,
    lastname_author VARCHAR(50)
);

CREATE TABLE book (
    book_id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    author_id INT REFERENCES author,
    price DECIMAL(8, 2) NOT NULL,
    amount INT NOT NULL
);

CREATE TABLE book_genre (
    book_genre_id SERIAL PRIMARY KEY,
    book_id INT NOT NULL REFERENCES book,
    genre_id INT NOT NULL REFERENCES genre,
    UNIQUE(book_id, genre_id)
);

/*Delete all database tables*/
DROP TABLE book_genre, book, genre, author;




/*Get all genre*/
`SELECT * FROM genre`

/*Get one genre*/
`SELECT * FROM genre WHERE genre_id = $1`, id

/*Added genre*/
`INSERT INTO genre (name_genre) VALUES ($1)`, genreName

/*Update genre*/
`UPDATE genre SET name_genre = $1 WHERE genre_id = $2`, genreNameNew, id

/*Delete genre*/
`DELETE FROM genre WHERE genre_id = $1`, id





/*Get all author*/
`SELECT * FROM author`

/*Get one author*/
`SELECT * FROM author WHERE author_id = $1`, id

/*Added author*/
`INSERT INTO author (firstname_author, lastname_author) VALUES ($1, $2)`, firstName, lastName

/*Update author*/
`UPDATE author SET firstname_author = $1, lastname_author = $2 WHERE author_id = $3`, firstNameNew, lastNameNew, id

/*Delete author*/
`DELETE FROM author WHERE author_id = $1`, id






/*Get all book*/
`SELECT book_id, title, author_id, firstname_author, lastname_author, price, amount FROM book INNER JOIN author USING(author_id)`

/*Get one book*/
`SELECT book_id, title, author_id, firstname_author, lastname_author, price, amount FROM book INNER JOIN author USING(author_id) WHERE book_id = $1`, id
`SELECT name_genre FROM genre INNER JOIN book_genre USING(genre_id) WHERE book_id = $1;`, id

/*Added book*/
`INSERT INTO book (title, author_id, price, amount) VALUES ($1, $2, $3, $4);`, title, authorID, price, amount

/*Added genres book*/
`INSERT INTO book_genre (book_id, genre_id) VALUES ($1, $2);`, title, firstName, lastName, nameGenre

/*Update book*/
`UPDATE book SET title = $1, author_id = $2, price = $3, amount = $4 WHERE book_id = $5`, title, authorID, price, amount, id

/*Update genre book*/
`UPDATE book_genre SET genre_id = $1 WHERE book_genre_id = $2`, genreID, bookGenreID

/*Delete book*/
`DELETE from book WHERE book_id = $1`, id

/*Delete genre book*/
`DELETE FROM book_genre WHERE book_genre_id = $1`, bookGenreID



/*Get book id*/
`SELECT book_id FROM book WHERE title = $1, author_id = $2`, title, authorID

/*Get genre id*/
`SELECT genre_id FROM genre WHERE name_genre = $1`, nameGenre

/*Get author id*/
`SELECT author_id FROM author WHERE firstname_author = $1 AND lastname_author = $2`, fistName,  lastName