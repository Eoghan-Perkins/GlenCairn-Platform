package controllers

import (
	"fmt"
	"log"
	"whisky-review-platform/models"

	"gorm.io/gorm"
)

// CRUD FUNCTIONS FOR MAIN.GO
// Review Controller Functions

func ReadUserReview(db *gorm.DB, id uint) {

	// Get Review, handle errors
	var review models.UserReview
	result := db.First(&review, id)
	if result.Error != nil {
		log.Fatalf("Could not locate review: ", result.Error)
	}

	// Read review
	fmt.Println(review.Whisky.Name)
	fmt.Println(review.Whisky.Distillery)
	fmt.Println(review.Whisky.Region)
	fmt.Println("Author:", review.Author.Username)
	fmt.Println(review.Whisky.Age)
	fmt.Println(review.Favorite)
	fmt.Println(review.Notes)

}

// Returns data regarding the desired whisky's averages
func ReadAverageReview(db *gorm.DB, id uint) {

	// Get Scotch, handle error
	var whisky models.Whisky
	result := db.First(&whisky, id)
	if result.Error != nil {
		log.Fatalf("Could not locate scotch: ", result.Error)
	}

	// Read Grandaddy review
	fmt.Println(whisky.Name)
	fmt.Println("Average Rating (out of 5): ", whisky.AverageRating)
	fmt.Println("Percentage of User's Favorite: ", whisky.UserFavorite)
	fmt.Println("Age: ", whisky.Age)
	fmt.Println("Distillery: ", whisky.Distillery)
	fmt.Println("Region: ", whisky.Region)
	fmt.Println("Peat PPM: ", whisky.PeatPPM)
	fmt.Println("Chill Filtered: ", whisky.ChillFiltering)
	fmt.Println("Tasting Notes:")
	for _, note := range whisky.Notes {
		fmt.Println("-", note)
	}
}

// This function is passed a set of parameters that are used to create and add a new UserReview struct
func AddUserReview(db *gorm.DB, userID uint, whiskyID uint, favorite bool, rating float32, sa []string) {

	// Creates []TastingNotes struct from passed string array
	var notes []models.TastingNotes
	for _, note := range sa {
		notes = append(notes, models.TastingNotes{
			Note: note,
		})
	}

	// Build struct
	review := models.UserReview{
		WhiskyID: whiskyID,
		AuthorID: userID,
		Favorite: favorite,
		Rating:   rating,
		Notes:    notes,
	}

	// Add struct to db, handle any errors
	result := db.Create(&review)
	if result.Error != nil {
		log.Fatalf("Could Not Add Review. Error : ", result.Error)
	}

	// Update mother whisky's metrics based on input
	UpdateAverageRating(db, whiskyID)  // Update Average rating to whisky struct post-review
	UpdateReviewCount(db, whiskyID, 1) // Update Review count

}

// Function Calculates and saves a whisky struct's average review score
func UpdateAverageRating(db *gorm.DB, id uint) error {

	// Retrieve whisky, handle any errors
	var whisky models.Whisky
	if err := db.First(&whisky, id).Error; err != nil {
		return fmt.Errorf("Whisky Not Found. Error: ", err)
	}

	// Retrieve reviews, handle any errors
	var reviews []models.UserReview
	if err := db.Where("whisky_id = ?", id).Find(&reviews).Error; err != nil {
		return fmt.Errorf("Error Finding Reviews: ", err)
	}

	// Get length of reviews
	var reviews_len = float32(len(reviews))

	// Calculate and save average review score
	if reviews_len == 0 {
		whisky.AverageRating = 2.5
	} else {
		var total = float32(0)
		for _, review := range reviews {
			total += float32(review.Rating)
		}
		var final_rating = (total / reviews_len)
		whisky.AverageRating = final_rating
	}

	// Handle any errors during saving
	if err := db.Save(&whisky).Error; err != nil {
		fmt.Errorf("Could not save score. Error: ", err)
	}

	log.Printf("Review Score Successfully Updated!")
	return nil
}

// Function takes in a signed integer (1/-1) to increase or decrease
// a whisky struct's review count

func UpdateReviewCount(db *gorm.DB, id uint, pm int16) {

	// Get empyt whisky model
	var whisky models.Whisky
	// Retrieve whisky, handle any errors
	if err := db.First(&whisky, id).Error; err != nil {
		fmt.Println("Whisky Not Found. Error: ", err)
		log.Fatalf("Whisky Not Found. Error: ", err)
	}

	// Increase or decrease review count
	whisky.ReviewCount += pm

	// Handle any errors during saving
	if err := db.Save(&whisky).Error; err != nil {
		fmt.Println("Could not save score. Error: ", err)
		log.Fatalf("Could not save score. Error: ", err)
	}

	log.Printf("Review count for", whisky.Name, "successfully updated!")

}
