package controllers

import (
	"fmt"
	"log"
	"whisky-review-platform/models"

	"gorm.io/gorm"
)

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

func GetWhisky(db *gorm.DB, id uint) {

	// Get Scotch, handle error
	var whisky models.Whisky
	result := db.First(&whisky, id)
	if result.Error != nil {
		log.Fatalf("Could not locate scotch: ", result.Error)
	}

	// Read Granddaddy reeview
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
		whisky.AverageRating = 5.0
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

	return nil
}
