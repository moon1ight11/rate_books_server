package database

import (
	"log"
	"rate_books/internal/model"
)

// запрос на добавление новой книги
func InsertNewBook(b model.Book2) error {
	AuthorID, err := SearchAuthorId(b.Author)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Айдишник автора нашли:", AuthorID)

	query := "INSERT INTO rate_books (title, author_id, year_public, year_read, rate, time_stamp) VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP)"

	_, err = DB.Exec(query, b.Title, AuthorID, b.Year_public, b.Year_read, b.Rate)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
