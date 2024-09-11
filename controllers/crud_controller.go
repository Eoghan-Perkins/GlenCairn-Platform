package controllers

import (
	"fmt"
	"log"
	"strconv"

	"whisky-review-platform/models"

	"gorm.io/gorm"
)

// CRUD FUNCTIONS FOR MAIN.GO
// Create, Read, Update, and Delete whisky records

// Add new whisky to database
func AddWhisky(db *gorm.DB, name string, region string, age uint, dist string, ppm uint, ch bool, rting float32, tn []string) {
	// Organize any included tasting notes into TastingNotes struct for later addition
	notes := []models.TastingNotes{}
	for _, note := range tn {
		notes = append(notes, models.TastingNotes{
			Note: note,
		})
	}
	// Complete whisky record
	whisky := models.Whisky{
		Name:           name,
		Region:         region,
		Age:            age,
		Distillery:     dist,
		PeatPPM:        ppm,
		ChillFiltering: ch,
		Notes:          notes,
		AverageRating:  rting,
	}
	// Error Handling
	result := db.Create(&whisky)
	if result.Error != nil {
		log.Fatalf("Fatal error while trying to add whisky record: %v", result.Error)
	}
	// Success
	fmt.Println("Whisky record created sucessfully")
}

// Retrieve Whisky data
func ReadWhisky(db *gorm.DB, id uint) {
	// Declare empty whisky struct to hold data after retrieval from db
	var whisky models.Whisky
	// Preload any tasting notes
	result := db.Preload("Notes").First(&whisky, id)
	// Kill function if data retrieval fails
	if result.Error != nil {
		log.Fatalf("Failed to retrieve scotch from database: %v", result.Error)
	}
	// Print out data regarding whisky
	var name = whisky.Name
	var dist = whisky.Distillery
	var ppm = whisky.PeatPPM
	var region = whisky.Region
	// Handle NAS whiskies (No-Age-Statement)
	if whisky.Age == 00 {
		fmt.Println(name+" is from "+dist+" in the "+region+". It is peated at %d PPM and has no age statement",
			name, dist, region, ppm)
	} else {
		var age = strconv.Itoa(int(whisky.Age))
		fmt.Println("%s is from %s in the %s region. It is peated at %d PPM and aged for %s years",
			name, dist, region, ppm, age)
	}
	// Print Average Rating
	fmt.Println("Average Rating: %f", whisky.AverageRating)
	// Print Tasting Notes
	fmt.Println("Tasting Notes:")
	for _, note := range whisky.Notes {
		fmt.Println("- %s", note.Note)
	}
}

// Update a whisky's data
func AddTastingNote(db *gorm.DB, id uint, note string) {
	// Empty whisky struct for loading
	var whisky models.Whisky
	// First whisky in db that has id-match
	result := db.First(&whisky, id)
	// Check for accessing error
	if result.Error != nil {
		log.Fatalf("Failed to retrive whisky from database: %v", result.Error)
	}
	// New tasting note
	newnote := models.TastingNotes{
		Note: note,
	}
	// Append to whisky's TastingNotes struct
	whisky.Notes = append(whisky.Notes, newnote)
	// Save note to whisky record
	result = db.Save(&whisky)
	// Error handling for note saving
	if result.Error != nil {
		log.Fatalf("could not save tasting note. Error: %v", result.Error)
	}
	// Success
	fmt.Println("Tasting note saved successfully")
}

// Delete a whisky from the database
func DeleteWhisky(db *gorm.DB, id uint) {
	// Empty whisky struct for loading
	var whisky models.Whisky
	// First whisky in db that has id match
	result := db.First(&whisky, id)
	// Access error handling
	if result.Error != nil {
		log.Fatalf("Failed to retrieve whisky from database: %v", result.Error)
	}
	// Delete error handling
	result = db.Delete(&whisky)
	if result.Error != nil {
		log.Fatalf("Failed to delete whisky from database. Error: %v", result.Error)
	}
	// Success
	fmt.Println("Whisky deleted successfully")
}
