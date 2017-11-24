package game

const (
	DEFAULT_DIFFICULTY = -1
)

// SetDifficulty sets the games difficulty
func (d *Data) SetDifficulty(difficulty uint) {

	// Get the previous difficulty if this person has won a game (i.e is on the winner board)
	//pd := PrevDifficulty(d.userID)
}
