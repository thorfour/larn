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
	character.Character
	maps *maps.Maps
	rng  *rand.Rand
}

func New() *State {
	s := new(State)
	s.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	s.maps = maps.New()
	return s
}

func (s *State) CurrentMap() [][]io.Runeable {
	return s.maps.CurrentMap()
}
