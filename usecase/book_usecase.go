package usecase

import (
	"time"

	"github.com/FRFebi/template-service/domain"
)

type BookUsecase struct {
	BookRepository domain.BookRepository
}

func NewBookUsecase(repo domain.BookRepository) domain.BookUsecase {
	return &BookUsecase{BookRepository: repo}
}

func (u *BookUsecase) CreateBook(book *domain.Book) (*domain.Book, error) {
	return u.BookRepository.Create(book)
}

func (u *BookUsecase) ShowBooks() ([]*domain.Book, error) {
	return u.BookRepository.FindAll()
}

func (u *BookUsecase) UpdateBook(book *domain.Book) (*domain.Book, error) {
	oldBook, err := u.BookRepository.FindById(book.Id)
	if err != nil || oldBook == nil {
		return nil, err
	}
	now := time.Now()
	oldBook.Name = book.Name
	oldBook.UpdateAt = &now

	return u.BookRepository.Update(oldBook)
}

func (u *BookUsecase) DeleteBook(book *domain.Book) error {
	_, err := u.BookRepository.FindById(book.Id)
	if err != nil {
		return err
	}
	return u.BookRepository.Delete(book)
}
