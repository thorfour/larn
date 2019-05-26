package maps

import (
	"math"
	"math/rand"

	log "github.com/sirupsen/logrus"
	"github.com/thorfour/larn/pkg/game/state/character"
	"github.com/thorfour/larn/pkg/game/state/conditions"
	"github.com/thorfour/larn/pkg/game/state/monster"
	"github.com/thorfour/larn/pkg/game/state/types"
	"github.com/thorfour/larn/pkg/io"
)

// Standard height and width of maps
const (
	height = 17
	width  = 67
)

// Maps is the collection of all the levels in the game
type Maps struct {
	monsters [][]*monster.Monster // list of all monsters on all levels
	mazes    [][][]io.Runeable    // slice of all mazes in the game
	entrance []types.Coordinate   // list of all entrances in each maze (i.e where a ladder from the previous maze drops you)
	active   [][]io.Runeable      // current active maze
	current  int                  // index of the active maze. active = mazes[current]
}

// EnterLevel moves a character from one level to the next by way of entrance or stairs
func (m *Maps) EnterLevel(c *character.Character, lvl int) {
	m.RemoveCharacter(c)
	m.SetCurrent(lvl)
	m.SpawnCharacter(m.entrance[lvl], c)
}

// New returns a set of maps to represent the game
func New(c *character.Character) *Maps {

	log.Info("Generating new maps")

	m := new(Maps)
	m.monsters = make([][]*monster.Monster, MaxVolcano)
	for i := uint(0); i < MaxVolcano; i++ {

		nm := newMap(i) // create the new map with items

		switch i {
		case homeLevel:
			m.entrance = append(m.entrance, walkToEmpty(randMapCoord(), nm)) // TODO change this to be next to the dungeon entrance
		case 1: // dungeon 0 has an entrance
			nm[height-1][width/2] = (Empty{})
			m.entrance = append(m.entrance, types.Coordinate{X: width / 2, Y: height - 2})
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
	m.active[l.Y][l.X] = c.Displaced
	c.Displaced = nil
}

// SpawnCharacter places the character on the home level
func (m *Maps) SpawnCharacter(coord types.Coordinate, c *character.Character) {
	log.WithField("coord", coord).Info("Spawning Character")

	// Place the character on the map
	l, d := placeObject(coord, c, m.active)

	// Save the displaced element for the character
	c.Displaced = d

	// Set the character to the location
	c.Teleport(int(l.X), int(l.Y))

	m.SetVisible(c)
}

type cell struct {
	x int
	y int
	io.Runeable
}

func (c *cell) X() int { return c.x }
func (c *cell) Y() int { return c.y }

// Move a character on the map. First bool indiactes if the character moved. Second bool indicates if the character attacked.
// They will never both be set
func (m *Maps) Move(d types.Direction, c *character.Character) (bool, bool) {

	// Validate the move
	if !m.validMove(d, c) {
		return false, m.isAttack(d, c)
	}

	old := c.Location()
	new := c.MoveCharacter(d)

	// Reset the displaced
	m.active[old.Y][old.X] = c.Displaced

	// Check if they character moved to the exit of the dungeon
	if m.current == 1 && new.Y == height-1 {
		// Move the character to the home level
		m.RemoveCharacter(c)
		m.SetCurrent(homeLevel)
		m.SpawnCharacter(m.entrance[homeLevel], c)
		return true, false
	}

	// Save the newly displaced item
	c.Displaced = m.active[new.Y][new.X]

	// Set the character to the location
	m.active[new.Y][new.X] = c

	return true, false
}

// validMove returns true if the move is allowed (i.e not off the edge, not into a wall
func (m *Maps) validMove(d types.Direction, c *character.Character) bool {

	// Make the move and check its validity
	newLoc := types.Move(c.Location(), d)

	log.WithField("coord", types.Coordinate{X: newLoc.X, Y: newLoc.Y}).Info("valid move")

	// Ensure the character isn't going off the grid, tron
	inBounds := m.ValidCoordinate(newLoc)

	// If the player has walk through walls active they are allowed to displace anything
	if c.Cond.EffectActive(conditions.WalkThroughWalls) {
		return inBounds
	}

	// Ensure the character is going onto an empty location
	isDisplaceable := false
	if inBounds {
		switch m.active[newLoc.Y][newLoc.X].(type) {
		case Displaceable:
			isDisplaceable = true
		}
	}

	return inBounds && isDisplaceable
}

// isAttack returns true if the move would result in attacking a monster
func (m *Maps) isAttack(d types.Direction, c *character.Character) bool {

	// Make the move and check its validity
	newLoc := types.Move(c.Location(), d)

	// Check if the character would be attacking a monster
	if m.ValidCoordinate(newLoc) {
		switch m.active[newLoc.Y][newLoc.X].(type) {
		case types.Attackable:
			return true
		}
	}
	return false
}

// SetCurrent sets the current map level to display (i.e the character moved between levels)
func (m *Maps) SetCurrent(lvl int) {
	log.WithField("lvl", lvl).Info("setting current level")
	m.current = lvl
	m.active = m.mazes[m.current]
}

// CurrentLevel returns the current level the character is on
func (m *Maps) CurrentLevel() int { return m.current }

// SetVisible changes the visibilty of surrounding objects
func (m *Maps) SetVisible(c *character.Character) {

	coord := c.Location()
	adj := append(adjacent(coord, false), diagonal(coord, false)...)
	for _, l := range adj {
		switch m.active[l.Y][l.X].(type) {
		case Visible:
			m.active[l.Y][l.X].(Visible).Visible(true)
		}
	}
}

// RemoveAt removes the object at the given coordinate (i.e a monster died)
func (m *Maps) RemoveAt(c types.Coordinate) {
	m.active[c.Y][c.X] = Empty{}
}

// VaporizeAdjacent to vaporize walls at adjacent locations
func (m *Maps) VaporizeAdjacent(c *character.Character) {
	coord := c.Location()
	adj := append(adjacent(coord, true), diagonal(coord, true)...)
	for _, l := range adj { // set all to empty
		switch m.active[l.Y][l.X].(type) {
		case *Wall: // only vaporize walls
			m.active[l.Y][l.X] = Empty{}
		}
	}
}

// AdjacentCoords returns all adjacent coordinates to the location
func (m *Maps) AdjacentCoords(c types.Coordinate) []types.Coordinate {
	return append(adjacent(c, false), diagonal(c, false)...)
}

// Adjacent returns all adjacent spaces to the location
func (m *Maps) Adjacent(c types.Coordinate) []io.Runeable {

	// get adjacent locations to the player
	coords := m.AdjacentCoords(c)

	var adj []io.Runeable
	lvl := m.CurrentMap()

	// Populate the adjacent spaces
	for _, l := range coords {
		adj = append(adj, lvl[l.Y][l.X])
	}

	return adj
}

// LevelMonsters returns the list of monsters on the current level
func (m *Maps) LevelMonsters() []*monster.Monster { return m.monsters[m.current] }

// ValidCoordinate returns true if the coordinate provided is within the map boundaries
func (m *Maps) ValidCoordinate(c types.Coordinate) bool {
	return c.X < width && c.X >= 0 && c.Y < height && c.Y >= 0
}

// Distance returns the distance between coordinates
func (m *Maps) Distance(c0, c1 types.Coordinate) int {
	return int(math.Abs(float64(c0.X-c1.X)) + math.Abs(float64(c0.Y-c1.Y)))
}

// At returns whatever is at the given location on the active map
func (m *Maps) At(c types.Coordinate) io.Runeable {
	return m.active[c.Y][c.X]
}

// NewEmptyTile returns a new Empty map tile
func (m *Maps) NewEmptyTile() Empty {
	return Empty{m.current == homeLevel}
}

// TouchAllInteriorCoordinates walks the maps internal coordinats, and executes the given function on them
func (m *Maps) TouchAllInteriorCoordinates(f func(io.Runeable, int, int)) {
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			f(m.CurrentMap()[y][x], x, y)
		}
	}
}

