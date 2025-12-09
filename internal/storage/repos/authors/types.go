package authors

import (
	"rate_books/internal/storage"

	"github.com/google/uuid"
)

type Repo struct {
	storage.DataBase
}

func NewAuthorsRepo(db *storage.DataBase) *Repo {
	return &Repo{DataBase: *db}
}

type Author struct {
	ID          uuid.UUID `json:"author_id"`
	Surname     string    `json:"surname"`
	Name        string    `json:"name"`
	Country     string    `json:"country"`
	YearBorn    string    `json:"year_born"`
	Description string    `json:"author_description"`
}
