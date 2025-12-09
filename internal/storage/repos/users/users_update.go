package users

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

// обновление имени пользователя
func (db *Repo) UpdateUserName(newUserName string, userID uuid.UUID, tx *sql.Tx) error {
	query := `
				UPDATE rate_books.users
				SET name = $1, updated_at = NOW()
				WHERE id = $2
			`
	_, err := tx.Exec(query, newUserName, userID)
	if err != nil {
		return fmt.Errorf("error in UpdateUserName query: %w", err)
	}

	return nil
}

// обновление email пользователя
func (db *Repo) UpdateUserEmail(newUserEmail string, userID uuid.UUID, tx *sql.Tx) error {
	query := `
				UPDATE rate_books.users
				SET email = $1, updated_at = NOW()
				WHERE id = $2
			`
	_, err := tx.Exec(query, newUserEmail, userID)
	if err != nil {
		return fmt.Errorf("error in UpdateUserEmail query: %w", err)
	}

	return nil
}

// обновление пароля пользователя
func (db *Repo) UpdateUserPass(newUserPass string, userID uuid.UUID, tx *sql.Tx) error {
	query := `
				UPDATE rate_books.users
				SET pass = $1, updated_at = NOW()
				WHERE id = $2
			`
	_, err := tx.Exec(query, newUserPass, userID)
	if err != nil {
		return fmt.Errorf("error in UpdateUserPass query: %w", err)
	}

	return nil
}
