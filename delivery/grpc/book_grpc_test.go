package grpc

import (
	"context"
	"testing"

	"github.com/FRFebi/template-service/domain"
	"github.com/FRFebi/template-service/infrastructure"
	"github.com/FRFebi/template-service/proto"
	"github.com/FRFebi/template-service/repository"
	"github.com/FRFebi/template-service/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BookUsecaseMock struct {
	mock.Mock
}

func (m *BookUsecaseMock) CreateBook(book *domain.Book) (*domain.Book, error) {
	args := m.Called(book)
	return args.Get(0).(*domain.Book), args.Error(1)
}

func (m *BookUsecaseMock) ShowBooks() ([]*domain.Book, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Book), args.Error(1)
}

func TestBookHandler(t *testing.T) {
	db := infrastructure.ConnectDB()
	bookRepo := repository.NewBookRepository(db)
	bookUC := usecase.NewBookUsecase(bookRepo)
	bookGRPC := NewBookGRPC(bookUC)

	book := &domain.Book{
		Name: "History",
	}
	t.Run("Create a new Book with mock", func(t *testing.T) {
		res, err := bookGRPC.CreateBook(context.Background(), &proto.CreateBookRequest{Name: book.Name})
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
	t.Run("Show a Book", func(t *testing.T) {
		res, err := bookGRPC.ShowBooks(context.Background(), &proto.EmptyRequest{})
		assert.Nil(t, err)
		assert.NotNil(t, res)
		book.Id = res.Books[0].Id
	})

	t.Run("Update a Book", func(t *testing.T) {
		res, err := bookGRPC.UpdateBook(context.Background(), &proto.UpdateBookRequest{Id: book.Id, Name: "World War II"})
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Delete a Book", func(t *testing.T) {
		res, err := bookGRPC.UpdateBook(context.Background(), &proto.UpdateBookRequest{Id: book.Id, Name: "World War II"})
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

}
