package conditions

type condition int

const (
	// Blindness means the player no longer reveals objects when encountering them
	Blindness condition = iota
	// Confusion character is confuesd
	Confusion
	// Heroic status
	Heroic
	// GiantStrength spell
	GiantStrength
	// FireResistance resistance to hell hounds and other fire sources
	FireResistance
	// HalfDamage player only deals half damange
	HalfDamage
	// SeeInvisible allows player to see invisible stalkers
	SeeInvisible
	// HoldMonsters keeps all monsters from moving unless provoked
	HoldMonsters
	// TimeStop means no time passes
	TimeStop
	// GlobeOfInvul globe of invulnerability
	GlobeOfInvul
	// SpellOfStrength gives the player strength
	SpellOfStrength
	// SpellOfDexterity gives the player dexterity
	SpellOfDexterity
	// SpellOfProtection gives the player armor
	SpellOfProtection
	// Invisiblity means monsters can't see the player
	Invisiblity
	// CharmMonsters monsters are more likely to be charmed
	CharmMonsters
	// Cancellation TODO
	Cancellation
	// HasteSelf increases character speed
	HasteSelf
)

// ActiveConditions represents all active conditions a character might have
type ActiveConditions struct {
	active map[condition]func(int)
}

// New returns a new active conditions struct
func New() *ActiveConditions {
	a := new(ActiveConditions)
	a.active = make(map[condition]func(int))
	return a
}

// EffectActive returns true if the given condition is active
func (a *ActiveConditions) EffectActive(c condition) bool {
	_, ok := a.active[c]
	return ok
}

// DecayAll calls the decay function on all active conditions
func (a *ActiveConditions) DecayAll() {
	for i := range a.active {
		a.active[i](0)
	}
}

// Decay calls the decay function on a single condition
func (a *ActiveConditions) Decay(c condition) {
	a.active[c](0)
}

// Refresh adds time onto a given condition, adds a new condition if the condition doesn't exist
func (a *ActiveConditions) Refresh(c condition, n int, decay func()) {
	if _, ok := a.active[c]; ok {
		a.active[c](n)
		return
	}

	// Add condition
	a.Add(c, n, decay)
}

// Remove an active condition
func (a ActiveConditions) Remove(c condition) {
	delete(a.active, c)
}

// Add an active condition with the given decay function
func (a ActiveConditions) Add(c condition, dur int, decay func()) {
	a.active[c] = func(refresh int) {
		if refresh != 0 { // refresh instead of decay
			dur += refresh
			return
		}

		dur--
		if dur == 0 {
			if decay != nil {
				decay() // execure the decay func
			}
			a.Remove(c)
		}
	}
}
