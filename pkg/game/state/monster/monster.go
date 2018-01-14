package monster

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/game/state/character"
)

// Monster interface is used to represent monsters
type Monster interface {
	MoveTowards(*character.Character) // moves a monster towards a character and attacks if able
	Damage(*character.Character)      // character attacks the monster
	Fg() termbox.Attribute
	Bg() termbox.Attribute
	Visible(v bool)
}