// TouchAllInteriorCoordinatesOnAllMaps is the same as TouchAllInteriorCoordinates except it operates on all maps not just current
func (m *Maps) TouchAllInteriorCoordinatesOnAllMaps(f func(io.Runeable, int, int)) {
	// TODO THOR
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			f(m.CurrentMap()[y][x], x, y)
		}
	}
}

// Swap places the object at the given coordinate and returns the item that was previously there
func (m *Maps) Swap(c types.Coordinate, o io.Runeable) io.Runeable {
	displaced := m.CurrentMap()[c.Y][c.X]
	m.CurrentMap()[c.Y][c.X] = o

	return displaced
}

// OutOfBounds returns true of the given coordinate c is out of map boundaries
func (m *Maps) OutOfBounds(c types.Coordinate) bool {
	return c.X < 0 || c.X > width-1 || c.Y < 0 || c.Y > height-1
}

// OuterWall returns true if the corrdinate c locates an outer dungeon/volcano wall
func (m *Maps) OuterWall(c types.Coordinate) bool {
	if m.CurrentLevel() == homeLevel {
		return false
	}

	return c.X == 0 || c.X == width-1 || c.Y == 0 || c.Y == height-1
}

// RandomDisplaceableCoordinate returns a coordinate with a displaceable object anywhere in the current maze
func (m *Maps) RandomDisplaceableCoordinate() types.Coordinate {

	// randomly select a coordinate in the maze
	c := types.Coordinate{X: rand.Intn(width), Y: rand.Intn(height)}
	for { // continue selecting new coordinates until a displaceable coordinate is found
		if _, ok := m.At(c).(Displaceable); ok {
			return c
		}

		c = types.Coordinate{X: rand.Intn(width), Y: rand.Intn(height)}
	}
}
