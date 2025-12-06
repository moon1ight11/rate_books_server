package users

import (
	"fmt"
	"github.com/google/uuid"
)

// добавление пользователя в БД
func (db *Repo) CreateUser(NewUser User) (uuid.UUID, error) {
	transaction, err := db.DB.Begin()
	if err != nil {
		return uuid.Nil, err
	}

	defer transaction.Rollback()

	query := `
				INSERT INTO rate_books.users (name, pass, email)
				VALUES ($1, $2, $3)
				RETURNING id
			`

	var user_id uuid.UUID

	err = transaction.QueryRow(query, NewUser.Name, NewUser.Pass, NewUser.Email).Scan(&user_id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error in AddUser query: %w", err)
	}

	transaction.Commit()

	return user_id, nil
}
