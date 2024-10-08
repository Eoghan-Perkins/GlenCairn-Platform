package controllers

import (
	"fmt"
	"log"
	"net/http"
	"whisky-review-platform/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegistrationHandler(db *gorm.DB, k *gin.Context) {

	// Attach incoming JSON user registration data to empty RegistrationData struct
	var userdata models.RegistrationData
	if err := k.ShouldBindJSON(&userdata); err != nil { // Check for bad JSON
		k.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Secure password string via hashing
	pwHash, err := bcrypt.GenerateFromPassword([]byte(userdata.Password), bcrypt.DefaultCost)
	if err != nil {
		k.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Hash Password"})
		return
	}

	// Fill out user struct
	user := models.User{
		Username:   userdata.Username,
		Password:   string(pwHash),
		UserAge:    1,
		NumReviews: 0,
	}

	// Add to db, check for errors
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal("Failed to upload new user to database: ", result.Error)
	}

	// Success
	fmt.Println("New User Record Created Successfully")

}
