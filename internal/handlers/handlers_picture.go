package handlers

import (
	"log"
	"net/http"
	"rate_books/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

// добавление картинки в БД
func AddPicture(c *gin.Context) {
	file, _ := c.FormFile("file")

	err := c.SaveUploadedFile(file, "./files/covers/"+file.Filename)
	if err != nil {
		log.Println("Error in SaveUploadedFile", err)
		return
	}

	file_name := file.Filename

	c_id := database.InsertIMG(file_name)

	c.JSON(http.StatusOK, gin.H{"c_id": c_id})
}

// отправление картинки на фронт
func GetPicture(c *gin.Context) {
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error in convert cover_id(GetPicture)", err)
		return
	}

	cover_name, err := database.SelectNameIMGByID(id)
	if err != nil {
		log.Println("Error in SelectNameIMGByID(GetPicture)", err)
		return
	}

	filePath := "./files/covers/" + cover_name

	c.File(filePath)
}
