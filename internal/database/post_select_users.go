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
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var uzer string
		err := rows.Scan(&uzer)
		if err != nil {
			log.Println(err)
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
		log.Println(err)
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
		log.Println("error:", err)
		return 0, "", err
	}

	defer rows.Close()

	if !rows.Next() {
		log.Println("user not find")
		return 0, "", err
	}

	var UserIdDB int
	var UserPassDB string
	if err := rows.Scan(&UserIdDB, &UserPassDB); err != nil {
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
		log.Println("error:", err)
		return false
	}

	defer rows.Close()

	if !rows.Next() {
		log.Println("user not find")
		return false
	}

	err = rows.Scan(&user_name)
	if err != nil {
		log.Println(err)
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
		log.Println("error:", err)
		return ""
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user_name)
		if err != nil {
			log.Println(err)
		}
	}

	return user_name
}
