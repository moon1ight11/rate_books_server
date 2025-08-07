package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// gin-роутер
func Router() {
	// инициализация gin
	r := gin.Default()

	// разрешения для CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// устанавливаем максимальный размер входящего файла
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	// принимаем файл с фронта
	r.POST("/upload", AddPicture)
	// отправляем файл на фронт
	r.GET("/image/:id", GetPicture)

	// роуты регистрации и авторизации
	r.POST("/user/register", NewUser)
	r.POST("/user/login", OldUser)
	r.GET("/user/auth_check", CheckAut)
	r.GET("/user/log_out", LogOut)

	// get для книг
	r.GET("/all_books", GetAllBooks)
	r.GET("/rec_books", BooksRecomm)

	// get для авторов
	r.GET("/all_authors", GetAllAuthors)

	// post
	r.POST("/new_book", PostNewBook)
	r.POST("/new_author", PostNewAuthor)

	// запуск сервера
	r.Run(":8080")
}
