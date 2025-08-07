package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rate_books/internal/jwtmod"
)

// вытаскиваем id пользователя из куков
func AuthCheck(c *gin.Context) (int, error) {
	value, err := c.Cookie("my_cookie")
	if err != nil {
		log.Println("Cookie not found")
		c.JSON(http.StatusForbidden, gin.H{"error": "Cookie not found"})
		return 0, err
	}

	myClaims := jwtmod.MyClaims{}

	token, err := jwtmod.ParseToken(value, &myClaims)
	if err != nil {
		log.Println("Err in parse token", err)
	}

	if !token.Valid {
		log.Println("Token not valid")
		c.JSON(http.StatusForbidden, gin.H{"error": "Token not valid"})
		return 0, err
	}

	return myClaims.UserId, nil
}
