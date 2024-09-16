package models

// Whisky struct
type Whisky struct {
	ID             uint
	Name           string
	Region         string
	Age            uint
	Distillery     string
	PeatPPM        uint
	ChillFiltering bool
	AverageRating  float32
	Notes          []TastingNotes `gorm:"foreignKey:WhiskyID;constraint:OnDelete:CASCADE;"`
	// Cascade deletion for Tasting Notes struct
}

// Tasting Notes struct
type TastingNotes struct {
	ID       uint `gorm:"primaryKey"`
	Note     string
	WhiskyID uint
}
