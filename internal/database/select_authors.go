package database

import (
	"fmt"
	"log"
	"rate_books/internal/model"
	// "rate_books/internal/handlers"
	// "github.com/gin-gonic/gin"
)

// запрос на всех авторов из списка
func SelectAuthors(page_number int, page_size int, where string, sort_field string, sort_order string, args []interface{}) ([]model.Authors, error) {
	queryArgs := make([]interface{}, len(args))
	copy(queryArgs, args)
	queryArgs = append(queryArgs, page_number*page_size, page_size)

	query := fmt.Sprintf(`SELECT 
							author_name, 
							year_b, 
							country
						FROM authors
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

	var Authors []model.Authors
	for rows.Next() {
		var Author model.Authors
		err := rows.Scan(&Author.Author_name, &Author.Year_born, &Author.Country)
		if err != nil {
			return nil, err
		}
		Authors = append(Authors, Author)
	}

	return Authors, nil
}

// запрос на общее количество авторов
func SelectAmountOfAuthors(where string, args []interface{}) (AmountOfAuthors int, err error) {
	var total int
	query := fmt.Sprintf(`
				SELECT count(*) 
				FROM authors
				 %s`, where,
	)

	err = DB.QueryRow(query, args...).Scan(&total)
	if err != nil {
		return 0, err
	}

	AmountOfAuthors = total

	return AmountOfAuthors, nil
}

// проверка на совпадения в списке авторов
func CheckAuthorsList(AuthorName string) bool {

	query := `SELECT author_name AS avtor
							FROM authors
							`

	rows, err := DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var avtor string
		err := rows.Scan(&avtor)
		if err != nil {
			log.Fatal(err)
		}

		if avtor == AuthorName {
			log.Println("Такой автор уже есть")
			return true
		}
	}
	log.Println("Такого автора еще нет")
	return false
}
