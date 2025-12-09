package users

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

// получение пользователя по id
func (db *Repo) GetUserById(userID uuid.UUID) (User, error) {
	query := `
				SELECT id, name, email, pass
				FROM rate_books.users
				WHERE id = $1
			`
	var user User
	err := db.DB.QueryRow(query, userID).Scan(&user.Id, &user.Name, &user.Email, &user.Pass)
	if err != nil {
		return User{}, fmt.Errorf("error in GetUserById query: %w", err)
	}

	return user, nil
}

// получение пользователя по email
func (db *Repo) GetUserByEmail(userEmail string) (User, error) {
	query := `
				SELECT id, name, email, pass
				FROM rate_books.users
				WHERE email = $1
			`
	var user User
	err := db.DB.QueryRow(query, userEmail).Scan(&user.Id, &user.Name, &user.Email, &user.Pass)
	if err != nil {
		// Проверяем, это "запись не найдена" или реальная ошибка БД
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("No users with this email")
		}
		return User{}, fmt.Errorf("error in GetUserByEmail query: %w", err)
	}

	return user, nil
}

// проверка занятости имени пользователя
func (db *Repo) CheckUserName(userName string) (bool, error) {
	query := `
				SELECT EXISTS(
					SELECT 1
					FROM rate_books.users
					WHERE name = $1)
			`
	var exist bool

	err := db.DB.QueryRow(query, userName).Scan(&exist)
	if err != nil {
		return false, fmt.Errorf("error in CheckUserName query: %w", err)
	}

	return exist, nil
}

// проверка занятости почты
func (db *Repo) CheckUserEmail(userEmail string) (bool, error) {
	query := `
				SELECT EXISTS(
					SELECT 1
					FROM rate_books.users
					WHERE email = $1)
			`
	var exist bool

	err := db.DB.QueryRow(query, userEmail).Scan(&exist)
	if err != nil {
		return false, fmt.Errorf("error in CheckUserEmail query: %w", err)
	}

	return exist, nil
}
