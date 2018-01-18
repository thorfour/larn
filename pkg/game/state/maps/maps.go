package maps

import (
	"github.com/golang/glog"
	"github.com/thorfour/larn/pkg/game/state/character"
	"github.com/thorfour/larn/pkg/game/state/monster"
	"github.com/thorfour/larn/pkg/io"
)

// Standard height and width of maps
const (
	height = 17
	width  = 67
)

// Maps is the collection of all the levels in the game
type Maps struct {
	monsters  [][]monster.Monster // list of all monsters on all levels
	mazes     [][][]io.Runeable   // slice of all mazes in the game
	entrance  []Coordinate        // list of all entrances in each maze (i.e where a ladder from the previous maze drops you)
	active    [][]io.Runeable     // current active maze
	displaced io.Runeable         // object the player is currently standing on TODO should be moved to the character type
	current   int                 // index of the active maze. active = mazes[current]
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
	m.monsters = make([][]monster.Monster, maxVolcano)
	for i := uint(0); i < maxVolcano; i++ {

		nm := newMap(i) // create the new map with items

		switch i {
		case homeLevel:
		case 1: // dungeon 0 has an entrance
			nm[height-1][width/2] = (Empty{})
			m.entrance = append(m.entrance, Coordinate{width / 2, height - 2})
			m.monsters[i] = spawnMonsters(nm, i, true) // spawn monsters onto the map
		default:
			// Set the entrace for the maze to a random location
			m.entrance = append(m.entrance, walkToEmpty(randMapCoord(), nm))
			m.monsters[i] = spawnMonsters(nm, i, true) // spawn monsters onto the map
		}

		m.mazes = append(m.mazes, nm)
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
	adj := append(adjacent(Coordinate{coord.X, coord.Y}, false), diagonal(Coordinate{coord.X, coord.Y}, false)...)
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

// function to vaporize walls at adjacent locations
func (m *Maps) VaporizeAdjacent(c *character.Character) {
	coord := c.Location()
	adj := append(adjacent(Coordinate{coord.X, coord.Y}, true), diagonal(Coordinate{coord.X, coord.Y}, true)...)
	for _, l := range adj { // set all to empty
		switch m.active[l.Y][l.X].(type) {
		case *Wall: // only vaporize walls
			m.active[l.Y][l.X] = Empty{}
		}
	}
}

// Adjacent returns all adjacent spaces to the character
func (m *Maps) Adjacent(c *character.Character) []io.Runeable {

	loc := Coordinate{c.Location().X, c.Location().Y}

	// get adjacent locations to the player
	coords := append(adjacent(loc, false), diagonal(loc, false)...)

	var adj []io.Runeable
	lvl := m.CurrentMap()

	// Populate the adjacent spaces
	for _, l := range coords {
		adj = append(adj, lvl[l.Y][l.X])
	}

	return adj
}

// LevelMonsters returns the list of monsters on the current level
func (m *Maps) LevelMonsters() []monster.Monster { return m.monsters[m.current] }

// ValidCoordindate returns true if the coordinate provided is within the map boundaries
func (m *Maps) ValidCoordinate(c Coordinate) bool {
	return c.X < width && c.X >= 0 && c.Y < height && c.Y >= 0
}
