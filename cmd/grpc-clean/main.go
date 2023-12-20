package main

import (
	dbConfig "Agency_management/internal/db"
	"Agency_management/internal/models"
	interfaces "Agency_management/pkg/v1"
	handler "Agency_management/pkg/v1/handler/grpc"
	repo "Agency_management/pkg/v1/repository"
	usecase "Agency_management/pkg/v1/usecase"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	db := dbConfig.Conn()
	migrations(db)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	agencyUseCase := initUserServer(db)
	handler.NewServer(grpcServer, agencyUseCase)

	log.Fatalf("Failed to serve: %v", grpcServer.Serve(lis))
}

func initUserServer(db *gorm.DB) interfaces.UseCaseInterface {
	agencyRepo := repo.New(db)
	agencyUseCase := usecase.New(agencyRepo)
	return agencyUseCase
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.Agency{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migration did run successfully")
	}
}
