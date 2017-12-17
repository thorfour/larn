package items

import (
	"fmt"
	"math/rand"

	"github.com/thorfour/larn/pkg/game/state/stats"
)

const (
	bookRune = 'B'
)

type Book struct {
	Level uint // the level the book was placed in
	DefaultItem
	NoStats
}

// Rune implements the io.Runeable interface
func (b *Book) Rune() rune {
	if b.Visibility {
		return bookRune
	} else {
		return invisibleRune
	}
}

// Log implements the Loggable interface
func (b *Book) Log() string {
	return "You have found a book"
}

// String returns the texutal representation of the item
func (b *Book) String() string { return "a book" }

// Read implements the Readable interface
func (b *Book) Read(s *stats.Stats) []string {
	// Generate a spell based on level
	var i int
	if b.Level < 4 {
		i = rand.Intn(spellLevel[b.Level])
	} else {
		i = rand.Intn(spellLevel[b.Level]-9) + 9
	}

	spell := SpellFromIndex(i)

	// Mark the spell as known
	s.KnownSpells[spell.Code] = true

	logs := []string{"", fmt.Sprintf("Spell %s: %s", spell.Code, spell.Name), spell.Desc}

	// Reading can gain player knowledge
	if rand.Intn(10) == 0 {
		s.Intelligence++
		logs = append(logs, "Your int went up by one!")
	}

	return logs
}
