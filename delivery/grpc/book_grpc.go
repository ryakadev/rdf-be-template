package grpc

import (
	"context"

	"github.com/FRFebi/template-service/domain"
	"github.com/FRFebi/template-service/helper"
	"github.com/FRFebi/template-service/proto"
)

type BookGRPC struct {
	proto.UnimplementedBookServiceServer
	BookUsecase domain.BookUsecase
}

func NewBookGRPC(usecase domain.BookUsecase) *BookGRPC {
	return &BookGRPC{BookUsecase: usecase}
}

func (g *BookGRPC) CreateBook(ctx context.Context, req *proto.CreateBookRequest) (*proto.Book, error) {
	book := &domain.Book{
		Name: req.Name,
	}

	book, err := g.BookUsecase.CreateBook(book)
	if err != nil {
		return nil, err
	}

	res := &proto.Book{
		Id:        book.Id,
		Name:      book.Name,
		CreatedAt: helper.SafeTimeString(book.CreateAt),
		UpdatedAt: helper.SafeTimeString(book.UpdateAt),
		DeletedAt: helper.SafeTimeString(book.DeletedAt),
	}

	return res, nil
}

func (g *BookGRPC) ShowBooks(ctx context.Context, req *proto.EmptyRequest) (*proto.ShowBookResponse, error) {

	books, err := g.BookUsecase.ShowBooks()
	if err != nil {
		return nil, err
	}

	res := &proto.ShowBookResponse{Books: make([]*proto.Book, 0, len(books))}
	for _, v := range books {

		book := &proto.Book{
			Id:        v.Id,
			Name:      v.Name,
			CreatedAt: helper.SafeTimeString(v.CreateAt),
			UpdatedAt: helper.SafeTimeString(v.UpdateAt),
			DeletedAt: helper.SafeTimeString(v.DeletedAt),
		}
		res.Books = append(res.Books, book)
	}

	return res, nil
}

func (g *BookGRPC) UpdateBook(ctx context.Context, req *proto.UpdateBookRequest) (*proto.Book, error) {
	book := &domain.Book{
		Id:   req.Id,
		Name: req.Name,
	}

	book, err := g.BookUsecase.UpdateBook(book)
	if err != nil {
		return nil, err
	}

	res := &proto.Book{
		Id:        book.Id,
		Name:      book.Name,
		CreatedAt: helper.SafeTimeString(book.CreateAt),
		UpdatedAt: helper.SafeTimeString(book.UpdateAt),
		DeletedAt: helper.SafeTimeString(book.DeletedAt),
	}

	return res, nil
}

func (g *BookGRPC) DeleteBook(ctx context.Context, req *proto.DeleteBookRequest) (*proto.EmptyResponse, error) {
	book := &domain.Book{
		Id: req.Id,
	}

	err := g.BookUsecase.DeleteBook(book)
	if err != nil {
		return nil, err
	}

	return &proto.EmptyResponse{}, nil
}
