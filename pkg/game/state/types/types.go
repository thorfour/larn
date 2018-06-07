package types

// Direction cardinal direction
type Direction uint8

const (
	// Up direction
	Up Direction = iota
	// Down direction
	Down
	// Left direction
	Left
	// Right direction
	Right
	// UpLeft direction
	UpLeft
	// UpRight direction
	UpRight
	// DownLeft direction
	DownLeft
	// DownRight direction
	DownRight
	// Here direction
	Here
)

// Attackable indicates something is attackable
type Attackable interface {
	Damage(int) bool // take damage and return if dead
}

// Coordinate represents a cartesian map coordinate
type Coordinate struct {
	X int
	Y int
}

// Move coordinate in a given direction, and return the new coordinate
func Move(c Coordinate, d Direction) Coordinate {
	switch d {
	case Up:
		c.Y--
	case Down:
		c.Y++
	case Left:
		c.X--
	case Right:
		c.X++
	case UpLeft:
		c.Y--
		c.X--
	case UpRight:
		c.Y--
		c.X++
	case DownLeft:
		c.Y++
		c.X--
	case DownRight:
		c.Y++
		c.X++
	case Here:
	}

	return c
}
