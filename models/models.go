package models

// Whisky struct
type Whisky struct {
	ID         uint
	Name       string
	Region     string
	Age        int
	Distillery string
	PeatPPM    uint
	Chfilt     bool
	Notes      []TastingNotes
}

// Tasting Notes struct
type TastingNotes struct {
	ID       uint
	Note     string
	WhiskyID uint
}
