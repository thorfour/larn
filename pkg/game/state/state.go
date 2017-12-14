package state

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/golang/glog"
	"github.com/thorfour/larn/pkg/game/state/character"
	"github.com/thorfour/larn/pkg/game/state/items"
	"github.com/thorfour/larn/pkg/game/state/maps"
	"github.com/thorfour/larn/pkg/io"
)

const (
	logLength = 5 // Ideally should be the same as the game.logLength but is useful to be definde separately for debug
)

var (
	NoItemErr           = fmt.Errorf("You don't have item")
	AlreadyDisplacedErr = fmt.Errorf("There's something here already")
)

type logring []string

// add adds a new log to the log ring
func (log logring) add(s string) logring {
	log = append(log, s)      // Append the new string
	if len(log) > logLength { // remove first element if the log exceeds length
		log = log[1:]
	}
	glog.V(6).Infof("Add: %s", s)
	return log
}

// State holds all current game state
type State struct {
	StatLog logring
	C       *character.Character
	maps    *maps.Maps
	rng     *rand.Rand
}

func New() *State {
	glog.V(1).Info("Creating new state")
	s := new(State)
	s.C = new(character.Character)
	s.C.Init()
	s.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	s.maps = maps.New(s.C)

	// Display the welcome string at the bottom
	for i := 0; i < logLength-1; i++ {
		s.Log("")
	}
	s.Log("Welcome to larn -- Press ? for help")
	return s
}

// Drop drops an item where the player is standing, returns false if the player is already standing on an item
func (s *State) Drop(e rune) (items.Item, error) {
	if _, ok := s.maps.Displaced().(maps.Empty); !ok { // Check if player is already displacing an object
		return nil, AlreadyDisplacedErr
	}

	item, err := s.C.DropItem(e)
	if err != nil {
		return nil, err
	}
	s.maps.AddDisplaced(item)
	return item, nil
}

// Log adds the string to the statlog
func (s *State) Log(str string) {
	s.StatLog = s.StatLog.add(str)
}

// CurrentMap returns the current map the character is on
func (s *State) CurrentMap() [][]io.Runeable {
	return s.maps.CurrentMap()
}

// Move is for character movement
func (s *State) Move(d character.Direction) bool {

	// Move the character
	moved := s.maps.Move(d, s.C)

	if moved {
		// If the character is displacing something add it to the status log
		switch t := s.maps.Displaced().(type) {
		case *items.GoldPile:
			t.PickUp(s.C.Stats) // auto-pick up gold
			s.maps.RemoveDisplaced()
			s.Log(t.Log())
		case maps.Loggable:
			s.Log(t.Log())
		}
	}
	return moved
}

// Enter is used for entering into a building or dungeon/volcano
func (s *State) Enter() {

	glog.V(2).Infof("Enter request")

	// Check if character is standing on an enterable object
	switch t := s.maps.Displaced().(type) {
	case maps.Enterable:
		s.maps.EnterLevel(s.C, t.Enter())
	}
}

// PickUp will pick up the item the player is standing on
func (s *State) PickUp() {

	glog.V(2).Info("PickUp request")

	i, ok := s.maps.Displaced().(items.Item)
	if ok {
		i.PickUp(s.C.Stats)
		s.C.AddItem(i)
		s.maps.RemoveDisplaced()
	}
}

// Inventory request
func (s *State) Inventory() []string {
	glog.V(2).Info("Inventory request")
	return s.C.Inventory()
}
