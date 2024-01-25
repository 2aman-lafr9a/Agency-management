package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func getDataBaseUrl() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if host == "" || user == "" || password == "" {
		host = "localhost"
		user = "postgres"
		password = "password"
		db_name = "agency_management"
		port = "5432"
	}
	return "host=" + host + " user=" + user + " password=" + password + " dbname=" + db_name + " port=" + port
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
