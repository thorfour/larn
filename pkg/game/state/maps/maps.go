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
	m.home = newMap(homeLevel)
	m.SpawnCharacter(c)
	return m
}

// dungeon creates all the levels of the dungeon
func dungeon() [][][]io.Runeable {
	d := make([][][]io.Runeable, maxDungeon)
	for i := range d {
		d[i] = newMap(uint(i + 1)) // add 1 to avoid creating a home level
	}
	return d
}

// volcano creates all the levels of the volcano
func volcano() [][][]io.Runeable {
	v := make([][][]io.Runeable, maxVolcano)
	for i := range v {
		v[i] = newMap(uint(i + 1 + maxDungeon)) // +1 + maxDungeon to indicate it's a volcano
	}
	return v
}

// TODO
func (m *Maps) CurrentMap() [][]io.Runeable {
	return m.home
}

// SpawnCharacter places the character on the home level
func (m *Maps) SpawnCharacter(c *character.Character) {

	// Save the displaced element
	l := randMapCoord()
	m.displaced = m.home[l.Y][l.X]

	// Place the character on the map
	placeObject(l, c, m.home)

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
