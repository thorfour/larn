package game

// item is the generic item interface
type item interface {
	PickUp(c *character)
	Drop(c *character)
}

// food for edible items (fortune cookies)
type food interface {
	item
	Eat(c *character)
}

// potion for anything that can be quaffed
type potion interface {
	item
	Quaff(c *character)
}

// weapon for anything that can be wielded
type weapon interface {
	item
	Wield(c *character)
	Disarm(c *character)
}

// armor interface for anything that can be used as armor
type armorInterface interface {
	item
	Wear(c *character)
	TakeOff(c *character)
}
