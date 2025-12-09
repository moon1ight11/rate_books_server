package books

import (
	"fmt"

	"github.com/google/uuid"
)

// получение одной книги по id
func (db *Repo) GetBookById(bookID uuid.UUID, userID uuid.UUID) (Book, error) {
	query := `
				SELECT id, title, author_id, genre, year_public, year_read, grade, description
				FROM rate_books.books
				WHERE id = $1 AND owner_id = $2
			`
	var book Book
	err := db.DB.QueryRow(query, bookID, userID).Scan(
		&book.ID,
		&book.Title,
		&book.AuthorID,
		&book.Genre,
		&book.YearPublic,
		&book.YearRead,
		&book.Grade,
		&book.Description,
	)
	if err != nil {
		return Book{}, fmt.Errorf("error in GetBookById query: %w", err)
	}

	return book, nil
}

// получение списка всех книг пользователя
func (db *Repo) GetAllBooks(userID uuid.UUID) ([]Book, error) {
	query := `
				SELECT id, title, author_id, genre, year_public, year_read, grade, description
				FROM rate_books.books
				WHERE owner_id = $1
			`
	var books []Book
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error in GetAllBooks query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.AuthorID,
			&book.Genre,
			&book.YearPublic,
			&book.YearRead,
			&book.Grade,
			&book.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("error in GetAllBooks scan: %w", err)
		}
		books = append(books, book)
	}
	return books, nil
}
