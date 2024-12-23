package routes

import (
	"belajar-crud-go/controllers"

	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine, bookController *controllers.BookController) {
	books := router.Group("/books")
	{
		books.GET("", bookController.GetAllBooks)
		books.GET("/:id", bookController.GetBookById)
		books.POST("", bookController.CreateBook)
		books.PUT("/:id", bookController.UpdateBook)
		books.DELETE("/:id", bookController.DeleteBook)
	}
}
