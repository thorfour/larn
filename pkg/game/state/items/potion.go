package items

import (
	"fmt"
	"math/rand"

	"github.com/golang/glog"
	"github.com/thorfour/larn/pkg/game/state/stats"
)

const (
	potionRune = '!'
)

// PotionID to indetify a potion
type PotionID int

const (
	// Sleep potion
	Sleep PotionID = iota
	// Healing potion
	Healing
	// RaiseLevel potion
	RaiseLevel
	// IncreaseAbility potion
	IncreaseAbility
	// GainWisdom potion
	GainWisdom
	// GainStrength potion
	GainStrength
	// IncreaseCharisma potion
	IncreaseCharisma
	// Dizziness potion
	Dizziness
	// Learning potion
	Learning
	// ObjectDetection potion
	ObjectDetection
	// MonsterDetection potion
	MonsterDetection
	// Forgetfulness potion
	Forgetfulness
	// Water potion
	Water
	// Blindness potion
	Blindness
	// Confusion potion
	Confusion
	// Heroism potion
	Heroism
	// Sturdiness potion
	Sturdiness
	// GiantStrength potion
	GiantStrength
	// FireResistance potion
	FireResistance
	// TreasureFinding potion
	TreasureFinding
	// InstantHealing potion
	InstantHealing
	// CureDianthroritis potion
	CureDianthroritis
	// Poison potion
	Poison
	// SeeInvisible potion
	SeeInvisible
)

/*
 *  LUT to return a potion ID created with appropriate probability
 *  of occurrence
 *
 *  0 - sleep               1 - healing                 2 - raise level
 *  3 - increase ability    4 - gain wisdom             5 - gain strength
 *  6 - increase charisma   7 - dizziness               8 - learning
 *  9 - object detection    10 - monster detection      11 - forgetfulness
 *  12 - water              13 - blindness              14 - confusion
 *  15 - heroism            16 - sturdiness             17 - giant strength
 *  18 - fire resistance    19 - treasure finding       20 - instant healing
 *  21 - cure dianthroritis 22 - poison                 23 - see invisible
 */
var potprob = []PotionID{0, 0, 1, 1, 1, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 9,
	10, 10, 10, 11, 11, 12, 12, 13, 14, 15, 16, 17, 18, 19, 19, 19,
	20, 20, 22, 22, 23, 23}

var potionname = []string{
	"sleep",
	"healing",
	"raise level",
	"increase ability",
	"wisdom",
	"strength",
	"raise charisma",
	"dizziness",
	"learning",
	"object detection",
	"monster detection",
	"forgetfulness",
	"water",
	"blindness",
	"confusion",
	"heroism",
	"sturdiness",
	"giant strength",
	"fire resistance",
	"treasure finding",
	"instant healing",
	"cure dianthroritis",
	"poison",
	"see invisible",
}

// knownPotions map of all potions the player has learned
var knownPotions map[PotionID]bool

// Potion that a player may drink for an effect
type Potion struct {
	ID    PotionID
	Store bool // indicates if this potion is on display in the DND store (for name display purposes)
	DefaultItem
	NoStats
}

// Rune implements the io.Runeable interface
func (p *Potion) Rune() rune {
	if p.Visibility {
		return potionRune
	}
	return invisibleRune
}

// Log implements the Disaplceable interface
func (p *Potion) Log() string {
	if knownPotions[p.ID] {
		return fmt.Sprintf("You have found a magic potion of %s", potionname[p.ID])
	}
	return "You have found a magic potion"
}

// String implements the Item interface
func (p *Potion) String() string {
	if knownPotions[p.ID] || p.Store {
		return fmt.Sprintf("a potion of %s", potionname[p.ID])
	}
	return "a potion"
}

// KnownPotion returns true if the player knows the potion (used for tradig post)
func KnownPotion(p PotionID) bool {
	return knownPotions[p]
}

// LearnPotion adds a potion to the list of known potions
func LearnPotion(p PotionID) {
	knownPotions[p] = true
}

