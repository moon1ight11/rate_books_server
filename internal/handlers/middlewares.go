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
		log.Println("Cookie not found")
		c.JSON(http.StatusForbidden, gin.H{"error": "Cookie not found"})
		return 0, err
	}

	id, err := strconv.Atoi(value)
	if err != nil {
		log.Println("Error in convert cookie")
		return 0, err
	}

	return id, nil
}

