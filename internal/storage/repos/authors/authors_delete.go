package authors

import (
	"fmt"
	"github.com/google/uuid"
)

// удаление автора из БД
func (db *Repo) DeleteAuthor(authorID uuid.UUID) error {
	query := `
				DELETE FROM rate_books.authors
				WHERE id = $1
			`
	_, err := db.DB.Exec(query, authorID)
	if err != nil {
		return fmt.Errorf("error in DeleteAuthor query: %w", err)
	}

	return nil
}
