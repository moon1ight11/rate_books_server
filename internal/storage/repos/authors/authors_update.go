package authors

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

// обновление фамилии
func (db *Repo) UpdateAuthorSurname(authorID uuid.UUID, newAuthorSurname string, tx *sql.Tx) error {
	query := `
				UPDATE rate_books.authors
				SET surname = $1, updated_at = NOW()
				WHERE id = $2
			`
	_, err := tx.Exec(query, newAuthorSurname, authorID)
	if err != nil {
		return fmt.Errorf("error in UpdateAuthorSurname query: %w", err)
	}

	return nil
}

// обновление имени
func (db *Repo) UpdateAuthorName(authorID uuid.UUID, newAuthorName string, tx *sql.Tx) error {
	query := `
				UPDATE rate_books.authors
				SET name = $1, updated_at = NOW()
				WHERE id = $2
			`
	_, err := tx.Exec(query, newAuthorName, authorID)
	if err != nil {
		return fmt.Errorf("error in UpdateAuthorName query: %w", err)
	}

	return nil
}

// обновление года рождения
func (db *Repo) UpdateAuthorYearBorn(authorID uuid.UUID, newAuthorYearBorn int, tx *sql.Tx) error {
	query := `
				UPDATE rate_books.authors
				SET year_born = $1, updated_at = NOW()
				WHERE id = $2
			`
	_, err := tx.Exec(query, newAuthorYearBorn, authorID)
	if err != nil {
		return fmt.Errorf("error in UpdateAuthorYearBorn query: %w", err)
	}

	return nil
}

// обновление страны
func (db *Repo) UpdateAuthorCountry(authorID uuid.UUID, newAuthorCountry string, tx *sql.Tx) error {
	query := `
				UPDATE rate_books.authors
				SET country = $1, updated_at = NOW()
				WHERE id = $2
			`
	_, err := tx.Exec(query, newAuthorCountry, authorID)
	if err != nil {
		return fmt.Errorf("error in UpdateAuthorCountry query: %w", err)
	}

	return nil
}

// обновление описания
func (db *Repo) UpdateAuthorDescription(authorID uuid.UUID, newAuthorDescription string, tx *sql.Tx) error {
	query := `
				UPDATE rate_books.authors
				SET description = $1, updated_at = NOW()
				WHERE id = $2
			`
	_, err := tx.Exec(query, newAuthorDescription, authorID)
	if err != nil {
		return fmt.Errorf("error in UpdateAuthorDescription query: %w", err)
	}

	return nil
}