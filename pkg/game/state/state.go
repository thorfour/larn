package state

import (
	"math/rand"
	"time"

	"github.com/golang/glog"
	"github.com/thorfour/larn/pkg/game/state/character"
	"github.com/thorfour/larn/pkg/game/state/maps"
	"github.com/thorfour/larn/pkg/io"
)

// State holds all current game state
type State struct {
	C    *character.Character
	maps *maps.Maps
	rng  *rand.Rand
}

func New() *State {
	s := new(State)
	s.C = new(character.Character)
	s.C.Init()
	s.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	s.maps = maps.New(s.C)
	return s
}

// CurrentMap returns the current map the character is on
func (s *State) CurrentMap() [][]io.Runeable {
	return s.maps.CurrentMap()
}

// Move is for character movement
func (s *State) Move(d character.Direction) []io.Cell {
	return s.maps.Move(d, s.C)
}

// Enter is used for entering into a building or dungeon/volcano
func (s *State) Enter() {

	glog.V(2).Infof("Enter request")

	// Check if character is standing on an enterable object
	switch t := s.maps.Displaced().(type) {
	case maps.Enterable:
		s.maps.RemoveCharacter(s.C)
		s.maps.SetCurrent(t.Enter())
		s.maps.SpawnCharacter(s.C) // TODO THOR map entrances should be used
	}
}