// ForgetPotion removes a potion from the list of known potions
func ForgetPotion(p PotionID) {
	delete(knownPotions, p)
}

// NewPotion randomly returns a new potion
func NewPotion() *Potion {
	return &Potion{
		ID: potprob[rand.Intn(len(potprob))],
	}
}

// Quaff implemtents the Quaffable interface. Applies a potions effects to the given stats. Returns a log of events
func (p *Potion) Quaff(s *stats.Stats) []string {
	LearnPotion(p.ID)
	switch p.ID {
	case Sleep:
		// TODO
		return nil
	case Healing:
		if s.Hp == s.MaxHP { // if at max HP, raise max HP by 1
			s.RaiseMaxHP(1)
		} else { // heal the player
			s.GainHP(uint(rand.Intn(20)+1) + 20 + s.Level)
		}
		return []string{"You feel better"}
	case RaiseLevel:
		s.Level++
		s.RaiseMaxHP(1)
		return []string{"Suddenly, you feel much more skillful!"}
	case IncreaseAbility:
		// add 1 to random attribute
		switch rand.Intn(6) {
		case 0:
			s.Cha++
		case 1:
			s.Wisdom++
		case 2:
			s.Con++
		case 3:
			s.Dex++
		case 4:
			s.Str++
		case 5:
			s.Intelligence++
		}
		return []string{"You feel strange for a moment"}
	case GainWisdom:
		s.Wisdom += uint(rand.Intn(2)) + 1
		return []string{"You feel more self confident!"}
	case GainStrength:
		if s.Str < 12 {
			s.Str = 12
		} else {
			s.Str++
		}
		return []string{"Wow! You feel great!"}
	case IncreaseCharisma:
		s.Cha++
		return []string{"Your charm went up by one!"}
	case Dizziness:
		s.Str--
		if s.Str < 3 {
			s.Str = 3
		}
		return []string{"You become dizzy!"}
	case Learning:
		s.Intelligence++
		return []string{"Your intelligence went up by one!"}
	case ObjectDetection:
		// TODO don't reveal anything if player is blind
		// TODO reveal all items on a level
		return []string{"You sense the presence of objects!"}
	case MonsterDetection:
		// TODO reveal all monster unless blind
		return []string{"You sense the presence of monsters!"}
	case Forgetfulness:
		// TODO hide all objects on map
		return []string{"You stagger for a moment . ."}
	case Water:
		return []string{"This potion has no taste to it"}
	case Blindness:
		// TODO add active blind
		return []string{"You can't see anything!"}
	case Confusion:
		// TODO add active confusion
		return []string{"You feel confused"}
	case Heroism:
		// TODO check if already heroic
		// TODO add to heroism decay func +250
		s.Cha += 11
		s.Wisdom += 11
		s.Con += 11
		s.Dex += 11
		s.Str += 11
		s.Intelligence++
		return []string{"WOW!! You feel Super-fantastic!!!"}
	case Sturdiness:
		s.Con++
		return []string{"You have a greater intestinal constitude!"}
	case GiantStrength:
		/// TODO if gianstr active +21 strextra
		// TODO add giantstr decay func
		s.Str += 700
		return []string{"You now have incredibly bulgin muscles!!!"}
	case FireResistance:
		// TODO add fire resist active +1000
		return []string{"You feel a chill run up your spine!"}
	case TreasureFinding:
		// TODO reveal all diamonds and piles of gold
		return []string{"You feel greedy . . ."}
	case InstantHealing:
		s.Hp = s.MaxHP
		return nil
	case CureDianthroritis:
		return []string{"You don't seem to be affected"}
	case Poison:
		// TODO add half damage active += 200 + rand.Intn(200)
		return []string{"You feel a sickness engulf you"}
	case SeeInvisible:
		// TODO add see invisible active
		return []string{"You feel your vision sharpen"}
	default:
		// TODO log an error
		glog.Error("unknown potion consumed: %v", p.ID)
		return nil
	}
}
