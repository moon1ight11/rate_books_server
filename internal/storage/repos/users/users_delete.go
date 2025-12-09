package users

import (
	"fmt"
	"github.com/google/uuid"
)

// удаление пользователя из БД
func (db *Repo) DeleteUser(userID uuid.UUID) error {
	query := `
				DELETE FROM rate_books.users
				WHERE id = $1
			`
	_, err := db.DB.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("error in DeleteUser query: %w", err)
	}

	return nil
}
