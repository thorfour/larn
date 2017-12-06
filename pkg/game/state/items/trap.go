package items

const (
	TeleTrap = iota
	ArrowTrap
	DartTrap
	DoorTrap
)

const (
	teleTrapRune = ' '
	trapRune     = '^'
)

type Trap struct {
	TrapType   int
	discovered bool
	DefaultItem
}

// Rune implements the io.Runeable interface
func (t *Trap) Rune() rune {
	if t.discovered {
		return invisibleRune
	}

	switch t.TrapType {
	case TeleTrap:
		return teleTrapRune
	default:
		return trapRune
	}
}

func (t *Trap) Log() string {
	switch t.TrapType {
	case TeleTrap:
		return "Zaaaappp! You've been teleported"
	case ArrowTrap:
		return "You are hit by an arrow"
	case DartTrap:
		return "You are hit by a dart"
	default:
		return ""
	}
}
