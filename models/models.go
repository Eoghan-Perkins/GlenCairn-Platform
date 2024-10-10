package models

// Whisky struct
type Whisky struct {
	ID             uint
	Name           string
	Region         string
	Age            uint // Years
	Distillery     string
	PeatPPM        uint
	ChillFiltering bool
	AverageRating  float32 // Rating out of 5 stars
	ReviewCount    int16
	UserFavorite   float32        // Percentage
	Notes          []TastingNotes `gorm:"foreignKey:WhiskyID;constraint:OnDelete:CASCADE;"` // Cascade deletion for Tasting Notes struct
	Reviews        []UserReview   `gorm:"foreignKey:WhiskyID;constraint:OnDelete:CASCADE;"`
}

// Tasting Notes struct
type TastingNotes struct {
	ID       uint `gorm:"primaryKey"`
	Note     string
	WhiskyID uint
}

// User Struct
type User struct {
	ID         uint
	Username   string
	Password   string
	NumReviews uint // Number of scotch reviews
	UserAge    uint //days
	Reviews    []UserReview
}

// Individual User Reviews Struct
type UserReview struct {
	ID       uint `gorm:"primaryKey"`
	WhiskyID uint
	Whisky   Whisky `gorm:"foreignKey:WhiskyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relation to Whisky struct

	AuthorID uint
	Author   User `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Rating   float32
	Favorite bool
	Notes    []TastingNotes
}

// Struct for User Registration Data
type RegistrationData struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required, email"`
	Password string `json:"password" binding:"required, min=8"`
}

type LoginData struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
