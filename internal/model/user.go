package model

// модель юзера
type User struct {
	UserName string `json:"user_name"`
	Pass     string `json:"pass"`
}