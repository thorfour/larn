package game

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	termbox "github.com/nsf/termbox-go"
)

type course struct {
	name      string
	available bool
	mobuls    int
}

// used to order the map
var order = []string{"a", "b", "c", "d", "e", "f", "g", "h"}

var college = map[string]course{
	"a": {"Fighters Training I", true, 10},
	"b": {"Fighters Training II", true, 15},
	"c": {"Introduction to Wizardry", true, 10},
	"d": {"Applied Wizardry", true, 20},
	"e": {"Behavioral Psychology", true, 10},
	"f": {"Faith for Today", true, 10},
	"g": {"Contemporary Dance", true, 10},
	"h": {"History of Larn", true, 5},
}

func collegePage(gold int) string {
	s := "\n The College of Larn offers the exciting opportunity of higher education to\n"
	s += " all inhabitants of the caves.  Here is a list of the class schedule:\n\n\n"
	buf := bytes.NewBuffer(make([]byte, 100))
	w := tabwriter.NewWriter(buf, 5, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "                 Course Name\t\t\t\tTime Needed\n\n")
	for _, i := range order {
		c := college[i]
		if c.available {
			fmt.Fprintf(w, "            %s)  %s\t\t%v mobuls", i, c.name, c.mobuls)
		} else {
			fmt.Fprintf(w, "\n")
		}
		fmt.Fprintf(w, "\n")
	}

	w.Flush()
	costLine := "\n            All courses cost 250 gold pieces"
	goldLine := fmt.Sprintf("\n\n                       You are presently carrying %v gold pieces", gold)
	choiceLine := "\n\n What is your choice? [Press escape to leave]"
	return s + string(buf.Bytes()) + costLine + goldLine + choiceLine
}

// collegeHandler displays the college of larn
func (g *Game) collegeHandler() func(termbox.Event) {
	g.renderSplash(collegePage(int(g.currentState.C.Stats.Gold)))
	return func(e termbox.Event) {
		switch e.Key {
		case termbox.KeyEsc: // Exit
			g.inputHandler = g.defaultHandler
			g.render(display(g.currentState))
		default:
			switch e.Ch {
			case 'a':
				fallthrough
			case 'b':
				fallthrough
			case 'c':
				fallthrough
			case 'd':
				fallthrough
			case 'e':
				fallthrough
			case 'f':
				fallthrough
			case 'g':
				fallthrough
			case 'h':
				// Attempt to take a course

			}
		}
	}
}
