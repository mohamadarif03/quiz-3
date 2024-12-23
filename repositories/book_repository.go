package repositories

import (
	"belajar-crud-go/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll() ([]models.Book, error)
	Create(book models.Book) (models.Book, error)
	GetById(id uint) (models.Book, error)
	Update(book models.Book) (models.Book, error)
	Delete(book models.Book) (models.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) Create(book models.Book) (models.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *bookRepository) GetById(id uint) (models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	return book, err
}

func (r *bookRepository) Update(book models.Book) (models.Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}

func (r *bookRepository) Delete(book models.Book) (models.Book, error) {
	err := r.db.Delete(&book).Error
	return book, err
}
