package controllers

import (
	"fmt"
	"log"
	"net/http"
	"whisky-review-platform/models"
	"whisky-review-platform/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Register new users
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

// LOGIN HANDLER
func LoginHandler(db *gorm.DB, k *gin.Context) {

	// Transfer JSON payload to LoginData struct
	var payload models.LoginData
	if err := k.ShouldBindJSON(&payload); err != nil {
		k.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Search for payload's email
	var user models.User
	if err := db.Where("email = ?", payload.Email).First(&user).Error; err != nil {
		k.JSON(http.StatusUnauthorized, gin.H{"email": "Incorrect Email or Password"})
		return
	}

	// Compare saved password with JSON payload
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		k.JSON(http.StatusUnauthorized, gin.H{"email": "Incorrect Email or Password"})
		return
	}

	// Generate JSON token for remainder of session
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		k.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Generate Session Token"})
		return
	}

	// Return token as JSON
	k.JSON(http.StatusOK, gin.H{"token": token})

}

// LOGOUT HANDLER
