package main

import (
	"github.com/idj1997/book-rent-core/config"
	"github.com/idj1997/book-rent-core/repository"
	"github.com/idj1997/book-rent-core/service"
	"github.com/idj1997/book-rent-grpc/api"
	"github.com/idj1997/book-rent-grpc/proto/proto_book"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	config.InitConfig("dev", "config.yml")
	db := config.OpenPostgresDB()

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	bookRepo := repository.GormBookRepository{Db: db}
	bookService := service.NewBookService(&bookRepo)

	s := grpc.NewServer()
	proto_book.RegisterBookServiceServer(s, &api.BookServer{Service: *bookService})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
