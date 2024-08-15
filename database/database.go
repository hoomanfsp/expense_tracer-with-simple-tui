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
	dsn := dsn_genr()
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Expense{})
	return db, nil
}

func dsn_genr() string {
	err := godotenv.Load("./database/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Access the environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")
	return fmt.Sprintf("%s:%s@%s:%s/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, port, dbName)
}
