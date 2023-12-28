package db

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env", ".env.prod")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func Conn() *gorm.DB {
	db, err := gorm.Open(
		postgres.Open("host="+goDotEnvVariable("host")+" user="+goDotEnvVariable("username")+" password="+goDotEnvVariable("password")+" dbname=agency_management port=5432"), &gorm.Config{},
	)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
