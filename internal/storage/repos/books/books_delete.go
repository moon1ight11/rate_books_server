package books

import (
	"fmt"
	"github.com/google/uuid"
)

// удаление книги из БД
func (db *Repo) DeleteBook(bookID uuid.UUID) error {
	query := `
				DELETE FROM rate_books.books
				WHERE id = $1
			`
	_, err := db.DB.Exec(query, bookID)
	if err != nil {
		return fmt.Errorf("error in DeleteBook query: %w", err)
	}

	return nil
}
