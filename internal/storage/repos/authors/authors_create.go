package authors

import (
	"fmt"
	"github.com/google/uuid"
)

// добавление автора в БД
func (db *Repo) CreateAuthor(NewAuthor Author) (uuid.UUID, error) {
		query := `
				INSERT INTO rate_books.authors (surname, name, country, year_born, description)
				VALUES ($1, $2, $3, $4, $5)
				RETURNING id
			`

	err := db.DB.QueryRow(
		query,
		NewAuthor.Surname,
		NewAuthor.Name,
		NewAuthor.Country,
		NewAuthor.YearBorn,
		NewAuthor.Description,
	).Scan(&NewAuthor.ID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error in CreateAuthor query: %w", err)
	}

	return NewAuthor.ID, nil
}