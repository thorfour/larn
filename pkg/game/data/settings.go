package data

// Settings contains all game settings that are needed for opening/saving of games
type Settings struct {
	// SaveFile filepath location of the save file
	SaveFile string
	// UserID unique id of the user
	UserID uint64
	// Difficulty current game difficulty
	Difficulty uint
	// FromSaveFile if the current game was loaded from a save file
	FromSaveFile bool
}
