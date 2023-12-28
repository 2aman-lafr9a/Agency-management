package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func getDataBaseUrl() string {
	host := os.Getenv("HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	if host == "" || user == "" || password == "" {
		host = "localhost"
		user = "postgres"
		password = "password"
	}
	return "host=" + host + " user=" + user + " password=" + password + " dbname=agency_management port=5432"
}

func Conn() *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(getDataBaseUrl()), &gorm.Config{},
	)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
