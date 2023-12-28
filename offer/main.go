package main

import (
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
	dbConfig "offer/internal/db"
	models2 "offer/internal/models"
	interfaces "offer/pkg/v1"
	handler "offer/pkg/v1/handler/grpc"
	repo "offer/pkg/v1/repository"
	usecase "offer/pkg/v1/usecase"
)

func main() {
	db := dbConfig.Conn()
	migrations(db)

	lis, err := net.Listen("tcp", ":50002")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	OfferUseCase := initUserServer(db)
	handler.NewServer2(grpcServer, OfferUseCase)

	log.Fatalf("Failed to serve: %v", grpcServer.Serve(lis))
}

func initUserServer(db *gorm.DB) interfaces.UseCaseInterface {
	offerRepo := repo.New2(db)
	offerUseCase := usecase.New2(offerRepo)
	return offerUseCase
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models2.Offer{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migration did run successfully")
	}
}
