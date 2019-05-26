package conditions

// Condition is an effect that's applied to the character
type Condition int

const (
	// Blindness means the player no longer reveals objects when encountering them
	Blindness Condition = iota
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
	// ScareMonster makes them scared of you
	ScareMonster
	// WalkThroughWalls allows the player to move through walls
	WalkThroughWalls
	// Awareness allows the user to see invisible
	Awareness
	// UndeadProtection gives the user protection from the undead
	UndeadProtection
	// Dexterity increases user dexterity
	Dexterity
	// AltPro what does this do? TODO
	AltPro
	// Strength increases user strength
	Strength
	// Charm increases user charm
	Charm
	// AggravateMonsters aggravates monsters in the users area
	AggravateMonsters
	// Stealth makes the user not cause monsters to move
	Stealth
	// HasteMonsters increases monster speed within the area
	HasteMonsters
	// Globe of protection ? TODO
	Globe
	// SpiritProtection protection against spirit monsters
	SpiritProtection
	// Itching forces user to take off all armor
	Itching
	// Clumsiness TODO what does this do?
	Clumsiness
)

// ActiveConditions represents all active conditions a character might have
type ActiveConditions struct {
	active map[Condition]func(int)
}

// New returns a new active conditions struct
func New() *ActiveConditions {
	a := new(ActiveConditions)
	a.active = make(map[Condition]func(int))
	return a
}

// EffectActive returns true if the given condition is active
func (a *ActiveConditions) EffectActive(c Condition) bool {
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
func (a *ActiveConditions) Decay(c Condition) {
	a.active[c](0)
}

// MakePermanent adds a condition without a decay function
func (a *ActiveConditions) MakePermanent(c Condition) {
	a.active[c] = func(_ int) {}
}

// Refresh adds time onto a given condition, adds a new condition if the condition doesn't exist
func (a *ActiveConditions) Refresh(c Condition, n int, decay func()) {
	if _, ok := a.active[c]; ok {
		a.active[c](n)
		return
	}

	// Add condition
	a.Add(c, n, decay)
}

// Remove an active condition
func (a ActiveConditions) Remove(c Condition) {
	delete(a.active, c)
}

// Add an active condition with the given decay function
func (a ActiveConditions) Add(c Condition, dur int, decay func()) {
	a.active[c] = func(refresh int) {
		if refresh != 0 { // refresh instead of decay
			dur += refresh
			return
		}

		dur--
		if dur == 0 {
			if decay != nil {
				decay() // execute the decay func
			}
			a.Remove(c)
		}
	}
}
