package items

import (
	"math/rand"
	"strconv"

	"github.com/thorfour/larn/pkg/game/state/stats"
)

// WeaponType indicate the type of weapon
type WeaponType int

const weaponRune = ')'

const (
	// SunSword weapon
	SunSword WeaponType = iota
	// TwoHandedSword weapon
	TwoHandedSword
	// Spear weapon
	Spear
	// Dagger weapon
	Dagger
	// BattleAxe weapon
	BattleAxe
	// LongSword weapon
	LongSword
	// Flail weapon
	Flail
	// LanceOfDeath weapon
	LanceOfDeath
	// SwordOfSlashing weapon
	SwordOfSlashing
	// BessmansHammer weapon
	BessmansHammer
)

// Map of all the weapon base values
var weaponBase = map[WeaponType]int{
	Dagger:          3,
	Spear:           10,
	Flail:           14,
	BattleAxe:       17,
	LongSword:       22,
	TwoHandedSword:  26,
	SunSword:        32,
	SwordOfSlashing: 30,
	BessmansHammer:  35,
	LanceOfDeath:    19,
}

// Map of all the displayable weapon names
var weaponName = map[WeaponType]string{
	SunSword:        "sun sword",
	TwoHandedSword:  "two handed sword",
	Spear:           "spear",
	Dagger:          "dagger",
	BattleAxe:       "battleaxe",
	LongSword:       "long sword",
	Flail:           "flail",
	LanceOfDeath:    "lance of death",
	SwordOfSlashing: "sword of slashing",
	BessmansHammer:  "Bessman's flailing hammer",
}

// WeaponClass satisfies the item interface as well as the Weapon Interface
type WeaponClass struct {
	Type WeaponType // the type of weapon
	DefaultAttribute
	DefaultItem
	NoStats
}

// Rune implements the io.Runeable interface
func (a *WeaponClass) Rune() rune {
	if a.Visibility {
		return weaponRune
	}
	return invisibleRune
}

// Log implements the Loggable interface
func (a *WeaponClass) Log() string {
	return "You have found a " + a.String()
}

// String implements the Item interface
func (a *WeaponClass) String() string {
	if a.Attr() < 0 {
		return weaponName[a.Type] + " " + strconv.Itoa(a.Attr())
	} else if a.Attr() > 0 {
		return weaponName[a.Type] + " +" + strconv.Itoa(a.Attr())
	}
	return weaponName[a.Type]
}

// Wield implements the Weapon interface
func (a *WeaponClass) Wield(c *stats.Stats) {
	switch a.Type { // Special weapon handling
	case SwordOfSlashing:
		c.Dex += 5 // sword of slashing increases dexterity
	case BessmansHammer:
		c.Dex += 10
		c.Str += 10
		c.Intelligence -= 10 // hammers make you stupid
	}
	c.Wc += (weaponBase[a.Type] + a.Attr())
}

// Disarm implements the Weapon interface
func (a *WeaponClass) Disarm(c *stats.Stats) {
	switch a.Type { // Special weapon handling
	case SwordOfSlashing:
		c.Dex -= 5 // sword of slashin increases dexterity
	case BessmansHammer:
		c.Dex -= 10
		c.Str -= 10
		c.Intelligence += 10 // hammers make you stupid
	}
	c.Wc -= (weaponBase[a.Type] + a.Attr())
}

// GetNewWeapon returns a new default weapon of the given type
func GetNewWeapon(id WeaponType, l int) *WeaponClass {
	attr := 0
	switch id {
	case Dagger:
		x := rand.Intn(13)
		switch {
		case x < 3:
		case x < 7:
			attr = 1
		case x < 9:
			attr = 2
		case x < 11:
			attr = 3
		case x < 12:
			attr = 4
		default:
			attr = 5
		}
	case Spear:
		fallthrough
	case BattleAxe:
		fallthrough
	case TwoHandedSword:
		attr = rand.Intn(l/3 + 1)
	case Flail:
		attr = rand.Intn(l/2 + 1)
	case LongSword:
		x := rand.Intn(13)
		switch {
		case x < 6:
		case x < 11:
			attr = 1
		case x < 12:
			attr = 2
		case x < 13:
			attr = 3
		}
	}
	ac := &WeaponClass{
		Type: id,
	}
	ac.ResetAttr(attr)

	return ac
}
