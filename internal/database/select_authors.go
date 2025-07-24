package database

import (
	"fmt"
	"log"
	"rate_books/internal/model"
)

// запрос на всех авторов из списка
func SelectAuthors(page_number int, page_size int, sort_field string, sort_order string, filters []interface{}, us_id int) ([]model.Authors, error) {
	filterByAuthor := filters[0]
	filterByCountry := filters[1]
	filterYearBFrom := filters[2]
	filterYearBTo := filters[3]

	query := fmt.Sprintf(`SELECT DISTINCT
							author_name, 
							year_b, 
							country
						FROM 
							authors a
						JOIN 
							rate_books rb ON rb.author_id = a.id
						WHERE
								CASE WHEN $1::text IS NULL THEN true ELSE a.author_name ILIKE '%%' || $1::text || '%%' END
							AND CASE WHEN $2::text IS NULL THEN true ELSE a.country ILIKE '%%' || $2::text || '%%' END
							AND CASE WHEN $3::text IS NULL THEN true ELSE a.year_b >= $3::integer END
    						AND CASE WHEN $4::text IS NULL THEN true ELSE a.year_b <= $4::integer END
							AND CASE WHEN $7::integer IS NULL THEN true ELSE rb.book_owner = $7::integer END
						ORDER BY %s %s 
						OFFSET $5 LIMIT $6
						`, sort_field, sort_order)

	rows, err := DB.Query(
		query, filterByAuthor, filterByCountry, filterYearBFrom, filterYearBTo, page_number*page_size, page_size, us_id,
	)

	if err != nil {
		log.Println("Error in query SelectAuthors", err)
		return nil, err
	}
	defer rows.Close()

	var Authors []model.Authors
	for rows.Next() {
		var Author model.Authors
		err := rows.Scan(&Author.Author_name, &Author.Year_born, &Author.Country)
		if err != nil {
			log.Println("Error in scan SelectAuthors", err)
			return nil, err
		}
		Authors = append(Authors, Author)
	}

	return Authors, nil
}

// запрос на общее количество авторов
func SelectAmountOfAuthors(filters []interface{}, us_id int) (AmountOfAuthors int, err error) {
	filterByAuthor := filters[0]
	filterByCountry := filters[1]
	filterYearBFrom := filters[2]
	filterYearBTo := filters[3]

	var total int
	query := fmt.Sprintf(`SELECT COUNT(DISTINCT a.id)
						FROM 
							authors a
						LEFT JOIN 
							rate_books rb ON rb.author_id = a.id
						WHERE
								CASE WHEN $1::text IS NULL THEN true ELSE a.author_name ILIKE '%%' || $1::text || '%%' END
							AND CASE WHEN $2::text IS NULL THEN true ELSE a.country ILIKE '%%' || $2::text || '%%' END
							AND CASE WHEN $3::text IS NULL THEN true ELSE a.year_b >= $3::integer END
    						AND CASE WHEN $4::text IS NULL THEN true ELSE a.year_b <= $4::integer END
							AND CASE WHEN $5::integer IS NULL THEN true ELSE rb.book_owner = $5::integer END
				 `)

	err = DB.QueryRow(query, filterByAuthor, filterByCountry, filterYearBFrom, filterYearBTo, us_id).Scan(&total)
	if err != nil {
		log.Println("Error in query SelectAmountOfAuthors", err)
		return 0, err
	}

	AmountOfAuthors = total

	return AmountOfAuthors, nil
}

// проверка на совпадения в списке авторов
func CheckAuthorsList(AuthorName string, us_id int) bool {
	query := `SELECT 
				author_name AS avtor
			FROM authors
							`

	rows, err := DB.Query(query)
	if err != nil {
		log.Println("Error in query CheckAuthorsList", err)
	}
	defer rows.Close()

	for rows.Next() {
		var avtor string
		err := rows.Scan(&avtor)
		if err != nil {
			log.Println("Error in scan CheckAuthorsList", err)
		}

		if avtor == AuthorName {
			log.Println("Такой автор уже есть")
			return true
		}
	}
	log.Println("Такого автора еще нет")
	return false
}
