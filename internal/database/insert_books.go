package database

import (
	"log"
	"rate_books/internal/model"
)

// запрос на добавление новой книги
func InsertNewBook(b model.Book2, us_id int) error {
	AuthorID, err := SearchAuthorId(b.Author)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Айдишник автора нашли:", AuthorID)

	query := "INSERT INTO rate_books (title, author_id, year_public, year_read, rate, time_stamp, book_owner) VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, $6)"

	_, err = DB.Exec(query, b.Title, AuthorID, b.Year_public, b.Year_read, b.Rate, us_id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
