package controllers

import (
	"fmt"
	"log"
	"whisky-review-platform/models"

	"gorm.io/gorm"
)

// CRUD functions for MAIN.GO
// User Controller Functions

// Rewrriten in auth_controllers
/*
func NewUser(db *gorm.DB, username string, password string) {

	user := models.User{
		Username:   username,
		Password:   password,
		UserAge:    1,
		NumReviews: 0,
	}

	result := db.Create(&user)
	if result.Error != nil {
		log.Fatalf("Failed to upload new user to database: ", result.Error)
	}
	// Success
	fmt.Println("New User Record Created Successfully")

}
*/

func ReadUser(db *gorm.DB, id uint) {

	// Empty user for loading
	var user models.User
	// Get user
	result := db.First(&user, id)
	// Error Handling
	if result.Error != nil {
		log.Fatal("Could Not Read User From Database: ", result.Error)
	}

	// Print out info
	fmt.Println("Username: ", user.Username)
	//fmt.Println("User Age: ", user.UserAge, "days") Uncomment after adding day-counting functionality
	fmt.Println("Number of Reviews: ", user.NumReviews)

}
