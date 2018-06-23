package game

import (
	"bytes"
	"fmt"
	"text/tabwriter"
	"time"

	termbox "github.com/nsf/termbox-go"
)

// all courses cost 250
const courseCost = 250

type course struct {
	name      string
	available bool
	mobuls    uint
}

// used to order the map
var order = []string{"a", "b", "c", "d", "e", "f", "g", "h"}

var college = map[string]*course{
	"a": {"Fighters Training I", true, 10},
	"b": {"Fighters Training II", true, 15},
	"c": {"Introduction to Wizardry", true, 10},
	"d": {"Applied Wizardry", true, 20},
	"e": {"Behavioral Psychology", true, 10},
	"f": {"Faith for Today", true, 10},
	"g": {"Contemporary Dance", true, 10},
	"h": {"History of Larn", true, 5},
}

// takeCourse allows a user to take a course. Handles payment, mobuls, and college updates
func (g *Game) takeCourse(c string) error {
	if g.currentState.C.Stats.Gold <= courseCost {
		return fmt.Errorf(" You don't have enough gold to pay for that!")
	}

	crs := college[c]
	if !crs.available {
		return fmt.Errorf(" Sorry, but that class is filled.")
	}

	// Check for prerequisites
	switch {
	case c == "b" && college["a"].available:
		return fmt.Errorf(" Sorry, but this class has a prerequisite of Fighters Training I")
	case c == "d" && college["c"].available:
		return fmt.Errorf(" Sorry, but this class has a prerequisite of Introduction to Wizardry")
	}

	// Take the couse
	g.currentState.C.Stats.Gold -= courseCost
	college[c].available = false

	switch c {
	case "a":
		g.currentState.C.Stats.Str += 2
		g.currentState.C.Stats.Con++
	case "b":
		g.currentState.C.Stats.Str += 2
		g.currentState.C.Stats.Con += 2
	case "c":
		g.currentState.C.Stats.Intelligence += 2
	case "d":
		g.currentState.C.Stats.Intelligence += 2
	case "e":
		g.currentState.C.Stats.Cha += 3
	case "f":
		g.currentState.C.Stats.Wisdom += 2
	case "g":
		g.currentState.C.Stats.Dex += 3
	case "h":
		g.currentState.C.Stats.Intelligence++
	}

	// Use time for course
	g.currentState.UseTime(college[c].mobuls * 100)

	// Regen for the time used
	g.currentState.C.Stats.Hp = g.currentState.C.Stats.MaxHP
	g.currentState.C.Stats.Spells = g.currentState.C.Stats.MaxSpells
	g.currentState.Active["blind"] = nil
	g.currentState.Active["confuse"] = nil
	return nil
}

// diploma returns the message of what the characrter learned from taking a given course
func diploma(c string) string {
	switch c {
	case "a":
		return "You feel stronger!"
	case "b":
		return "You feel much stronger!"
	case "c":
		return "The task before you now seems more attainable!"
	case "d":
		return "The task before you now seems very attainable!"
	case "e":
		return "You now feel like a born leader!"
	case "f":
		return "You now feel more confident that you can find the potion in time!"
	case "g":
		return "You feel like dancing!"
	case "h":
		return "Your instructor told you that the Eye of Larn is rumored to be guarded\nby a platinum dragon who possesses psionic abilities."
	default:
		return ""
	}
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
				err := g.takeCourse(string(e.Ch))
				if err != nil {
					g.renderSplash(collegePage(int(g.currentState.C.Stats.Gold)) + "\n\n" + err.Error())
					time.Sleep(time.Second) // Blink the message
				} else {
					g.renderSplash(collegePage(int(g.currentState.C.Stats.Gold)) + "\n\n" + diploma(string(e.Ch)))
					time.Sleep(time.Second) // Blink the message
				}
				g.renderSplash(collegePage(int(g.currentState.C.Stats.Gold)))
			}
		}
	}
}
