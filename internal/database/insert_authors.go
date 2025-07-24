package database

import (
	"log"
	"rate_books/internal/model"
)

// запрос на добавление нового автора
func InsertNewAuthor(a model.Authors) error {
	query := "INSERT INTO authors (author_name,year_b,country) VALUES ($1, $2, $3)"

	_, err := DB.Exec(query, a.Author_name, a.Year_born, a.Country)
	if err != nil {
		log.Println("Error in query InsertNewAuthor",err)
		return err
	}
	return nil
}

// запрос на id автора
func SearchAuthorId(author_name string) (int, error) {
	query :=
		`SELECT id 
			FROM 
				authors
			WHERE author_name = $1`

	rows, err := DB.Query(
		query, author_name,
	)
	if err != nil {
		log.Println("Error in query SearchAuthorId", err)
		return 0, err
	}
	defer rows.Close()

	var AuthorID int
	for rows.Next() {
		err = rows.Scan(&AuthorID)
		if err != nil {
			log.Println("Error in scan SearchAuthorId", err)
			return 0, err
		}
	}

	return AuthorID, nil
}
