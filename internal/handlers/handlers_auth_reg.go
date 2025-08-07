package handlers

import (
	"log"
	"net/http"
	"rate_books/internal/jwtmod"
	"rate_books/internal/database"
	"rate_books/internal/model"
	"github.com/gin-gonic/gin"
)

// добавление нового юзера
func NewUser(c *gin.Context) {
	// получение структуры "пользователь" с фронта
	var NewUser model.User
	if err := c.ShouldBindJSON(&NewUser); err != nil {
		log.Println("Error in ShouldBindJSON", err)
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}

	UserName := NewUser.UserName

	// проверка свободности имени пользователя
	if !database.CheckUsersList(UserName) {
		log.Println("Error in CheckUsersList")
		c.JSON(http.StatusOK, gin.H{"check_user": false})
		return
	}

	// добавление нового пользователя в БД
	user_id, err := database.UserInsert(NewUser)
	if err != nil {
		log.Println("Error in UserInsert", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// генерация токена для нового пользователя
	token, err := jwtmod.GenerateToken(user_id, UserName)
    if err != nil {
        log.Println("Error in token generation", err)
        return
    }

	// установка куков новому пользователю
	c.SetCookie("my_cookie", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "succesful"})
}

// проверка старого юзера
func OldUser(c *gin.Context) {
	// получение структуры "пользователь" с фронта
	var OldUser model.User
	if err := c.ShouldBindJSON(&OldUser); err != nil {
		log.Println("Error in ShouldBindJSON", err)
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}

	UserNameInput := OldUser.UserName
	UserPassInput := OldUser.Pass

	// поиск пользователя в БД по заданному имени
	UserIdDB, UserPassDB, err := database.SelectUserName(UserNameInput)
	if err != nil {
		log.Println("Error in SelectUserName", err)
		c.JSON((http.StatusForbidden), gin.H{"error": "error in select"})
		return
	}

	// проверка пароля введенного и из БД
	if UserPassInput != UserPassDB {
		log.Println("Password dont match")
		c.JSON((http.StatusForbidden), gin.H{"error": "passwords dont match"})
		return
	}

	// генерация токена пользователю
	token, err := jwtmod.GenerateToken(UserIdDB, UserNameInput)
    if err != nil {
        log.Println("Error in token generation", err)
        return
    }

	// установка куков
	c.SetCookie("my_cookie", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "succesful"})
}

// проверка залогиненности
func CheckAut(c *gin.Context) {
	// получение значение куков
	value, err := c.Cookie("my_cookie")
	if err != nil {
		log.Println("Cookie not found")
		c.JSON(http.StatusForbidden, gin.H{"error": "Cookie not found"})
		return
	}

	// место, куда парсер распарсит клеймы
	myClaims := jwtmod.MyClaims {}

	// получаем токен
	token, err := jwtmod.ParseToken(value, &myClaims)
	if err != nil {
		log.Println("Err in parse token", err)
	}

	// проверка валидности токена
	if !token.Valid {
		log.Println("Token not valid")
		c.JSON(http.StatusForbidden, gin.H{"error": "Token not valid"})
		return
	}

	// выбор пользователя с заданным id 
	if !database.SelectUserId(myClaims.UserId) {
		log.Println("Error in SelectUserID")
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid auth"})
		return
	}

	user_name := database.NameById(myClaims.UserId)

	c.JSON(http.StatusOK, gin.H{"user_name": user_name})
}

// лог аут
func LogOut(c *gin.Context) {
	c.SetCookie("my_cookie", "1", -1, "/", "", false, false)
}