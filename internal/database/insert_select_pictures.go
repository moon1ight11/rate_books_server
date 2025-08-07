package database

import (
	"log"
)

// добавление картинки в БД
func InsertIMG(cover_name string) int {
	query := `	INSERT INTO covers (original_name, created_at) 
				VALUES ($1, CURRENT_TIMESTAMP)
				RETURNING id`

	var c_id int
	err := DB.QueryRow(query, cover_name).Scan(&c_id)
	if err != nil {
		log.Println("Error in query InsertIMG:", err)
		return 0
	}
	return c_id
}

// название картинки по id
func SelectNameIMGByID(cover_id int) (string, error) {
	query := `SELECT original_name FROM covers WHERE id = $1 `
	
	var cover_name string

	err := DB.QueryRow(query, cover_id).Scan(&cover_name)
	if err != nil {
		log.Println("Error in SelectNameIMGByID query", err)
		return "", err
	}

	return cover_name, err
}