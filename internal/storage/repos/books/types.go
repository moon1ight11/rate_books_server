package books

import (
	"rate_books/internal/storage"

	"github.com/google/uuid"
)

type Repo struct {
	storage.DataBase
}

func NewBooksRepo(db *storage.DataBase) *Repo {
	return &Repo{DataBase: *db}
}

type Book struct {
	ID          uuid.UUID `json:"book_id"`
	Title       string    `json:"title"`
	AuthorID    uuid.UUID `json:"author_id"`
	Genre       string    `json:"genre"`
	YearPublic  int       `json:"year_public"`
	YearRead    int       `json:"year_read"`
	Grade       int       `json:"grade"`
	Description string    `json:"book_description"`
}
