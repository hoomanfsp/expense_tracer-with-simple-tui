package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

type Expense struct {
	ID          uint `gorm:"primary_key"`
	Amount      float64
	Description string
}

func InitDB() (*gorm.DB, error) {
	// Generate DSN
	dsn := dsnGen()

	// Open the database connection
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// AutoMigrate the Expense struct into database table
	db.AutoMigrate(&Expense{})

	return db, nil
}

func dsnGen() string {
	// Load .env file
	err := godotenv.Load("./database/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Access environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT") // Typically, the environment variable for port would be DB_PORT, not just PORT

	// Generate the DSN string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	return dsn
}
