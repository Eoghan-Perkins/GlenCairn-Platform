package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// Models
	"whisky-review-platform/models"
)

func main() {

	// Load Environment variables for datbase connection
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// .env variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Build dsn (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Connect to database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Successfully connected to database")

	// Datbase Auto-Migration
	db.AutoMigrate(&models.Whisky{}, &models.TastingNotes{})

	fmt.Println("Database migration completed successfully")

	// DATABASE TESTING --- TO BE DELETED
	whisky := models.Whisky{
		Name:       "Port Charlotte",
		Region:     "Islay",
		Age:        10,
		Distillery: "Bruichladdich",
		PeatPPM:    100,
		Notes: []models.TastingNotes{
			{Note: "Medicinal"},
			{Note: "Heavily Peated"},
			{Note: "Spicy"},
		},
	}

	result := db.Create(&whisky)
	if result.Error != nil {
		log.Fatalf("Fatal error while trying to add whisky record: %v", result.Error)
	}

	fmt.Println("Whisky record created sucessfully")
}
