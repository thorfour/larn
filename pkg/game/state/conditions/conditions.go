package conditions

type condition string

const (
	// Blindness means the player no longer reveals objects when encountering them
	Blindness condition = "blindness"
)

// ActiveConditions represents all active conditions a character might have
type ActiveConditions struct {
	active map[condition]func()
}

// EffectActive returns true if the given condition is active
func (a *ActiveConditions) EffectActive(c condition) bool {
	_, ok := a.active[c]
	return ok
}

// DecayAll calls the decay function on all active conditions
func (a *ActiveConditions) DecayAll(c condition) {
	for i := range a.active {
		a.active[i]()
	}
}
