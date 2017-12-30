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
	logLength = 5     // Ideally should be the same as the game.logLength but is useful to be definde separately for debug
	timeLimit = 30000 // max time to win a game
)

var (
	NoItemErr           = fmt.Errorf("You don't have item")
	AlreadyDisplacedErr = fmt.Errorf("There's something here already")
	DidntWork           = fmt.Errorf("  It didn't seem to work")
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
	StatLog  logring
	C        *character.Character
	Active   map[string]func()
	maps     *maps.Maps
	rng      *rand.Rand
	timeUsed uint
}

func New() *State {
	glog.V(1).Info("Creating new state")
	s := new(State)
	s.C = new(character.Character)
	s.C.Init()
	s.Active = make(map[string]func())
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
	s.timeUsed++
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
	s.timeUsed++

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
	s.timeUsed++

	// Check if character is standing on an enterable object
	switch t := s.maps.Displaced().(type) {
	case maps.Enterable:
		s.maps.EnterLevel(s.C, t.Enter())
	}
}

// PickUp will pick up the item the player is standing on
func (s *State) PickUp() {
	glog.V(2).Info("PickUp request")
	s.timeUsed++

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

// TimeStr returns the current time elapsed in the game
func (s *State) TimeStr() string {
	return fmt.Sprintf("Elapsed time is %v. You have %v mobuls left", (s.timeUsed+99)/100+1, (timeLimit-s.timeUsed)/100)
}

// Read is for the player to read a scroll or book
func (s *State) Read(e rune) error {
	glog.V(2).Info("Read requested")
	s.timeUsed++

	l, err := s.C.Read(e)
	if err != nil {
		return err
	}

	// Log all the information that read returned
	for _, r := range l {
		s.Log(r)
	}

	return nil
}

// Cast casts the requested spell
func (s *State) Cast(spell string) error {
	s.timeUsed++
	var sp *items.Spell
	if !DEBUG {
		var err error
		sp, err = s.C.Cast(spell)
		if err != nil {
			return err
		}

		if s.Active["stp"] != nil { // can't cast spells when time is stopped
			return DidntWork
		}
	}

	if DEBUG { // Pass through the spell for debugging
		sp = &items.Spell{Code: spell}
	}

	switch sp.Code {
	case "vpr": // vaporize rock
		s.maps.VaporizeAdjacent(s.C)
	case "cbl": // cure blindness
		s.Active[sp.Code] = nil
	case "hel": // healing
		s.C.Heal(20 + int(s.C.Stats.Level<<1))
	case "sca": // scare monsters
		fallthrough
	case "hld": // hold monsters
		s.decay(sp.Code, rand.Intn(10)+int(s.C.Stats.Level), func() {})
	case "stp": // time stop
		s.decay(sp.Code, rand.Intn(20)+(int(s.C.Stats.Level)<<1), func() {})
	case "glo":
		if s.Active[sp.Code] == nil {
			s.C.Stats.Ac += 10
		}
		if s.C.Stats.Intelligence > 3 { // globe decreases intelligence to minimum of 3
			s.C.Stats.Intelligence--
		}
		s.decay(sp.Code, 200, func() { s.C.Stats.Ac -= 10 })
	case "str":
		if s.Active[sp.Code] == nil {
			s.C.Stats.Str += 3
		}
		s.decay(sp.Code, 150+rand.Intn(100), func() { s.C.Stats.Str -= 3 })
	case "dex":
		if s.Active[sp.Code] == nil {
			s.C.Stats.Dex += 3
		}
		s.decay(sp.Code, 400, func() { s.C.Stats.Dex -= 3 })
	case "pro":
		if s.Active[sp.Code] == nil {
			s.C.Stats.Ac += 2 // protection field +2
		}
		s.decay(sp.Code, 250, func() { s.C.Stats.Ac -= 2 })
	case "cld":
		fallthrough
	case "ssp":
		fallthrough
	case "bal":
		fallthrough
	case "lit":
		fallthrough
	case "mle":
		panic("TODO")
	}

	return nil
}

// decay adds a decay function to the Active functions map
func (s *State) decay(code string, dur int, f func()) {
	s.Active[code] = func() {
		dur--
		if dur == 0 {
			f() // execute the func
			// remove it from the list of actives
			s.Active[code] = nil
		}
	}
}

// IdentTrap notifies the player if there are traps adjacent
func (s *State) IdentTrap() {

	// Get adjacent spaces
	adj := s.maps.Adjacent(s.C)

	// Check all loc for traps
	var found bool
	for _, l := range adj {
		if t, ok := l.(*items.Trap); ok {
			switch t.TrapType {
			case items.TeleTrap:
				s.Log("It's a teleport trap")
			case items.ArrowTrap:
				s.Log("It's an arrow trap")
			case items.DartTrap:
				s.Log("It's an dart trap")
			case items.DoorTrap:
				s.Log("It's a trapdoor")
			}
			found = true
		}
	}

	if !found {
		s.Log("No traps are visible")
	}
}
