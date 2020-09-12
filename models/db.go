package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB //database

func init() {
	envLoadError := godotenv.Load()
	if envLoadError != nil {
		log.Fatal("Error loading .env file")
	}

	dbURI := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSL_MODE"), os.Getenv("DB_TIMEZONE"))
	fmt.Println(dbURI)
	connection, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting database")
		panic(err.Error())
	} else {
		fmt.Println("Connected to database")
		db = connection
		connection.Debug().AutoMigrate(
			User{},
			Todo{},
			UserDetails{},
		)
	}
}

func GetDBInstance() *gorm.DB {
	return db
}

// gorm.Model definition
type Model struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
