package state

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/thorfour/larn/pkg/game/state/character"
	"github.com/thorfour/larn/pkg/game/state/conditions"
	"github.com/thorfour/larn/pkg/game/state/items"
	"github.com/thorfour/larn/pkg/game/state/maps"
	"github.com/thorfour/larn/pkg/game/state/monster"
	"github.com/thorfour/larn/pkg/game/state/types"
	"github.com/thorfour/larn/pkg/io"
)

const (
	logLength = 5     // Ideally should be the same as the game.logLength but is useful to be definde separately for debug
	timeLimit = 30000 // max time to win a game
)

var (
	// ErrAlreadyDisplacedErr indicates player can't move to location
	ErrAlreadyDisplacedErr = fmt.Errorf("There's something here already")
	// ErrDidntWork player failed to cast a spell
	ErrDidntWork = fmt.Errorf("  It didn't seem to work")
)

type logring []string

// add adds a new log to the log ring
func (log logring) add(s string) logring {
	log = append(log, s)      // Append the new string
	if len(log) > logLength { // remove first element if the log exceeds length
		log = log[1:]
	}
	return log
}

// State holds all current game state
type State struct {
	StatLog    logring
	C          *character.Character
	Active     map[string]func()
	maps       *maps.Maps
	rng        *rand.Rand
	Taxes      int
	Name       string
	timeUsed   uint
	difficulty int
}

