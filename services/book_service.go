package services

import (
	"belajar-crud-go/models"
	"belajar-crud-go/repositories"

	"github.com/davecgh/go-spew/spew"
)

type BookService interface {
	GetAllBooks() ([]models.Book, error)
	CreateBook(book models.Book) (models.Book, error)
	GetBookById(id int) (models.Book, error)
	UpdateBook(book models.Book) (models.Book, error)
	DeleteBook(book models.Book) (models.Book, error)
}

type bookService struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{repo}
}

func (s *bookService) GetAllBooks() ([]models.Book, error) {
	return s.repo.GetAll()
}

func (s *bookService) CreateBook(book models.Book) (models.Book, error) {
	return s.repo.Create(book)
}

func (s *bookService) GetBookById(id int) (models.Book, error) {
	spew.Dump("Value of id:", id)
	return s.repo.GetById(uint(id))
}

func (s bookService) UpdateBook(book models.Book) (models.Book, error) {
	return s.repo.Update(book)
}

func (s bookService) DeleteBook(book models.Book) (models.Book, error) {
	return s.repo.Delete(book)
}
