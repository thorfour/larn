package game

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/game/state/items"
)

func bankSplash() string {
	return `        Welcome to the First National Bank of Larn.`
}

func bankPage(stones []items.Gem) string {
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
	fmt.Fprintf(w, "                       You have %v gold pieces in the bank\n", 0) // TODO get bank count
	fmt.Fprintf(w, "                       You have %v gold pieces\n", 0)             // TODO get gold count
	fmt.Fprintln(w)
	fmt.Fprintf(w, "  Your wish? [(d) deposit, (w) withdraw, (s) sell a stone, or escape]")

	w.Flush()
	return pg + string(buf.Bytes())
}

func (g *Game) bankHandler() func(termbox.Event) {
	g.renderSplash(bankPage(nil))
	return func(e termbox.Event) {
		switch e.Key {
		case termbox.KeyEsc: // Exit
			g.inputHandler = g.defaultHandler
			g.render(display(g.currentState))
		}
	}
}
