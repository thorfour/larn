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
