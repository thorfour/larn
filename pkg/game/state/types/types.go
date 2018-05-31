package types

// Attackable indicates something is attackable
type Attackable interface {
	Damage(int) bool // take damage and return if dead
}

// Coordinate represents a map coordinate
type Coordinate interface {
	X() int
	Y() int
}
