package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// Models
	"whisky-review-platform/controllers"
	"whisky-review-platform/models"
	// Controllers
	//"whisky-review-platform/controllers" -- Uncomment after integration
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

	// Database Auto-Migration
	db.AutoMigrate(&models.Whisky{}, &models.TastingNotes{})

	// CRUD functions testing via CLI --- DELETE LATER

	// Check for CLI arguments
	if len(os.Args) < 2 {
		log.Fatal("Please Provide a Command")
	}

	command := os.Args[1]

	switch command {
	case "create":
		controllers.AddWhisky(db, "Ardbeg 10", "Islay", 10, "Ardbeg Distillery", 76, false, 8.5, []string{
			"Smoky and rich",
			"Long peaty finish",
		})
	case "read":
		// Test ReadWhisky function (pass the whisky ID as the second argument)
		if len(os.Args) < 3 {
			log.Fatal("Please provide the whisky ID to read")
		}
		id, _ := strconv.ParseUint(os.Args[2], 10, 32)
		controllers.ReadWhisky(db, id)
	case "update":
		// Test AddTastingNote function (pass the whisky ID and new tasting note as arguments)
		if len(os.Args) < 4 {
			log.Fatal("Please provide the whisky ID and new tasting note")
		}
		id, _ := strconv.ParseUint(os.Args[2], 10, 32)
		newNote := os.Args[3]
		controllers.AddTastingNote(db, uint(id), newNote)
	case "delete":
		// Test DeleteWhisky function (pass the whisky ID as the second argument)
		if len(os.Args) < 3 {
			log.Fatal("Please provide the whisky ID to delete")
		}
		id, _ := strconv.ParseUint(os.Args[2], 10, 32)
		controllers.DeleteWhisky(db, uint(id))
	default:
		log.Fatalf("Unknown command: %s", command)
	}

}
