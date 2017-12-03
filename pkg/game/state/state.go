package state

import (
	"math/rand"
	"time"

	"github.com/golang/glog"
	"github.com/thorfour/larn/pkg/game/state/character"
	"github.com/thorfour/larn/pkg/game/state/maps"
	"github.com/thorfour/larn/pkg/io"
)

const (
	logLength = 5 // Ideally should be the same as the game.logLength but is useful to be definde separately for debug
)

type logring []string

// Add adds a new log to the log ring
func (log logring) Add(s string) logring {
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
		s.StatLog = s.StatLog.Add("")
	}
	s.StatLog = s.StatLog.Add("Welcome to larn -- Press ? for help")
	return s
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
		case maps.Loggable:
			s.StatLog = s.StatLog.Add(t.Log())
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
