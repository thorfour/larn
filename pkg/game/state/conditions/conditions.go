package conditions

import "github.com/golang/glog"

type condition int

const (
	// Blindness means the player no longer reveals objects when encountering them
	Blindness condition = iota
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

// Refresh adds time onto a given condition
func (a *ActiveConditions) Refresh(c condition, n int) {
	a.active[c](n)
}

// Remove an active condition
func (a ActiveConditions) Remove(c condition) {
	delete(a.active, c)
}

// Add an active condition with the given decay function
func (a ActiveConditions) Add(c condition, dur int, decay func()) {
	a.active[c] = func(refresh int) {
		glog.V(6).Infof("Decay %s: %v", c, dur)
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
