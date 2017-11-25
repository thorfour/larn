package state

import (
	"math/rand"
	"time"

	"github.com/thorfour/larn/pkg/game/state/character"
)

// State holds all current game state
type State struct {
	character.Character
	rng *rand.Rand
}

func New() *State {
	s := new(State)
	s.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	return s
}
