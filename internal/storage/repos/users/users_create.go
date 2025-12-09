package users

import (
	"fmt"
	"github.com/google/uuid"
)

// добавление пользователя в БД
func (db *Repo) CreateUser(NewUser User) (uuid.UUID, error) {
	query := `
				INSERT INTO rate_books.users (name, pass, email)
				VALUES ($1, $2, $3)
				RETURNING id
			`

	var userID uuid.UUID

	err := db.DB.QueryRow(query, NewUser.Name, NewUser.Pass, NewUser.Email).Scan(&userID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error in CreateUser query: %w", err)
	}

	return userID, nil
}
