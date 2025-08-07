package jwtmod

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("112900112900")

// создание токена по id и имени
func GenerateToken(UserID int, UserName string) (string, error) {
	// обозначаем клеймы
	claims := MyClaims {
		UserId: UserID,
		UserName: UserName,
	}

	// устанавливаем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// декодировка токена
func ParseToken(tokenString string, myClaims *MyClaims) (*jwt.Token, error) {
    return jwt.ParseWithClaims(tokenString, myClaims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid method")
		}
        return secretKey, nil
    })
}

