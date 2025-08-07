package jwtmod

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

// создание структуры с кастомными и стандартными клеймами
type MyClaims struct {
	UserId int `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

// метод для клеймов - валидация кастомных клеймов
func (c *MyClaims) Validate() error {
	if c.UserId <= 0 {
		return errors.New("invalid user_id")
	}
	if c.UserName == "" {
		return errors.New("invalid user_name")
	}
	return nil
}