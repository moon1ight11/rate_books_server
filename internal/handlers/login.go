package handlers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserName string `json:"user_name"`
	Pass     string `json:"pass"`
}

func Login(c *gin.Context) {
	var User User
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON((http.StatusUnauthorized), gin.H{"error": err.Error()})
		return
	}

	my_cookie, err := c.Cookie("my_cookie")
	if err != nil {
		log.Println(err)
	}

	log.Println(my_cookie)


	c.SetCookie("my_cookie", User.Pass, 3600, "/", "", false, false)
	c.JSON(http.StatusOK, "")
}
