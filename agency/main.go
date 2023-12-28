package main

import (
	dbConfig "agency/internal/db"
	"agency/internal/models"
	interfaces "agency/pkg/v1"
	handler "agency/pkg/v1/handler/grpc"
	repo "agency/pkg/v1/repository"
	usecase "agency/pkg/v1/usecase"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	db := dbConfig.Conn()
	migrations(db)

	lis, err := net.Listen("tcp", ":50001")
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
