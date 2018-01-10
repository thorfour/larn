package monster

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/game/state/character"
)

const (
	Bat = iota
	Gnome
	Hobgoblin
	Jackal
	Kobold
	Orc
	Snake
	Centipede
	Jaculi
	Troglodyte
	Ant
	Eye
	Leprechaun
	Nymph
	Quasit
	Rustmonster
	Zombie
	Assassinbug
	Bugbear
	Hellhound
	Icelizard
	Centaur
	Troll
	Yeti
	Whitedragon
	Elf
	Cube
	Metamorph
	Vortex
	Ziller
	Violetfungi
	Wraith
	Forvalaka
	Lamanobe
	Osequip
	Rothe
	Xorn
	Vampire
	Invisiblestalker
	Poltergeist
	Disenchantress
	Shamblingmound
	Yellowmold
	Umberhulk
	Gnomeking
	Mimic
	Waterlord
	Bronzedragon
	Greendragon
	Purpleworm
	Xvart
	Spiritnaga
	Silverdragon
	Platinumdragon
	Greenurchin
	Reddragon
	Demonlord
	Demonprince = 64 // requires special spawn
)

// Monster interface is used to represent monsters
type Monster interface {
	MoveTowards(*character.Character) // moves a monster towards a character and attacks if able
	Damage(*character.Character)      // character attacks the monster
	Fg() termbox.Attribute
	Bg() termbox.Attribute
	Visible(v bool)
	Displace() bool
}
