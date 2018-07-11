package items

import (
	"fmt"
	"math/rand"
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
