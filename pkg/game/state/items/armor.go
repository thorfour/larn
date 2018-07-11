package items

import (
	"math/rand"
	"strconv"

	"github.com/thorfour/larn/pkg/game/state/stats"
)

// ArmorType indicates the type of armor
type ArmorType int

const armorRune = '['

const (
	// Leather armor
	Leather ArmorType = iota
	// StuddedLeather armor
	StuddedLeather
	// RingMail armor
	RingMail
	// ChainMail armor
	ChainMail
	// SplintMail armor
	SplintMail
	// PlateMail armor
	PlateMail
	// PlateArmor armor
	PlateArmor
	// StainlessPlateArmor armor
	StainlessPlateArmor
)

// Map of all the armor base values
var armorBase = map[ArmorType]int{
	Leather:             2,
	StuddedLeather:      3,
	RingMail:            5,
	ChainMail:           6,
	SplintMail:          7,
	PlateMail:           9,
	PlateArmor:          10,
	StainlessPlateArmor: 12,
}

// Map of all the displayable armor names
var armorName = map[ArmorType]string{
	Leather:             "leather",
	StuddedLeather:      "studded leather",
	RingMail:            "ring mail",
	ChainMail:           "chain mail",
	SplintMail:          "splint mail",
	PlateMail:           "plate mail",
	PlateArmor:          "plate armor",
	StainlessPlateArmor: "stainless plate armor",
}

// ArmorClass satisfies the item interface as well as the Armor Interface
type ArmorClass struct {
	Type ArmorType // the type of armor
	DefaultAttribute
	DefaultItem
	NoStats
}

// Rune implements the io.Runeable interface
func (a *ArmorClass) Rune() rune {
	if a.Visibility {
		return armorRune
	}
	return invisibleRune
}

// Log implements the Loggable interface
func (a *ArmorClass) Log() string {
	return "You have found a " + a.String()
}

// String implements the Item interface
func (a *ArmorClass) String() string {
	if a.Attr() < 0 {
		return armorName[a.Type] + " " + strconv.Itoa(a.Attr())
	} else if a.Attr() > 0 {
		return armorName[a.Type] + " +" + strconv.Itoa(a.Attr())
	}
	return armorName[a.Type]
}

// Wear implements the Armor interface
func (a *ArmorClass) Wear(c *stats.Stats) {
	c.Ac += (armorBase[a.Type] + a.Attr())

}

// TakeOff implements the Armor interface
func (a *ArmorClass) TakeOff(c *stats.Stats) {
	c.Ac -= (armorBase[a.Type] + a.Attr())
}

// NewArmor returns a new defauly armor of type id
func NewArmor(id ArmorType, l int) *ArmorClass {
	attr := 0
	switch id {
	case StuddedLeather:
		fallthrough
	case SplintMail:
		fallthrough
	case RingMail:
		attr = rand.Intn(l/2 + 1)
	case Leather:
		x := rand.Intn(15) // TODO this should adjust for game difficulty
		switch {
		case x < 5:
		case x < 7:
			attr = 1
		case x < 9:
			attr = 2
		case x < 11:
			attr = 3
		case x < 12:
			attr = 4
		case x < 13:
			attr = 5
		case x < 14:
			attr = 6
		case x < 15:
			attr = 7
		}
	case ChainMail:
		x := rand.Intn(10)
		switch {
		case x < 3:
		case x < 6:
			attr = 1
		case x < 8:
			attr = 2
		case x < 9:
			attr = 3
		default:
			attr = 4
		}
	case PlateMail:
		x := rand.Intn(10)
		switch {
		case x < 4:
		case x < 6:
			attr = 1
		case x < 8:
			attr = 2
		case x < 9:
			attr = 3
		default:
			attr = 4
		}
	}
	ac := &ArmorClass{
		Type: id,
	}
	ac.ResetAttr(attr)

	return ac
}
