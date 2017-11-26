package state

import (
	"math/rand"
	"time"

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

func (s *State) CurrentMap() [][]io.Runeable {
	return s.maps.CurrentMap()
}

func (s *State) Move(d character.Direction) []io.Cell {
	return s.maps.Move(d, s.C)
}
