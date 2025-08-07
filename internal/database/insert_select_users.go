package database

import (
	"log"
	"rate_books/internal/model"
)

// проверка свободности имени пользователя
func CheckUsersList(UserName string) bool {
	query :=
			`SELECT 
				user_name AS uzer
			FROM 
				users`

	rows, err := DB.Query(query)

	if err != nil {
		log.Println("Error in query CheckUserList:", err)
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var uzer string
		err := rows.Scan(&uzer)
		if err != nil {
			log.Println("Error in Scan CheckUserList", err)
			return false
		}

		if uzer == UserName {
			log.Println("Имя пользователя занято")
			return false
		}
	}
	log.Println("Имя пользователя свободно")
	return true
}

// добавление нового юзера и возврат его id
func UserInsert(NewUser model.User) (int, error) {
	var user_id int
	query := `	INSERT INTO users (user_name, user_password, time_stamp) 
				VALUES ($1, $2, CURRENT_TIMESTAMP) 
				RETURNING id`

	err := DB.QueryRow(query, NewUser.UserName, NewUser.Pass).Scan(&user_id)
	if err != nil {
		log.Println("Error in query UserInsert:", err)
		return 0, err
	}

	return user_id, nil
}

// выбор пользователя с заданным именем
func SelectUserName(UserName string) (int, string, error) {
	query :=
		`SELECT
				id,
				user_password
			FROM
				users
			WHERE user_name = $1`

	rows, err := DB.Query(
		query, UserName,
	)
	if err != nil {
		log.Println("Error in query SelectUserName:", err)
		return 0, "", err
	}

	defer rows.Close()

	if !rows.Next() {
		log.Println("User not found")
		return 0, "", err
	}

	var UserIdDB int
	var UserPassDB string
	if err := rows.Scan(&UserIdDB, &UserPassDB); err != nil {
		log.Println("Error in Scan SelectUserName:", err)
		return 0, "", err
	}

	return UserIdDB, UserPassDB, nil
}

// выбор пользователя с заданным id
func SelectUserId(UserId int) bool {
	var user_name string
	query :=
		`SELECT user_name FROM users WHERE id = $1`

	rows, err := DB.Query(
		query, UserId,
	)
	if err != nil {
		log.Println("Error in SelectUserId query", err)
		return false
	}

	defer rows.Close()

	if !rows.Next() {
		log.Println("User not found")
		return false
	}

	err = rows.Scan(&user_name)
	if err != nil {
		log.Println("Error in Scan SelectUserID:", err)
		return false
	}

	return true
}

// имя пользователя по id
func NameById(UserId int) string {
	var user_name string
	query :=
		`SELECT user_name FROM users WHERE id = $1`

	rows, err := DB.Query(
		query, UserId,
	)
	if err != nil {
		log.Println("Error in NameByID query", err)
		return ""
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user_name)
		if err != nil {
			log.Println("Error in Scan NameByID", err)
			return ""
		}
	}

	return user_name
}
