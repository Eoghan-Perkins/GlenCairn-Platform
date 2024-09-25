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
