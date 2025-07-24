package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rate_books/internal/database"
	"rate_books/internal/model"
	"strconv"
)

// добавление нового юзера
func NewUser(c *gin.Context) {
	var NewUser model.User
	if err := c.ShouldBindJSON(&NewUser); err != nil {
		log.Println("Error in ShouldBindJSON", err)
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}

	UserName := NewUser.UserName

	if !database.CheckUsersList(UserName) {
		log.Println("Error in CheckUsersList")
		c.JSON(http.StatusOK, gin.H{"check_user": false})
		return
	}

	user_id, err := database.UserInsert(NewUser)
	if err != nil {
		log.Println("Error in UserInsert", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	id := strconv.Itoa(user_id)

	c.SetCookie("my_cookie", id, 3600, "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{"message": "succesful"})
}

// проверка старого юзера
func OldUser(c *gin.Context) {
	var OldUser model.User
	if err := c.ShouldBindJSON(&OldUser); err != nil {
		log.Println("Error in ShouldBindJSON", err)
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}

	UserNameInput := OldUser.UserName
	UserPassInput := OldUser.Pass

	UserIdDB, UserPassDB, err := database.SelectUserName(UserNameInput)
	if err != nil {
		log.Println("Error in SelectUserName", err)
		c.JSON((http.StatusForbidden), gin.H{"error": "error in select"})
		return
	}

	if UserPassInput != UserPassDB {
		log.Println("Password dont match")
		c.JSON((http.StatusForbidden), gin.H{"error": "passwords dont match"})
		return
	}

	id := strconv.Itoa(UserIdDB)

	c.SetCookie("my_cookie", id, 3600, "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{"message": "succesful"})
}

// проверка залогиненности
func CheckAut(c *gin.Context) {
	value, err := c.Cookie("my_cookie")
	if err != nil {
		log.Println("Cookie not found")
		c.JSON(http.StatusForbidden, gin.H{"error": "Cookie not found"})
		return
	}

	id, err := strconv.Atoi(value)
	if err != nil {
		log.Println("Err in convert cookie")
		return
	}

	if !database.SelectUserId(id) {
		log.Println("Error in SelectUserID")
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid auth"})
		return
	}

	user_name := database.NameById(id)

	c.JSON(http.StatusOK, gin.H{"user_name": user_name})
}

// лог аут
func LogOut(c *gin.Context) {
	c.SetCookie("my_cookie", "1", -1, "/", "", false, false)
}