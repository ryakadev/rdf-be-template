package main

import (
	"log"
	"net"
	"os"

	"github.com/FRFebi/template-service/delivery/grpc"
	"github.com/FRFebi/template-service/domain"
	"github.com/FRFebi/template-service/infrastructure"
	"github.com/FRFebi/template-service/proto"
	"github.com/FRFebi/template-service/repository"
	"github.com/FRFebi/template-service/usecase"
	gogrpc "google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	db := infrastructure.ConnectDB()

	if err := db.AutoMigrate(&domain.Book{}); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	bookRepo := repository.NewBookRepository(db)
	bookUC := usecase.NewBookUsecase(bookRepo)
	bookGRPC := grpc.NewBookGRPC(bookUC)

	gRPCServer := gogrpc.NewServer()
	proto.RegisterBookServiceServer(gRPCServer, bookGRPC)

	log.Println("gRPC server is running on port 9000")
	err = gRPCServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
