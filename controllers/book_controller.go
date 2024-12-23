package controllers

import (
	"belajar-crud-go/helpers"
	"belajar-crud-go/models"
	"belajar-crud-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	service services.BookService
}

func NewBookController(service services.BookService) *BookController {
	return &BookController{service}
}

func (c *BookController) GetAllBooks(ctx *gin.Context) {
	books, err := c.service.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseSuccess(books, "Berhasil Mengambil Data Buku", http.StatusOK))
}

func (c *BookController) CreateBook(ctx *gin.Context) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBook, err := c.service.CreateBook(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, helpers.ResponseSuccess(createdBook, "Berhasil Menambahkan Data Buku", http.StatusCreated))
}

func (c *BookController) GetBookById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	book, err := c.service.GetBookById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, helpers.ResponseSuccess(book, "Berhasil Mengambil Data Buku", http.StatusOK))
}

func (c *BookController) UpdateBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = uint(id)

	updatedBook, err := c.service.UpdateBook(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedBook)
}

func (c *BookController) DeleteBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = uint(id)

	deletedBook, err := c.service.DeleteBook(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, deletedBook)
}
