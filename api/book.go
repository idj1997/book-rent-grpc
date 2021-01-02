package api

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/idj1997/book-rent-core/domain"
	"github.com/idj1997/book-rent-core/service"
	"github.com/idj1997/book-rent-grpc/mappings"
	"github.com/idj1997/book-rent-grpc/proto/proto_book"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookServer struct {
	proto_book.UnimplementedBookServiceServer
	Service service.BookService
}

func (b *BookServer) GetBookByID(ctx context.Context, request *proto_book.GetBookByIDRequest) (*proto_book.Book, error) {
	book, err := b.Service.GetByID(int(request.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Book not found.")
	}
	return mappings.BookToProtoBook(book), nil
}

func (b *BookServer) CreateBook(ctx context.Context, request *proto_book.CreateBookRequest) (*proto_book.Book, error) {
	book := domain.Book{
		Title:   request.Title,
		Content: request.Content,
		Stock:   int(request.Stock)}

	_, err := b.Service.Create(&book)
	if err != nil {
		serviceErr := mappings.ErrorToServiceError(err)
		if serviceErr.Type == service.InvalidArguments {
			return nil, status.Errorf(codes.FailedPrecondition, "Invalid book data.")
		} else {
			return nil, status.Errorf(codes.Internal, "")
		}
	}
	return mappings.BookToProtoBook(&book), nil
}

func (b *BookServer) UpdateBookStock(ctx context.Context, request *proto_book.UpdateBookStockRequest) (*proto_book.Book, error) {
	book, err := b.Service.UpdateStock(int(request.Id), int(request.NewStock))
	if err != nil {
		serviceErr := mappings.ErrorToServiceError(err)
		if serviceErr.Type == service.NotFound {
			return nil, status.Errorf(codes.Unknown, "Book not found.")
 		} else if serviceErr.Type == service.InvalidArguments{
 			return nil, status.Errorf(codes.FailedPrecondition, "Invalid stock value.")
		} else {
			return nil, status.Errorf(codes.Internal, "")
		}
	}
	return mappings.BookToProtoBook(book), nil
}

func (b *BookServer) DeleteBook(ctx context.Context, request *proto_book.DeleteBookRequest) (*empty.Empty, error) {
	err := b.Service.Delete(int(request.Id))
	if err != nil {
		serviceErr := mappings.ErrorToServiceError(err)
		if serviceErr.Type == service.NotFound {
			return nil, status.Errorf(codes.NotFound, "Book not found.")
		} else {
			return nil, status.Errorf(codes.Internal, "")
		}
	}
	return &empty.Empty{}, nil
}
