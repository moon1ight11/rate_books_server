package books

import (
	"fmt"
	"github.com/google/uuid"
)

// добавление книги в БД
func (db *Repo) CreateBook(NewBook Book, ownerID uuid.UUID) (uuid.UUID, error) {
	query := `
				INSERT INTO rate_books.books (title, author_id, genre, owner_id, year_public, year_read, grade, description)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
				RETURNING id
			`

	err := db.DB.QueryRow(
		query,
		NewBook.Title,
		NewBook.AuthorID,
		NewBook.Genre,
		ownerID,
		NewBook.YearPublic,
		NewBook.YearRead,
		NewBook.Grade,
		NewBook.Description,
	).Scan(&NewBook.ID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error in CreateBook query: %w", err)
	}

	return NewBook.ID, nil
}
