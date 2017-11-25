package game

// armor satisfies the item interface as well as the armorInterface
type armor struct {
	class uint
	name  string
}

var (
	shield = armor{2, "shield"}

	leather     = armor{2, "leather armor"}
	studleather = armor{3, "studded leather armor"}
	ring        = armor{5, "ring mai"}
	chain       = armor{6, "chain mail"}
	splint      = armor{7, "splint mail"}
	plateMail   = armor{9, "plate mail"}
	plateArmor  = armor{10, "plate armor"}
	stainless   = armor{12, "stainless plate armor"}
)

func (a *armor) String() string {
	return a.name
}

func (a *armor) PickUp(c *character) {
}

func (a *armor) Drop(c *character) {
}

func (a *armor) Sell(c *character) {
}

func (a *armor) Wear(c *character) {
	c.ac += a.class
}

func (a *armor) TakeOff(c *character) {
	c.ac -= a.class
}
