package repository

import (
	"testing"
	"time"

	"github.com/FRFebi/template-service/domain"
	"github.com/FRFebi/template-service/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repoBookMock struct {
	mock.Mock
}

func NewRepoBookMock() domain.BookRepository {
	return &repoBookMock{}
}

func (r *repoBookMock) Create(book *domain.Book) (*domain.Book, error) {
	args := r.Called(book)
	return args.Get(0).(*domain.Book), args.Error(1)
}

func (r *repoBookMock) FindAll() ([]*domain.Book, error) {
	args := r.Called()
	return args.Get(0).([]*domain.Book), args.Error(1)
}

func (r *repoBookMock) FindById(id string) (*domain.Book, error) {
	args := r.Called(id)
	return args.Get(0).(*domain.Book), args.Error(1)
}

func (r *repoBookMock) Update(book *domain.Book) (*domain.Book, error) {
	args := r.Called(book)
	return args.Get(0).(*domain.Book), args.Error(1)
}

func (r *repoBookMock) Delete(book *domain.Book) error {
	args := r.Called(book)
	return args.Error(0)
}

func TestCreateBook(t *testing.T) {
	repoBookMock := repoBookMock{}
	book := &domain.Book{Name: "History"}
	repoBookMock.On("Create", book).Return(nil)

	t.Run("Create a new Book with mock", func(t *testing.T) {
		book, err := repoBookMock.Create(book)
		assert.Nil(t, err)
		assert.NotNil(t, book)
	})
	db := infrastructure.ConnectDB()
	repoBook := NewBookRepository(db)

	t.Run("Create a new Book to DB", func(t *testing.T) {
		book, err := repoBook.Create(book)
		assert.NotNil(t, book.Id)
		assert.Nil(t, err)
	})
}

func TestShowBook(t *testing.T) {

	db := infrastructure.ConnectDB()
	bookRepo := NewBookRepository(db)
	t.Run("Show a Book", func(t *testing.T) {
		books, err := bookRepo.FindAll()
		assert.Nil(t, err)
		assert.NotNil(t, books)
	})
}

func TestUpdateBook(t *testing.T) {

	repoBookMock := repoBookMock{}
	book := &domain.Book{
		Name: "Wrong",
	}
	now := time.Now()
	CreateBookReponse := &domain.Book{
		Id:       "1",
		Name:     "Wrong",
		CreateAt: &now,
	}
	repoBookMock.On("Create", book).Return(CreateBookReponse, nil)
	CreateBookReponse = &domain.Book{
		Id:       "1",
		Name:     "Correction",
		CreateAt: &now,
	}
	now = time.Now()
	UpdateBookResponse := &domain.Book{
		Id:       "1",
		Name:     "Correction",
		CreateAt: &now,
		UpdateAt: &now,
	}
	repoBookMock.On("Update", CreateBookReponse).Return(UpdateBookResponse, nil)
	t.Run("Update a Book with mock", func(t *testing.T) {

		book, err := repoBookMock.Create(book)
		assert.Nil(t, err)
		assert.NotNil(t, book)

		book, err = repoBookMock.Update(CreateBookReponse)
		assert.Nil(t, err)
		assert.Equal(t, "Correction", book.Name)
	})

	db := infrastructure.ConnectDB()
	repoBook := NewBookRepository(db)

	t.Run("Update a Book to DB", func(t *testing.T) {
		book, err := repoBook.Create(book)
		assert.Nil(t, err)
		assert.NotNil(t, book)

		book.Name = "Correction"
		book, err = repoBook.Update(book)
		assert.Nil(t, err)
		assert.Equal(t, "Correction", book.Name)
	})
}

func TestDeleteBook(t *testing.T) {

	db := infrastructure.ConnectDB()
	repoBook := NewBookRepository(db)
	book := &domain.Book{
		Name: "Buku Baru",
	}

	t.Run("Delete a Book", func(t *testing.T) {
		book, err := repoBook.Create(book)
		assert.Nil(t, err)

		err = repoBook.Delete(book)
		assert.Nil(t, err)
	})
}
