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
	Notes          []TastingNotes
}

// Tasting Notes struct
type TastingNotes struct {
	ID       uint
	Note     string
	WhiskyID uint
}
