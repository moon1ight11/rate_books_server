package handlers

import (
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// вытаскиваем id пользователя из куков
func AuthCheck(c *gin.Context) (int, error) {
	value, err := c.Cookie("my_cookie")
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cookie not found"})
		log.Println("net kukov")
		return 0, err
	}

	id, err := strconv.Atoi(value)
	if err != nil {
		log.Println("kuki plohie")
		return 0, err
	}

	return id, nil
}

