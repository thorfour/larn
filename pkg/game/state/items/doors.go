package items

const (
	doorClosedRune = 'D'
	doorOpenRune   = 'O'
)

// Door are doors to treasure rooms
type Door struct {
	Open bool // indicates if the door is open or closed
	DefaultItem
}

// Rune implements the io.Runeable interface
func (d *Door) Rune() rune {
	if d.Visibility {
		if d.Open {
			return doorOpenRune
		}
		return doorClosedRune
	}
	return invisibleRune
}

// Log implementes the Loggable interface
func (d *Door) Log() string {
	if d.Open {
		return "There is an open door"
	}
	return "There is a closed door"
}
