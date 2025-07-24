package database

import (
	"fmt"
	"log"
	"rate_books/internal/model"
)

// запрос на все книги из списка
func SelectBooks(page_number int, page_size int, sort_field string, sort_order string, filters []interface{}, us_id int) ([]model.Book, error) {
	filterByTitle := filters[0]
	filterByAuthor := filters[1]
	filterYearPublFrom := filters[2]
	filterYearPublTo := filters[3]
	filterYearReadFrom := filters[4]
	filterYearReadTo := filters[5]
	filterRateFrom := filters[6]
	filterRateTo := filters[7]

	query := fmt.Sprintf(`SELECT
							rb.title, 
							a.author_name, 
							rb.year_public, 
							rb.year_read, 
							rb.rate
						FROM 
							rate_books rb
						JOIN 
							authors a ON rb.author_id = a.id
						WHERE
								CASE WHEN $1::text IS NULL THEN true ELSE rb.title ILIKE '%%' || $1::text || '%%' END
   							AND CASE WHEN $2::text IS NULL THEN true ELSE a.author_name ILIKE '%%' || $2::text || '%%' END
    						AND CASE WHEN $3::text IS NULL THEN true ELSE rb.year_public >= $3::integer END
    						AND CASE WHEN $4::text IS NULL THEN true ELSE rb.year_public <= $4::integer END
    						AND CASE WHEN $5::text IS NULL THEN true ELSE rb.year_read >= $5::integer END
    						AND CASE WHEN $6::text IS NULL THEN true ELSE rb.year_read <= $6::integer END
    						AND CASE WHEN $7::text IS NULL THEN true ELSE rb.rate >= $7::integer END
    						AND CASE WHEN $8::text IS NULL THEN true ELSE rb.rate <= $8::integer END
							AND CASE WHEN $11::integer IS NULL THEN true ELSE rb.book_owner = $11::integer END
						ORDER BY %s %s 
						OFFSET $9 LIMIT $10
						`, sort_field, sort_order)

	rows, err := DB.Query(
		query, filterByTitle, filterByAuthor, filterYearPublFrom, filterYearPublTo, filterYearReadFrom, filterYearReadTo, filterRateFrom, filterRateTo, page_number*page_size, page_size, us_id,
	)

	if err != nil {
		log.Println("Error in query SelectBooks", err)
		return nil, err
	}
	defer rows.Close()

	var Books []model.Book
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.Title, &book.Author.Author_name, &book.Year_public, &book.Year_read, &book.Rate)
		if err != nil {
			log.Println("Error in scan SelectBooks", err)
			return nil, err
		}
		Books = append(Books, book)
	}
	return Books, nil
}

// запрос на общее количество книг
func SelectAmountOfBooks(filters []interface{}, us_id int) (AmountOfBooks int, err error) {
	filterByTitle := filters[0]
	filterByAuthor := filters[1]
	filterYearPublFrom := filters[2]
	filterYearPublTo := filters[3]
	filterYearReadFrom := filters[4]
	filterYearReadTo := filters[5]
	filterRateFrom := filters[6]
	filterRateTo := filters[7]

	var total int
	query := fmt.Sprintln(
		`SELECT 
							COUNT(*)
       					FROM 
							rate_books rb
        				JOIN 
							authors a ON rb.author_id = a.id
       					WHERE
								CASE WHEN $1::text IS NULL THEN true ELSE rb.title ILIKE '%' || $1::text || '%' END
   							AND CASE WHEN $2::text IS NULL THEN true ELSE a.author_name ILIKE '%' || $2::text || '%' END
    						AND CASE WHEN $3::text IS NULL THEN true ELSE rb.year_public >= $3::integer END
    						AND CASE WHEN $4::text IS NULL THEN true ELSE rb.year_public <= $4::integer END
    						AND CASE WHEN $5::text IS NULL THEN true ELSE rb.year_read >= $5::integer END
    						AND CASE WHEN $6::text IS NULL THEN true ELSE rb.year_read <= $6::integer END
    						AND CASE WHEN $7::text IS NULL THEN true ELSE rb.rate >= $7::integer END
    						AND CASE WHEN $8::text IS NULL THEN true ELSE rb.rate <= $8::integer END
							AND CASE WHEN $9::integer IS NULL THEN true ELSE rb.book_owner = $9::integer END
						`)
	err = DB.QueryRow(query, filterByTitle, filterByAuthor, filterYearPublFrom, filterYearPublTo, filterYearReadFrom, filterYearReadTo, filterRateFrom, filterRateTo, us_id).Scan(&total)
	if err != nil {
		log.Println("Error in query SelectAmountOfBooks", err)
		return 0, err
	}

	AmountOfBooks = total

	return AmountOfBooks, nil
}
