package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {
	// инициализация gin
	r := gin.Default()

	//  разрешения для CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// get для книг
	r.GET("/all_books", GetAllBooks)

	// get для авторов
	r.GET("/all_authors", GetAllAuthors)

	// post
	r.POST("/new_book", PostNewBook)
	r.POST("/new_author", PostNewAuthor)

	// запуск сервера
	r.Run(":8080")
}
