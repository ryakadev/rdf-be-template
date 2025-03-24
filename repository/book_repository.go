package repository

import (
	"time"

	"github.com/FRFebi/template-service/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) domain.BookRepository {
	return &BookRepository{db}
}

func (r *BookRepository) Create(book *domain.Book) (*domain.Book, error) {
	now := time.Now()
	newBook := &domain.Book{
		Id:       uuid.NewString(),
		Name:     book.Name,
		CreateAt: &now,
	}
	err := r.db.Create(newBook).Error
	if err != nil {
		return nil, err
	}
	return newBook, nil
}

func (r *BookRepository) FindAll() ([]*domain.Book, error) {
	var books []*domain.Book
	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) FindById(id string) (*domain.Book, error) {
	var book domain.Book
	err := r.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) Update(book *domain.Book) (*domain.Book, error) {
	err := r.db.Save(book).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookRepository) Delete(book *domain.Book) error {
	err := r.db.Delete(book).Error
	if err != nil {
		return err
	}
	return nil
}
