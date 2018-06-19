package game

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
)

func homePage(name string, time int) string {
	s := "\n\n\n\n  Welcome home " + name + ". Latest word from the doctor is not good.\n"

	if time <= 0 {
		s += "\n  The doctor has the sad duty to inform you that your daughter died!\n"
		s += "  You didn't make it in time.  There was nothing he could do without the potion.\n"
	}

	s += "\n  The diagnosis is confirmed as dianthroritis.  He guesses that\n"
	s += fmt.Sprintf("  your daughter has only %d mobuls left in this world.  It's up to you,\n", time)
	s += "  " + name + "to find the only hope for your daughter, the very rare\n"
	s += "  potion of cure dianthroritis.  It is rumored that only deep in the\n"
	s += "  depths of the caves can this potion be found.\n\n\n"
	s += "  ----- Press escape to leave -----"

	return s
}

func (g *Game) homeHandler() func(termbox.Event) {
	g.renderSplash(homePage(g.currentState.Name, g.currentState.TimeLeft()))
	if g.currentState.TimeLeft() <= 0 {
		// TODO handle game over
	}
	return func(e termbox.Event) {
		switch e.Key {
		case termbox.KeyEsc: // Exit
			g.inputHandler = g.defaultHandler
			g.render(display(g.currentState))
		}
	}
}
