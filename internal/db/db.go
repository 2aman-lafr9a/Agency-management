package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Conn() *gorm.DB {
	db, err := gorm.Open(
		postgres.Open("host=localhost user=postgres password=password dbname=agency_management port=5432"), &gorm.Config{},
	)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
