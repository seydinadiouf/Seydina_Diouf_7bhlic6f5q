package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"school-manager/model"
	"strconv"
)

var database *gorm.DB
var e error

func DatabaseInit() {
	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, port)
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	errorMigration := database.AutoMigrate(&model.SchoolClass{})
	if errorMigration != nil {
		fmt.Println("Error migrating database :", errorMigration)
	}

	if e != nil {
		panic(e)
	}
}

func DB() *gorm.DB {
	return database
}
