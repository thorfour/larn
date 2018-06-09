package game

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/game/state/items"
)

// account is the number of gold pieces the user has in the bank
var account int

func bankSplash() string {
	return `        Welcome to the First National Bank of Larn.`
}

func bankPage(gold int, stones []items.Gem) string {
	pg := bankSplash() + "\n\n"
	var b []byte
	buf := bytes.NewBuffer(b)
	w := tabwriter.NewWriter(buf, 5, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "Gemstone\t\t\tAppraisal\t\tGemstone\t\t\tAppraisal")

	// Display all stones
	for i := range stones { // TODO handle doing two gems on each line
		fmt.Fprintf(w, "%s\t\t\t%v\t\t", stones[i].String(), 0) // TODO get appraisal value
	}

	// Padd out the rest
	for i := 0; i < 16-len(stones); i++ {
		fmt.Fprintln(w, " \t\t\t \t\t \t\t\t ")
	}

	// Append helper lines at the bottom
	fmt.Fprintf(w, "                       You have %v gold pieces in the bank\n", account)
	fmt.Fprintf(w, "                       You have %v gold pieces\n", gold)
	fmt.Fprintln(w)
	fmt.Fprintf(w, "  Your wish? [(d) deposit, (w) withdraw, (s) sell a stone, or escape]\n")

	w.Flush()
	return pg + string(buf.Bytes())
}

// howmuch returns the how much string for deposits
func howmuch() string {
	return "  How much? [* for all]"
}

func (g *Game) bankHandler() func(termbox.Event) {
	g.renderSplash(bankPage(int(g.currentState.C.Stats.Gold), nil))
	return func(e termbox.Event) {
		switch e.Key {
		case termbox.KeyEsc: // Exit
			g.inputHandler = g.defaultHandler
			g.render(display(g.currentState))
		default:
			switch e.Ch {
			case 'd': // deposit into bank
				g.renderSplash(bankPage(int(g.currentState.C.Stats.Gold), nil) + howmuch())
				// TODO switch to a deposit handler
			case 'w': // witdraw from the bank
				g.renderSplash(bankPage(int(g.currentState.C.Stats.Gold), nil) + howmuch())
				// TODO switch to withdraw handler
			case 's': // sell a stone
				g.renderSplash(bankPage(int(g.currentState.C.Stats.Gold), nil))
				// TODO switch to sell handler
			}
		}
	}
}

func (g *Game) accountHandler() func(termbox.Event) {
	return nil
}
