package types

// Attackable indicates something is attackable
type Attackable interface {
	Damage(int) bool // take damage and return if dead
}

// Coordinate represents a cartesian map coordinate
type Coordinate struct {
	X int
	Y int
}
