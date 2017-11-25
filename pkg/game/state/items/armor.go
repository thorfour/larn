package items

import "github.com/thorfour/larn/pkg/game/state/stats"

// armorclass satisfies the item interface as well as the armor Interface
type armorclass struct {
	class uint
	name  string
}

var (
	shield = armorclass{2, "shield"}

	leather     = armorclass{2, "leather armor"}
	studleather = armorclass{3, "studded leather armor"}
	ring        = armorclass{5, "ring mai"}
	chain       = armorclass{6, "chain mail"}
	splint      = armorclass{7, "splint mail"}
	plateMail   = armorclass{9, "plate mail"}
	plateArmor  = armorclass{10, "plate armor"}
	stainless   = armorclass{12, "stainless plate armor"}
)

func (a *armorclass) String() string {
	return a.name
}

func (a *armorclass) PickUp(c *stats.Stats) {
}

func (a *armorclass) Drop(c *stats.Stats) {
}

func (a *armorclass) Sell(c *stats.Stats) {
}

func (a *armorclass) Wear(c *stats.Stats)    { c.Ac += a.class }
func (a *armorclass) TakeOff(c *stats.Stats) { c.Ac -= a.class }
