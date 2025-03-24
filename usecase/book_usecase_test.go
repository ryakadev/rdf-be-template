package usecase

import (
	"testing"

	"github.com/FRFebi/template-service/domain"
	"github.com/FRFebi/template-service/infrastructure"
	"github.com/FRFebi/template-service/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {

	db := infrastructure.ConnectDB()
	bookRepo := repository.NewBookRepository(db)
	ucBook := NewBookUsecase(bookRepo)
	book := &domain.Book{
		Name: "Science",
	}
	t.Run("Create a new Book with mock", func(t *testing.T) {
		book, err := ucBook.CreateBook(book)
		assert.Nil(t, err)
		assert.NotNil(t, book)
	})
}

func TestShowBook(t *testing.T) {

	db := infrastructure.ConnectDB()
	bookRepo := repository.NewBookRepository(db)
	bookUC := NewBookUsecase(bookRepo)

	t.Run("Show a Book", func(t *testing.T) {
		books, err := bookUC.ShowBooks()
		assert.Nil(t, err)
		assert.NotNil(t, books)
	})
}

func TestUpdateBook(t *testing.T) {
	db := infrastructure.ConnectDB()
	bookRepo := repository.NewBookRepository(db)
	bookUC := NewBookUsecase(bookRepo)
	book := &domain.Book{
		Name: "Science X",
	}

	t.Run("Update a Book", func(t *testing.T) {
		book, err := bookUC.CreateBook(book)
		assert.Nil(t, err)

		book.Name = "Science G"
		_, err = bookUC.UpdateBook(book)
		assert.Nil(t, err)
	})
}
