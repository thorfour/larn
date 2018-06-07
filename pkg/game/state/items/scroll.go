package items

import (
	"fmt"
	"math/rand"
)

const (
	scrollRune = '?'
)

// ScrollID identifies the type of scroll
type ScrollID int

const (
	// EnchantArmor scroll
	EnchantArmor ScrollID = iota
	// EnchantWeapon scroll
	EnchantWeapon
	// Englightenment scroll
	Englightenment
	// Paper scroll
	Paper
	// CreateMonster scroll
	CreateMonster
	// CreateItem scroll
	CreateItem
	// Aggravate scroll
	Aggravate
	// TimeWarp scroll
	TimeWarp
	// Teleportation scroll
	Teleportation
	// ExpandedAwareness scroll
	ExpandedAwareness
	// HasteMonster scroll
	HasteMonster
	// HealMonster scroll
	HealMonster
	// SpiritProtection scroll
	SpiritProtection
	// UndeadProtection scroll
	UndeadProtection
	// Stealth scroll
	Stealth
	// MagicMapping scroll
	MagicMapping
	// HoldMonster scroll
	HoldMonster
	// GemPerfection scroll
	GemPerfection
	// SpellExtension scroll
	SpellExtension
	// Identify scroll
	Identify
	// RemoveCurse scroll
	RemoveCurse
	// Annihilation scroll
	Annihilation
	// Pulverization scroll
	Pulverization
	// LifeProtection scroll
	LifeProtection
)

var scrollname = []string{
	"enchant armor",
	"enchant weapon",
	"enlightenment",
	"blank paper",
	"create monster",
	"create artifact",
	"aggravate monsters",
	"time warp",
	"teleportation",
	"expanded awareness",
	"haste monsters",
	"monster healing",
	"spirit protection",
	"undead protection",
	"stealth",
	"magic mapping",
	"hold monsters",
	"gem perfection",
	"spell extension",
	"identify",
	"remove curse",
	"annihilation",
	"pulverization",
	"life protection",
}

/*
 *  LUT to create scroll IDs with appropriate probability of
 *  occurrence
 *
 *  0 - armor           1 - weapon      2 - enlightenment   3 - paper
 *  4 - create monster  5 - create item 6 - aggravate       7 - time warp
 *  8 - teleportation   9 - expanded awareness              10 - haste monst
 *  11 - heal monster   12 - spirit protection      13 - undead protection
 *  14 - stealth        15 - magic mapping          16 - hold monster
 *  17 - gem perfection 18 - spell extension        19 - identify
 *  20 - remove curse   21 - annihilation           22 - pulverization
 *  23 - life protection
 */
var scprob = []ScrollID{0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 3, 3,
	3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 9, 9,
	9, 9, 10, 10, 10, 10, 11, 11, 11, 12, 12, 12, 13, 13, 13, 13, 14, 14,
	15, 15, 16, 16, 16, 17, 17, 18, 18, 19, 19, 19, 20, 20, 20, 20, 21, 22,
	22, 22, 23}

// knownScrolls a list of all the scrolls a player has discovered
var knownScrolls map[ScrollID]bool

// Scroll a player can read
type Scroll struct {
	ID ScrollID
	DefaultItem
	NoStats
}

// LearnScroll marks a scroll has having been learned (via reading or identify)
func LearnScroll(id ScrollID) {
	knownScrolls[id] = true
}

// ForgetScroll forget having learned a scroll
func ForgetScroll(id ScrollID) {
	delete(knownScrolls, id)
}

// Rune implements the io.Runeable interface
func (s *Scroll) Rune() rune {
	if s.Visibility {
		return scrollRune
	}
	return invisibleRune
}

// Log implements the Disaplceable interface
func (s *Scroll) Log() string {
	return "You have found a magic scroll"
}

// String implements the Item interface
func (s *Scroll) String() string {
	if knownScrolls[s.ID] {
		return fmt.Sprintf("a scroll of %s", idToName(s.ID))
	}
	return "a scroll"
}

// NewScroll returns a random scroll
func NewScroll() *Scroll {
	return &Scroll{
		ID: scprob[rand.Intn(len(scprob))],
	}
}

// idToName returns the string representation of the scroll
func idToName(id ScrollID) string {
	return scrollname[id]
}
