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
	entrance  []Coordinate      // list of all entrances in each maze (i.e where a ladder from the previous maze drops you)
	active    [][]io.Runeable   // current active maze
	displaced io.Runeable       // object the player is currently standing on TODO should be moved to the character type
	current   int               // index of the active maze. active = mazes[current]
}

// EnterLevel moves a character from one level to the next by way of entrance or stairs
func (m *Maps) EnterLevel(c *character.Character, lvl int) {
	m.RemoveCharacter(c)
	m.SetCurrent(lvl)
	m.SpawnCharacter(m.entrance[lvl], c)
}

// New returns a set of maps to represent the game
func New(c *character.Character) *Maps {

	glog.V(2).Infof("Generating new maps")

	m := new(Maps)
	for i := uint(0); i < maxVolcano; i++ {

		m.mazes = append(m.mazes, newMap(i))

		if i == 1 { // dungeon 0 has an entrance
			m.mazes[i][height-1][width/2] = (Empty{})
			m.entrance = append(m.entrance, Coordinate{width / 2, height - 2})
		} else {
			// Set the entrace for the maze to a random location
			m.entrance = append(m.entrance, walkToEmpty(randMapCoord(), m.mazes[i]))
		}
	}
	m.active = m.mazes[homeLevel]
	m.SpawnCharacter(m.entrance[homeLevel], c)
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
func (m *Maps) SpawnCharacter(coord Coordinate, c *character.Character) {

	// Place the character on the map
	l, d := placeObject(coord, c, m.active)

	// Save the displaced element
	m.displaced = d

	// Set the character to the location
	c.Teleport(int(l.X), int(l.Y))

	m.setVisible(c)
}

type cell struct {
	x int
	y int
	io.Runeable
}

func (c *cell) X() int { return c.x }
func (c *cell) Y() int { return c.y }

func (m *Maps) Move(d character.Direction, c *character.Character) bool {

	// Validate the move
	if !m.validMove(d, c) {
		return false
	}

	old := c.Location()
	new := c.MoveCharacter(d)

	// Reset the displaced
	m.active[old.Y][old.X] = m.displaced

	// Check if they character moved to the exit of the dungeon
	if m.current == 1 && new.Y == height-1 {
		// Move the character to the home level
		m.RemoveCharacter(c)
		m.SetCurrent(homeLevel)
		m.SpawnCharacter(m.entrance[homeLevel], c)
		return true
	}

	// Save the newly displaced item
	m.displaced = m.active[new.Y][new.X]

	// Set the character to the location
	m.active[new.Y][new.X] = c

	m.setVisible(c)

	return true
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

// setVisible changes the visibilty of surrounding objects
func (m *Maps) setVisible(c *character.Character) {

	coord := c.Location()
	adj := append(adjacent(Coordinate{uint(coord.X), uint(coord.Y)}, false), diagonal(Coordinate{uint(coord.X), uint(coord.Y)}, false)...)
	for _, l := range adj {
		switch m.active[l.Y][l.X].(type) {
		case Visible:
			m.active[l.Y][l.X].(Visible).Visible(true)
		}
	}
}

// RemoveDisplaced removes the displaced object on the map (i.e the player picked up an item)
func (m *Maps) RemoveDisplaced() {
	m.displaced = Empty{m.current == homeLevel} // Set displaced to empty, so it gets replaced when the player moves
}

// AddDisplaced adds a displaced item to the map. (i.e the player dropped an item)
func (m *Maps) AddDisplaced(i io.Runeable) {
	m.displaced = i
}
