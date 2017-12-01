package maps

import (
	"github.com/golang/glog"
	"github.com/thorfour/larn/pkg/game/state/character"
	"github.com/thorfour/larn/pkg/io"
)

// Standard height and width of maps
const (
	height = 17
	width  = 67
)

// Maps is the collection of all the levels in the game
type Maps struct {
	mazes     [][][]io.Runeable // slice of all mazes in the game
	active    [][]io.Runeable   // current active maze
	displaced io.Runeable       // object the player is currently standing on TODO should be moved to the character type
	current   int               // index of the active maze. active = mazes[current]
}

// New returns a set of maps to represent the game
func New(c *character.Character) *Maps {
	m := new(Maps)
	for i := uint(0); i < maxVolcano; i++ {
		m.mazes = append(m.mazes, newMap(i))
	}
	m.active = m.mazes[homeLevel]
	m.SpawnCharacter(c)
	return m
}

// CurrentMap returns the current map where the character is located
func (m *Maps) CurrentMap() [][]io.Runeable {
	return m.active
}

// RemoveCharacter is used for when a character leaves a map
func (m *Maps) RemoveCharacter(c *character.Character) {
	l := c.Location()
	m.active[l.Y][l.X] = m.displaced
	m.displaced = nil
}

// SpawnCharacter places the character on the home level
func (m *Maps) SpawnCharacter(c *character.Character) {

	// Place the character on the map
	l, d := placeObject(randMapCoord(), c, m.active)

	// Save the displaced element
	m.displaced = d

	// Set the character to the location
	c.Teleport(int(l.X), int(l.Y))
}

type cell struct {
	x int
	y int
	io.Runeable
}

func (c *cell) X() int { return c.x }
func (c *cell) Y() int { return c.y }

func (m *Maps) Move(d character.Direction, c *character.Character) []io.Cell {

	// Validate the move
	if !m.validMove(d, c) {
		return nil
	}

	old := c.Location()
	new := c.MoveCharacter(d)

	// Reset the displaced
	m.active[old.Y][old.X] = m.displaced

	// Save the newly displaced item
	m.displaced = m.active[new.Y][new.X]

	// Set the character to the location
	m.active[new.Y][new.X] = c

	return []io.Cell{&cell{old.X, old.Y, m.active[old.Y][old.X]}, &cell{new.X, new.Y, c}}
}

// validMove returns true if the move is allowed (i.e not off the edge, not into a wall
func (m *Maps) validMove(d character.Direction, c *character.Character) bool {

	// Make the move and check its validity
	current := c.Location()
	current.Move(d)

	glog.V(6).Infof("ValidMove: (%v,%v)", current.X, current.Y)

	// Ensure the character isn't going off the grid, tron
	inBounds := current.X >= 0 && current.X < width && current.Y >= 0 && current.Y < height

	// Ensure the character is going onto an empty location
	isDisplaceable := false
	if inBounds {
		switch m.active[current.Y][current.X].(type) {
		case Displaceable:
			isDisplaceable = true
		}
	}

	return inBounds && isDisplaceable
}

// Displaced returns the displaced object
func (m *Maps) Displaced() io.Runeable {
	return m.displaced
}

// SetCurrent sets the current map level to display (i.e the character moved between levels)
func (m *Maps) SetCurrent(lvl int) {
	glog.V(2).Infof("Setting current level %v", lvl)
	m.current = lvl
	m.active = m.mazes[m.current]
}
