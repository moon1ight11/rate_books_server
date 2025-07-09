package database

import (
	"fmt"
	"rate_books/internal/model"
)

// запрос на все книги из списка
func SelectBooks(page_number int, page_size int, where string, sort_field string, sort_order string, args []interface{}) ([]model.Book, error) {
	queryArgs := make([]interface{}, len(args))
	copy(queryArgs, args)
	queryArgs = append(queryArgs, page_number*page_size, page_size)

	query := fmt.Sprintf(`SELECT
				rb.title, 
				a.author_name, 
				rb.year_public, 
				rb.year_read, 
				rb.rate
			FROM rate_books rb
			JOIN authors a ON rb.author_id = a.id
			%s 
			ORDER BY %s %s 
			OFFSET $%d LIMIT $%d`,
		where, sort_field, sort_order, len(args)+1, len(args)+2)

	rows, err := DB.Query(
		query, queryArgs...,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Books []model.Book
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.Title, &book.Author.Author_name, &book.Year_public, &book.Year_read, &book.Rate)
		if err != nil {
			return nil, err
		}
		Books = append(Books, book)
	}
	return Books, nil
}

// запрос на общее количество книг
func SelectAmountOfBooks(where string, args []interface{}) (AmountOfBooks int, err error) {

	var total int
	query := fmt.Sprintf(`
        SELECT COUNT(*)
        FROM rate_books rb
        JOIN authors a ON rb.author_id = a.id
        %s`, where)
	err = DB.QueryRow(query, args...).Scan(&total)
	if err != nil {
		return 0, err
	}

	AmountOfBooks = total

	return AmountOfBooks, nil
}
