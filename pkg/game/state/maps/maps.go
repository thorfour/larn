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
	volcano   [][][]io.Runeable // Gross type alias?
	dungeon   [][][]io.Runeable
	home      [][]io.Runeable
	displaced io.Runeable
}

// New returns a set of maps to represent the game
func New(c *character.Character) *Maps {
	m := new(Maps)
	m.dungeon = dungeon()
	m.volcano = volcano()
	m.home = newLevel(homeLevel + 1)
	m.SpawnCharacter(c)
	return m
}

// TODO
func dungeon() [][][]io.Runeable {
	return nil
}

// TODO
func volcano() [][][]io.Runeable {
	return nil
}

// TODO
func (m *Maps) CurrentMap() [][]io.Runeable {
	return m.home
}

func (m *Maps) SpawnCharacter(c *character.Character) {

	x := c.Location().X
	y := c.Location().Y

	// Save what the character is standing on top of
	m.displaced = m.home[y][x]

	// Set the character to the location
	m.home[y][x] = c
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
	if !validMove(d, c) {
		return nil
	}

	old := c.Location()
	new := c.MoveCharacter(d)

	// Reset the displaced
	m.home[old.Y][old.X] = m.displaced

	// Save the newly displaced item
	m.displaced = m.home[new.Y][new.X]

	// Set the character to the location
	m.home[new.Y][new.X] = c

	return []io.Cell{&cell{old.X, old.Y, m.displaced}, &cell{new.X, new.Y, c}}
}

// validMove returns true if the move is allowed (i.e not off the edge, not into a wall
func validMove(d character.Direction, c *character.Character) bool {

	// Make the move and check its validity
	current := c.Location()
	current.Move(d)

	glog.V(6).Infof("ValidMove: (%v,%v)", current.X, current.Y)

	return current.X >= 0 && current.X < width && current.Y >= 0 && current.Y < height
}