// New returns a new state and prints the welcome screen
func New(diff int) *State {
	log.Info("creating new state")
	s := new(State)
	s.difficulty = diff
	s.C = new(character.Character)
	s.C.Init(s.difficulty)
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
	defer s.update()
	if _, ok := s.C.Displaced.(maps.Empty); !ok { // Check if player is already displacing an object
		return nil, ErrAlreadyDisplacedErr
	}

	item, err := s.C.DropItem(e)
	if err != nil {
		return nil, err
	}
	s.C.Displaced = item
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
func (s *State) Move(d types.Direction) bool {
	defer s.maps.SetVisible(s.C)
	defer s.update()

	// Move the character
	moved, attacked := s.maps.Move(d, s.C)

	// The move results in an attack
	if attacked {
		s.playerAttack(d)
		return false
	}

	// If the character is displacing something add it to the status log
	if moved {
		switch t := s.C.Displaced.(type) {
		case *items.GoldPile:
			t.PickUp(s.C.Stats) // auto-pick up gold
			s.C.Displaced = s.maps.NewEmptyTile()
			s.Log(t.Log())
		case maps.Loggable:
			s.Log(t.Log())
		}
	}
	return moved
}

// Enter is used for entering into a building or dungeon/volcano
func (s *State) Enter() int {
	defer s.update()
	log.Debug("enter request")

	// Check if character is standing on an enterable object
	if t, ok := s.C.Displaced.(maps.Enterable); ok {
		lvl := t.Enter()
		if lvl < 0 { // Special cases for stores
			return lvl
		}
		s.maps.EnterLevel(s.C, t.Enter())
	}

	return 0
}

// PickUp will pick up the item the player is standing on
func (s *State) PickUp() {
	defer s.update()
	log.Debug("pickup request")

	i, ok := s.C.Displaced.(items.Item)
	if ok {
		i.PickUp(s.C.Stats)
		s.C.AddItem(i)
		s.C.Displaced = s.maps.NewEmptyTile()
	}
}

// Inventory request
func (s *State) Inventory() []string {
	log.Debug("inventory request")
	return s.C.Inventory()
}

// TimeStr returns the current time elapsed in the game
func (s *State) TimeStr() string {
	return fmt.Sprintf("Elapsed time is %v. You have %v mobuls left", (s.timeUsed+99)/100+1, s.TimeLeft())
}

// UseTime increment the amount of time used by t
func (s *State) UseTime(t uint) {
	s.timeUsed += t
}

// TimeLeft returns the amount of time a user has left in mobuls
func (s *State) TimeLeft() int {
	return int((timeLimit - s.timeUsed) / 100)
}

// Read is for the player to read a scroll or book
func (s *State) Read(e rune) error {
	defer s.update()
	log.Debug("read request")

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

// Cast casts the requested spell. May return a callback function
func (s *State) Cast(spell string) (func(types.Direction) bool, error) {
	defer s.update()
	var sp *items.Spell
	if !DEBUG {
		var err error
		sp, err = s.C.Cast(spell)
		if err != nil {
			return nil, err
		}

		if s.C.Cond.EffectActive(conditions.TimeStop) {
			return nil, ErrDidntWork
		}
	}

	if DEBUG { // Pass through the spell for debugging
		sp = &items.Spell{Code: spell}
	}

	switch sp.Code {
	//----------------------------------------------------------------------------
	//                            LEVEL 1 SPELLS
	//----------------------------------------------------------------------------
	case "pro": // protection
		if !s.C.Cond.EffectActive(conditions.SpellOfProtection) {
			s.C.Stats.Ac += 2 // protection field +2
		}
		s.C.Cond.Refresh(conditions.SpellOfProtection, 250, func() { s.C.Stats.Ac -= 2 })
	case "mle": // magic missile
		msg := "Your missile hit the %s"
		if s.C.Stats.Level >= 2 {
			msg = "Your missiles hit the %s"
		}
		dmg := rand.Intn((int(s.C.Stats.Level)+1)<<1) + int(s.C.Stats.Level) + 3
		return s.projectile(sp, dmg, msg, '+'), nil
	case "dex": // dexterity
		if !s.C.Cond.EffectActive(conditions.SpellOfDexterity) {
			s.C.Stats.Dex += 3
		}
		s.C.Cond.Refresh(conditions.SpellOfDexterity, 400, func() { s.C.Stats.Dex -= 3 })
	case "sle": // sleep
		hits := rand.Intn(3) + 2
		return s.directedHit(sp, s.hits(hits), fmt.Sprintf("While the %s slept, you smashed it %d times", "%s", hits)), nil
	case "chm": // charm monsters
		s.C.Cond.Refresh(conditions.CharmMonsters, int(s.C.Stats.Cha)<<1, nil)
	case "ssp": // sonic spear
		dmg := rand.Intn(10) + 16 + int(s.C.Stats.Level)
		return s.projectile(sp, dmg, "The sound damages the %s", '@'), nil
		//----------------------------------------------------------------------------
		//                            LEVEL 2 SPELLS
		//----------------------------------------------------------------------------
	case "web": // webs
		hits := rand.Intn(3) + 3
		return s.directedHit(sp, s.hits(hits), fmt.Sprintf("While the %s is entangled, you hit %d times", "%s", hits)), nil
	case "str": // strength
		if !s.C.Cond.EffectActive(conditions.SpellOfStrength) {
			s.C.Stats.Str += 3
		}
		s.C.Cond.Add(conditions.SpellOfStrength, 150+rand.Intn(100), func() { s.C.Stats.Str -= 3 })
	case "enl": // enlightenment
		s.maps.TouchAllInteriorCoordinates(func(obj io.Runeable) {
			if _, ok := obj.(types.Visibility); ok {
				obj.(types.Visibility).Visible(true)
			}
		})
	case "hel": // healing
		s.C.Heal(20 + int(s.C.Stats.Level<<1))
	case "cbl": // cure blindness
		s.C.Cond.Remove(conditions.Blindness)
	case "cre": // create monster
		// Select a random empty location next to the player to spawn the monster
		coords := s.maps.AdjacentCoords(s.C.Location())
		rand.Shuffle(len(coords), func(i, j int) {
			tmp := coords[j]
			coords[j] = coords[i]
			coords[i] = tmp
		})

		for _, c := range coords {
			if _, ok := s.maps.At(c).(maps.Displaceable); ok { // Found a displaceable object to place the monster onto
				mon := monster.New(monster.FromLevel(s.maps.CurrentLevel() + 1))
				mon.Visible(true)
				// TODO in the case of ROTHE, POLTERGEIST OR VAMPIRE stealth needs to be set on the monster
				// TODO figure out how monster stealth is utilized
				mon.Displaced = s.maps.Swap(c, mon)
				return nil, nil
			}
		}
	case "pha": // phantasmal forces
		if rand.Intn(11)+8 <= int(s.C.Stats.Wisdom) {
			return s.directedHit(sp, rand.Intn(20)+21+int(s.C.Stats.Level), "The %s believed!"), nil
		}
		s.Log("It didn't believe the illusions!")
	case "inv": // invsibility
		n := 0
		if am := s.C.CarryingSpecial(items.Amulet); am != nil { // Time added for amulet of invisibility
			n += 1 + am.Attr()
		}
		s.C.Cond.Refresh(conditions.Invisiblity, (n<<7)+12, nil)
		//----------------------------------------------------------------------------
		//                            LEVEL 3 SPELLS
		//----------------------------------------------------------------------------
	case "bal": // fireball
		dmg := rand.Intn(25+int(s.C.Stats.Level)) + 26 + int(s.C.Stats.Level)
		return s.projectile(sp, dmg, "A fireball hits the %s", '*'), nil
	case "cld": // cone of cold
		dmg := rand.Intn(25) + 21 + int(s.C.Stats.Level)
		return s.projectile(sp, dmg, "Your cone of cold strikes the %s", 'O'), nil
	case "ply": // polymorph
		return s.directedPolymorph(), nil
	case "can": // cancellation
		s.C.Cond.Refresh(conditions.Cancellation, 5+int(s.C.Stats.Level), nil)
	case "has": // haste self
		s.C.Cond.Refresh(conditions.HasteSelf, 7+int(s.C.Stats.Level), nil)
	case "ckl": // cloud kill
		s.omniDirect(sp, 31+rand.Intn(10), "The %s gasps for air")
	case "vpr": // vaporize rock
		//TODO may not be high level enough to break walls
		//TODO statues can drop books
		//TODO xorns take dmg from vpr
		//TODO thrones create gnome kings
		//TODO altars create demon princes
		//TODO fountains create waterlords
		s.maps.VaporizeAdjacent(s.C)
		//----------------------------------------------------------------------------
		//                            LEVEL 4 SPELLS
		//----------------------------------------------------------------------------
	case "dry": // dehydration
		return s.directedHit(sp, 100+int(s.C.Stats.Level), "The %s shrivels up"), nil
	case "lit": // lightning bolt
		dmg := (rand.Intn(25) + 1) + 20 + (int(s.C.Stats.Level) << 1)
		return s.projectile(sp, dmg, "A lightning bolt hits the %s", '~'), nil
	case "drl": // drain life
		i := int(math.Min(float64(s.C.Stats.Hp-1), float64(s.C.Stats.MaxHP/2)))
		s.C.Stats.Hp -= uint(i)
		return s.directedHit(sp, i+i, ""), nil
	case "glo": // globe of invulnerability
		if !s.C.Cond.EffectActive(conditions.GlobeOfInvul) {
			s.C.Stats.Ac += 10
		}
		if s.C.Stats.Intelligence > 3 { // globe decreases intelligence to minimum of 3
			s.C.Stats.Intelligence--
		}
		s.C.Cond.Add(conditions.GlobeOfInvul, 200, func() { s.C.Stats.Ac -= 10 })
	case "flo": // flood
		s.omniDirect(sp, 32+int(s.C.Stats.Level), "The %s struggles for air in your flood!")
	case "fgr": // finger of death
		if rand.Intn(150) == 63 {
			s.Log("Your heart stopped!")
			s.C.Stats.Hp = 0
			// TODO character died
			return nil, nil
		}

		if int(s.C.Stats.Wisdom) > rand.Intn(10)+11 {
			return s.directedHit(sp, 2000, "The %s's heart stopped"), nil
		}

		s.Log("It didn't work")
		//----------------------------------------------------------------------------
		//                            LEVEL 5 SPELLS
		//----------------------------------------------------------------------------
	case "sca": // scare monster
		s.C.Cond.Refresh(conditions.ScareMonster, rand.Intn(9)+1+int(s.C.Stats.Level), nil)
	case "hld": // hold monsters
		s.C.Cond.Add(conditions.HoldMonsters, rand.Intn(9)+1+int(s.C.Stats.Level), nil)
	case "stp": // time stop
		s.C.Cond.Add(conditions.TimeStop, rand.Intn(19)+1+(int(s.C.Stats.Level)<<1), nil)
	case "tel":
	case "mfi": // magic fire
		s.omniDirect(sp, 35+rand.Intn(9)+1+int(s.C.Stats.Level), "The %s cringes from the flame")
		//----------------------------------------------------------------------------
		//                            LEVEL 6 SPELLS
		//----------------------------------------------------------------------------
	case "sph":
	case "gen":
	case "sum":
	case "wtw":
	case "alt":
	case "per":
	default:
		return nil, ErrDidntWork
	}

	return nil, nil
}

// IdentTrap notifies the player if there are traps adjacent
func (s *State) IdentTrap() {
	defer s.update()

	// Get adjacent spaces
	adj := s.maps.Adjacent(s.C.Location())

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

// update function to handle time passage, spell decay and monster movement
func (s *State) update() {
	log.Debug("updating game state")
	if s.C.Cond.EffectActive(conditions.TimeStop) {
		s.C.Cond.Decay(conditions.TimeStop) // time stop, only thing to do is decay that spell
		return
	}

	// Move monsters
	s.moveMonsters()

	// increase the time used
	s.timeUsed++

	// Decay all active functions
	s.C.Cond.DecayAll()
}

func (s *State) moveMonsters() {
	log.Debug("move monsters")

	// Hold monsters, monsters don't move
	if s.C.Cond.EffectActive(conditions.HoldMonsters) {
		return
	}

	// TODO check for aggravate monsters
	// TODO check for stealth

	// Create a window from the current players position
	// c1 is the bottom left coordindate of a square, and c2 is the top right
	c := s.C.Location()
	c1 := types.Coordinate{int(c.X) - 5, int(c.Y) - 3}
	c2 := types.Coordinate{int(c.X) + 6, int(c.X) + 4}

	// Get a list of all monsters that appear in that window
	monsters := s.monstersInWindow(c1, c2)

	// Move all monsters in the window
	for _, m := range monsters {
		s.monsterMove(m)
	}
}

// TODO this window is still far to large
// monstersInWindow returns the list of coordinates of monsters withing a given section of the map
/*
	------c2
	|      |
	|      |
	|      |
	c1------
*/
// c1 is the lower left corner of a square and c2 is the upper right corner of a square
func (s *State) monstersInWindow(c1, c2 types.Coordinate) []types.Coordinate {
	log.WithFields(log.Fields{
		"c1": c1,
		"c2": c2,
	}).Debug("monster window")

	level := s.maps.CurrentMap()
	var ml []types.Coordinate

	// Walk through the window checking each space for a monster
	for i := c1.Y; i <= c2.Y; i++ {
		for j := c1.X; j <= c2.X; j++ {

			// Current coordinate within the window
			c := types.Coordinate{j, i}

			// First always check if the coordinate is within the map
			if !s.maps.ValidCoordinate(c) {
				continue
			}

			// Check if there is a monster at the coordinate
			if _, ok := level[c.Y][c.X].(*monster.Monster); ok {
				log.WithFields(log.Fields{
					"monster": string(level[c.Y][c.X].Rune()),
					"coord":   c,
				}).Debug("monster found")
				ml = append(ml, c) // add the monsters coordinate to the list
			}
		}
	}
	return ml
}

func (s *State) monsterMove(m types.Coordinate) {
	level := s.maps.CurrentMap()

	// Cast map location to a monster (this should never fail)
	mon := level[m.Y][m.X].(*monster.Monster)

	// Slow monsters only move every other tick
	if s.timeUsed&1 == 1 {
		switch mon.ID() {
		case monster.Troglodyte:
			fallthrough
		case monster.Hobgoblin:
			fallthrough
		case monster.Metamorph:
			fallthrough
		case monster.Xvart:
			fallthrough
		case monster.Invisiblestalker:
			fallthrough
		case monster.Icelizard:
			return
		}
	}

	// all spaces surrounding monster
	adj := s.maps.AdjacentCoords(m)

	// TODO handle scared monsters (randomly select valid location to move)
	// TODO handle intelligent monsters (they can navigate maze)

	//
	// Dumb monster movement (greedy)
	//

	// If the monster is already adjacent to the player attack player instead
	if s.maps.Distance(types.Coordinate(s.C.Location()), m) == 1 {
		s.attackPlayer(mon)
		return
	}

	// For each space calculate the space closest to the player
	minD := 10000
	var minC types.Coordinate
	for _, c := range adj {
		if _, ok := level[c.Y][c.X].(maps.Displaceable); !ok { // Invalid movement location
			log.WithField("coord", c).Debug("not displaceable")
			continue
		}

		if d := s.maps.Distance(s.C.Location(), c); d < minD {
			minD = d
			minC = c
		}
	}

	log.WithFields(log.Fields{
		"monster":   string(level[m.Y][m.X].Rune()),
		"min coord": minC,
		"distance":  minD,
	}).Debug("moving monster")

	// Perform the move
	level[m.Y][m.X] = mon.Displaced
	mon.Displaced = level[minC.Y][minC.X]
	level[minC.Y][minC.X] = mon
}

// attackPlayer attempts an attack on the player from the monster
func (s *State) attackPlayer(mon *monster.Monster) {
	// TODO check for negatespirit or spirit pro against poltergeis and naga
	// TODO cubeundead or undeadpro against vampire, wraith, zombie

	mName := s.monsterName(mon)

	// If character is invisble chance to miss
	if s.C.Cond.EffectActive(conditions.Invisiblity) {
		if rand.Intn(33) < 20 {
			s.Log(fmt.Sprintf("The %s misses wildly", mName))
			return
		}
	}

	if s.C.Cond.EffectActive(conditions.CharmMonsters) {
		if rand.Intn(30)+5*mon.Info.Lvl-int(s.C.Stats.Cha) < 30 {
			s.Log(fmt.Sprintf("The %s is awestruct at your magnificence!", mName))
			return
		}
	}

	s.hitPlayer(mon)
}

// hitPlayer deals the damage from a monster to a player
func (s *State) hitPlayer(mon *monster.Monster) {
	dmg := mon.BaseDamage()

	if mon.Info.Attack > 0 {
		if dmg+s.difficulty+8 > s.C.Stats.Ac || s.C.Stats.Ac <= 0 || rand.Intn(s.C.Stats.Ac) == 0 { // Check for special attack success
			// TODO check for special attack
			/*
				if special() {
					return
				}
			*/

			s.difficulty -= 2
		}
	}

	// No special attack, deal normal damage
	if (dmg+s.difficulty) > s.C.Stats.Ac || s.C.Stats.Ac <= 0 || rand.Intn(s.C.Stats.Ac) == 0 {
		s.Log(fmt.Sprintf("The %v hit you", s.monsterName(mon)))
		if s.C.Stats.Ac < dmg {
			s.C.Damage(dmg - s.C.Stats.Ac)
		}
	}

	s.Log(fmt.Sprintf("The %s missed", s.monsterName(mon)))
}

// playerAttack deals damage to a monster
func (s *State) playerAttack(d types.Direction) {

	// Get monster location NOTE: this isn't moving the character, just calculating the coordinate
	mLoc := types.Move(s.C.Location(), d)

	// Get the monster at the attempted location
	m := s.maps.At(mLoc)
	switch mon := m.(type) {
	case *monster.Monster: // nominal case
		// Deal damage to the monster
		dead := s.hitMonster(mon)
		if dead {
			s.Log(fmt.Sprintf("The %s died", s.monsterName(mon)))
			s.maps.RemoveAt(mLoc)                               // remove the mosnter at the location
			s.maps.CurrentMap()[mLoc.Y][mLoc.X] = mon.Displaced // replace the any items displaced by the monster
			s.monsterDrop(mLoc, mon)                            // have the monster drop gold/items
			if s.C.GainExperience(mon.Info.Experience) {
				s.Log(fmt.Sprintf("Welcome to level %d", s.C.Stats.Level))
			}
		}
	default:
		log.WithField("object", m).Error("attached non attackable object")
		return
	}

}

// hitMonster handles a charachter attempting to hit a monster
func (s *State) hitMonster(m *monster.Monster) bool {
	dead := false
	if s.C.Cond.EffectActive(conditions.TimeStop) {
		return dead
	}

	tmp := m.Info.Armor + int(s.C.Stats.Level) + int(s.C.Stats.Dex) + s.C.Stats.Wc/4 - 12
	if rand.Intn(20) < tmp-s.difficulty || rand.Intn(71) < 5 { // some random chance to hit
		s.Log(fmt.Sprintf("You hit the %s", s.monsterName(m)))
		dmg := s.hits(1)
		if dmg < 9999 {
			dmg = rand.Intn(dmg) + 1
		}

		log.WithFields(log.Fields{
			"monster": string(m.Rune()),
			"damage":  dmg,
		}).Debug("damanged monster")

		_, dead = m.Damage(dmg)
	} else {
		s.Log(fmt.Sprintf("You missed the %s", s.monsterName(m)))
	}

	// TODO handle dulled weapons
	// TODO handle turning vampires back into bats
	return dead
}

// hits returns the damage dealt for the given number of hits
func (s *State) hits(n int) int {
	if n <= 0 || n > 20 { // out of range
		return 0
	}

	if s.wieldingLance() {
		return 10000
	}

	c := s.C.Stats
	dmg := n * ((c.Wc >> 1) + int(c.Str) + c.StrExtra - s.difficulty - 12 + c.MoreDmg)
	if dmg >= 1 {
		return dmg
	}

	return n
}

// monsterDrop performs a item/gold drop from a slain monster at a given location.
// NOTE: in OG larn the items were always dropped next to the player. This version drops next to the monster
func (s *State) monsterDrop(c types.Coordinate, m *monster.Monster) {
	amt := m.Info.Gold
	if amt > 0 {
		amt = rand.Intn(amt) + amt
	}
	gp := &items.GoldPile{Amount: amt}
	if gp.Amount > 0 {
		s.drop(c, gp) // drop gold pile
	}

	var drop []items.Item
	switch m.ID() {
	case monster.Orc:
		fallthrough
	case monster.Nymph:
		fallthrough
	case monster.Elf:
		fallthrough
	case monster.Troglodyte:
		fallthrough
	case monster.Troll:
		fallthrough
	case monster.Rothe:
		fallthrough
	case monster.Violetfungi:
		fallthrough
	case monster.Platinumdragon:
		fallthrough
	case monster.Gnomeking:
		fallthrough
	case monster.Reddragon:
		drop = items.CreateItems(s.maps.CurrentLevel())
	case monster.Leprechaun:
		if rand.Intn(101) >= 75 {
			drop = append(drop, items.CreateGem())
		}
		for i := rand.Intn(5); i == 0; i = rand.Intn(5) {
			if rand.Intn(101) >= 75 {
				drop = append(drop, items.CreateGem())
			}
		}
	}
	for i := range drop { // drop items
		s.drop(c, drop[i])
	}
}

// dropAdjacent finds a location to drop an item
func (s *State) drop(c types.Coordinate, drop io.Runeable) {
	// Drop in location if coordinate is empty
	if _, ok := s.maps.CurrentMap()[c.Y][c.X].(maps.Empty); ok {
		s.maps.CurrentMap()[c.Y][c.X] = drop
		return
	}

	// Look for empty adjacent locations to drop
	for _, a := range s.maps.Adjacent(c) {
		if _, ok := a.(maps.Empty); ok {
			s.maps.CurrentMap()[c.Y][c.X] = drop
			return
		}
	}

	// NOTE: If we couldn't find a place to drop then nothing gets dropped
}

// Quaff performs a drink potion action
func (s *State) Quaff(e rune) (func() bool, error) {
	defer s.update()
	log.Debug("quaff requested")

	l, id, err := s.C.Quaff(e)
	if err != nil {
		return nil, err
	}

	// Log all the information that read returned
	for _, r := range l {
		s.Log(r)
	}

	// Special potion cases
	switch id {
	case items.TreasureFinding:
		// Don't if blind
		if s.C.Cond.EffectActive(conditions.Blindness) {
			return nil, nil
		}
		s.maps.TouchAllInteriorCoordinates(func(obj io.Runeable) {
			switch obj.(type) {
			case items.Gemstone:
				if _, ok := obj.(types.Visibility); ok {
					obj.(types.Visibility).Visible(true)
				}
			case items.Gold:
				if _, ok := obj.(types.Visibility); ok {
					obj.(types.Visibility).Visible(true)
				}
			}
		})
	case items.MonsterDetection:
		// Don't if blind
		if s.C.Cond.EffectActive(conditions.Blindness) {
			return nil, nil
		}
		s.maps.TouchAllInteriorCoordinates(func(obj io.Runeable) {
			if _, ok := obj.(monster.Interface); ok {
				if _, ok := obj.(types.Visibility); ok {
					obj.(types.Visibility).Visible(true)
				}
			}
		})
	case items.ObjectDetection:
		// Don't if blind
		if s.C.Cond.EffectActive(conditions.Blindness) {
			return nil, nil
		}
		s.maps.TouchAllInteriorCoordinates(func(obj io.Runeable) {
			switch obj.(type) {
			case items.Gemstone:
			case items.Gold:
				// no gems or gold piles
			case items.Item:
				if _, ok := obj.(types.Visibility); ok {
					obj.(types.Visibility).Visible(true)
				}
			}
		})
	case items.Forgetfulness:
		s.maps.TouchAllInteriorCoordinates(func(obj io.Runeable) {
			if _, ok := obj.(types.Visibility); ok {
				obj.(types.Visibility).Visible(false)
			}
		})
	case items.Sleep:
		// Return a callback function
		i := rand.Intn(11) + 1 - (int(s.C.Stats.Con) >> 2) + 2
		return func() bool {
			if i > 0 {
				i--
				s.update()
				s.maps.SetVisible(s.C)
				log.Info("sleeping from potion")
				time.Sleep(time.Second)
				return true
			}
			s.Log("You woke up!")
			return false
		}, nil
	}

	return nil, nil
}

// projectile returns a callback function that will handle the animation of a projectile
func (s *State) projectile(spell *items.Spell, dmg int, msg string, c rune) func(types.Direction) bool {
	current := s.C.Location()
	var obj io.Runeable
	return func(d types.Direction) bool {
		cleanup := func() {
			if obj != nil {
				s.maps.Swap(current, obj)
			}
		}
		if dmg <= 0 { // projectile ran out of power
			cleanup()
			return false
		}

		// Update state with location of projectile
		if obj != nil { // replace the object that was displaced
			s.maps.Swap(current, obj) // TODO replaced objects should probably be visible?
		}
		current = types.Move(current, d)
		if s.maps.OutOfBounds(current) { // If the projectile would go off the map, or into a dungeon wall
			return false
		}
		obj = s.maps.Swap(current, &items.ProjectileSpell{R: c})

		// Object collision handling
		switch o := obj.(type) {
		case *maps.Empty:
			dmg -= (3 + (s.difficulty >> 1)) // reduce power for each space traveled
		case *maps.Wall:
			msg = fmt.Sprintf(msg, "wall")
			bonusDmg := 0
			if DEBUG {
				bonusDmg = 100000
			}
			// Enough damage to destroy the wall?
			if (dmg+bonusDmg >= 50+s.difficulty) && s.maps.CurrentLevel() < maps.MaxVolcano && !s.maps.OuterWall(current) {
				msg = fmt.Sprintf(msg + "  The wall crumbles")
				s.maps.Swap(current, &maps.Empty{})
			} else {
				cleanup()
			}
			s.Log(msg)
			return false
		case *monster.Monster:
			s.Log(fmt.Sprintf(msg, s.monsterName(o)))
			dealt, dead := s.damageMonster(dmg, o, current)
			if dead {
				obj = nil
			}

			dmg -= dealt
		default:
			// TODO probably panic here
			dmg -= (3 + (s.difficulty >> 1)) // reduce power for each space traveled
		}

		return true
	}
}

// omniDirect deals damange to all adjacent coordinates to the character
func (s *State) omniDirect(spell *items.Spell, dmg int, msg string) {
	for _, c := range s.maps.AdjacentCoords(s.C.Location()) {
		obj := s.maps.At(c)
		switch o := obj.(type) {
		case *monster.Monster:
			s.Log(fmt.Sprintf(msg, s.monsterName(o)))
			s.damageMonster(dmg, o, c)
		}
	}
}

// directedPolymorph attempts to polymorph a monster in a given direction
func (s *State) directedPolymorph() func(types.Direction) bool {
	if s.C.Cond.EffectActive(conditions.Confusion) { // Do nothing if confused
		return nil
	}

	return func(d types.Direction) bool {
		monLoc := types.Move(s.C.Location(), d)
		obj := s.maps.At(monLoc)
		switch obj.(type) {
		case *monster.Monster:
			mon := monster.New(monster.Random())
			mon.Visible(true)
			s.maps.Swap(monLoc, mon)
		default:
			s.Log("There wasn't anything there!")
		}
		return false
	}
}

// directedHit attempts to damange a monster in a given direction
func (s *State) directedHit(spell *items.Spell, dmg int, msg string) func(types.Direction) bool {
	if s.C.Cond.EffectActive(conditions.Confusion) { // Do nothing if confused
		return nil
	}

	// TODO handle if spell affects the monster

	return func(d types.Direction) bool {
		monLoc := types.Move(s.C.Location(), d)
		obj := s.maps.At(monLoc)
		switch o := obj.(type) {
		case *monster.Monster:
			if msg != "" {
				s.Log(fmt.Sprintf(msg, s.monsterName(o)))
			}
			s.damageMonster(dmg, o, monLoc)
		case *items.Mirror:
			// TODO handle hitting a mirror
		default:
			s.Log("There wasn't anything there!")
		}
		return false
	}
}

// wieldingLance returns true if the character is wielding the lance of death
func (s *State) wieldingLance() bool {
	if w, ok := s.C.Wielding().(*items.WeaponClass); ok {
		if w.Type == items.LanceOfDeath {
			return true
		}
	}

	return false
}

func (s *State) damageMonster(dmg int, m *monster.Monster, loc types.Coordinate) (int, bool) {
	dealt, dead := m.Damage(dmg)
	if dead {
		// TODO handle gaining exp for killing a monster
		s.maps.Swap(loc, m.Displaced)
		s.Log(fmt.Sprintf("The %s died!", s.monsterName(m)))
	}

	return dealt, dead
}

// monsterName returns the name of the monster, handles if the character is blind
func (s *State) monsterName(m *monster.Monster) string {
	if s.C.Cond.EffectActive(conditions.Blindness) {
		return "monster"
	}

	return m.Name()
}
