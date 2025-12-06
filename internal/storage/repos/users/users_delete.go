package users

import (
	"fmt"
	"github.com/google/uuid"
)

// удаление пользователя из БД
func (db *Repo) DeleteUser(user_id uuid.UUID) error {
	transaction, err := db.DB.Begin()
	if err != nil {
		return err
	}

	defer transaction.Rollback()

	query := `
				DELETE FROM rate_books.users
				WHERE id = $1
			`
	_, err = transaction.Exec(query, user_id)
	if err != nil {
		return fmt.Errorf("Error in DeleteUser query: %w", err)
	}

	transaction.Commit()
	
	return nil
}