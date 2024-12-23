package main

import (
	"belajar-crud-go/configs"
	"belajar-crud-go/controllers"
	"belajar-crud-go/migration"
	"belajar-crud-go/repositories"
	"belajar-crud-go/routes"
	"belajar-crud-go/services"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDatabase()
	migration.Migration()

	bookRepo := repositories.NewBookRepository(configs.DB)
	bookService := services.NewBookService(bookRepo)
	bookController := controllers.NewBookController(bookService)

	router := gin.Default()
	routes.BookRoutes(router, bookController)
	router.Run(":8080")
}
