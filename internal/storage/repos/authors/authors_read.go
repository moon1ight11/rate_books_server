package authors

import (
	"fmt"
	"github.com/google/uuid"
)

// получение автора по id
func (db *Repo) GetAuthorByID(authorID uuid.UUID) (Author, error) {
	query := `
				SELECT id, surname, name, country, year_born, description
				FROM rate_books.authors
				WHERE id = $1
			`
	var author Author
	err := db.DB.QueryRow(query, authorID).Scan(
		&author.ID,
		&author.Surname,
		&author.Name,
		&author.Country,
		&author.YearBorn,
		&author.Description,
	)
	if err != nil {
		return Author{}, fmt.Errorf("error in GetAuthorByID query: %w", err)
	}

	return author, nil
}

// получение списка всех авторов пользователя
func (db *Repo) GetAuthorsByUser(userID uuid.UUID) ([]Author, error) {
	query := `
				SELECT DISTINCT 
					a.id, 
					a.surname, 
					a.name, 
					a.country, 
					a.year_born, 
					a.description
				FROM rate_books.authors a
				JOIN rate_books.books b ON a.id = b.author_id
				WHERE b.owner_id = $1
			`
	var authors []Author
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error in GetAuthorsByUser query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var author Author
		err := rows.Scan(
			&author.ID,
			&author.Surname,
			&author.Name,
			&author.Country,
			&author.YearBorn,
			&author.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("error in GetAuthorsByUser scan: %w", err)
		}
		authors = append(authors, author)
	}
	return authors, nil
}

// получение списка всех авторов в системе
func (db *Repo) GetAllAuthors() ([]Author, error) {
	query := `
				SELECT id, surname, name, country, year_born, description
				FROM rate_books.authors
			`
	var authors []Author
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error in GetAllAuthors query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var author Author
		err := rows.Scan(
			&author.ID,
			&author.Surname,
			&author.Name,
			&author.Country,
			&author.YearBorn,
			&author.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("error in GetAllAuthors scan: %w", err)
		}
		authors = append(authors, author)
	}
	return authors, nil
}
