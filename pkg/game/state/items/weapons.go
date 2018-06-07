package items

import (
	"strconv"

	"github.com/thorfour/larn/pkg/game/state/stats"
)

type WeaponType int

const weaponRune = ')'

const (
	SunSword WeaponType = iota
	TwoHandedSword
	Spear
	Dagger
	BattleAxe
	LongSword
	Flail
	LanceOfDeath
	SwordOfSlashing
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
	Type      WeaponType // the type of weapon
	Attribute int        // the attributes of the weapon that add/subtract from the class
	DefaultItem
	NoStats
}

// Rune implements the io.Runeable interface
func (a *WeaponClass) Rune() rune {
	if a.Visibility {
		return weaponRune
	} else {
		return invisibleRune
	}
}

// Log implements the Loggable interface
func (a *WeaponClass) Log() string {
	return "You have found a " + a.String()
}

// String implements the Item interface
func (a *WeaponClass) String() string {
	if a.Attribute < 0 {
		return weaponName[a.Type] + " " + strconv.Itoa(a.Attribute)
	} else if a.Attribute > 0 {
		return weaponName[a.Type] + " +" + strconv.Itoa(a.Attribute)
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
	c.Wc += (weaponBase[a.Type] + a.Attribute)
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
	c.Wc -= (weaponBase[a.Type] + a.Attribute)
}

// GetNewWeapon returns a new default weapon of the given type
func GetNewWeapon(id WeaponType) *WeaponClass {
	return &WeaponClass{
		Type: id,
	}
}
