package books

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

// обновление названия
func (db *Repo) UpdateBookTitle(bookID uuid.UUID, newBookTitle string, tx *sql.Tx) error {
	query := `
				UPDATE rate_books.books
				SET title = $1, updated_at = NOW()
				WHERE id = $2
			`
	_, err := tx.Exec(query, newBookTitle, bookID)
	if err != nil {
		return fmt.Errorf("error in UpdateBookTitle query: %w", err)
	}

	return nil
}

// обновление автора
func (db *Repo) UpdateBookAuthor(bookID uuid.UUID, newAuthorID uuid.UUID, tx *sql.Tx) error {
	query := `
				UPDATE rate_books.books
				SET author_id = $1, updated_at = NOW()
				WHERE id = $2
			`
	_, err := tx.Exec(query, newAuthorID, bookID)
	if err != nil {
		return fmt.Errorf("error in UpdateBookAuthor query: %w", err)
	}

	return nil
}

// обновление жанра
func (db *Repo) UpdateBookGenre(bookID uuid.UUID, newGenre string, tx *sql.Tx) error {
	query := `
				UPDATE rate_books.books
				SET genre = $1, updated_at = NOW()
				WHERE id = $2
			`
	_, err := tx.Exec(query, newGenre, bookID)
	if err != nil {
		return fmt.Errorf("error in UpdateBookGenre query: %w", err)
	}

	return nil
}

// обновление года выпуска
func (db *Repo) UpdateBookYearPublic(bookID uuid.UUID, newYearPublic int, tx *sql.Tx) error {
	query := `
				UPDATE rate_books.books
				SET year_public = $1, updated_at = NOW()
				WHERE id = $2
			`
	_, err := tx.Exec(query, newYearPublic, bookID)
	if err != nil {
		return fmt.Errorf("error in UpdateBookYearPublic query: %w", err)
	}

	return nil
}

// обновление года прочтения
func (db *Repo) UpdateBookYearRead(bookID uuid.UUID, newYearRead int, tx *sql.Tx) error {
	query := `
				UPDATE rate_books.books
				SET year_read = $1, updated_at = NOW()
				WHERE id = $2
			`
	_, err := tx.Exec(query, newYearRead, bookID)
	if err != nil {
		return fmt.Errorf("error in UpdateBookYearRead query: %w", err)
	}

	return nil
}

// обновление оценки
func (db *Repo) UpdateBookGrade(bookID uuid.UUID, newGrade int, tx *sql.Tx) error {
	query := `
				UPDATE rate_books.books
				SET grade = $1, updated_at = NOW()
				WHERE id = $2
			`
	_, err := tx.Exec(query, newGrade, bookID)
	if err != nil {
		return fmt.Errorf("error in UpdateBookGrade query: %w", err)
	}

	return nil
}

// обновление описания
func (db *Repo) UpdateBookDescription(bookID uuid.UUID, newDescription string, tx *sql.Tx) error {
	query := `
				UPDATE rate_books.books
				SET description = $1, updated_at = NOW()
				WHERE id = $2
			`
	_, err := tx.Exec(query, newDescription, bookID)
	if err != nil {
		return fmt.Errorf("error in UpdateBookDescription query: %w", err)
	}

	return nil
}
