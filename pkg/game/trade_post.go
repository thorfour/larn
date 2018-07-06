package game

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	termbox "github.com/nsf/termbox-go"
)

func tradingPostWelcome() string {
	return `
  Welcome to the Larn Trading Post.  We buy items that explorers no longer find
  useful.  Since the condition of the items you bring in is not certain,
  and we incur great expense in reconditioning the items, we usually pay
  only 20% of their value were they to be new.  If the items are badly
  damaged, we will pay only 10% of their new value.
	  
  Here are the items we would be willing to buy from you:
`
}

func tradingPost(inv []string) string {
	s := tradingPostWelcome()
	buf := bytes.NewBuffer(make([]byte, 100))
	w := tabwriter.NewWriter(buf, 5, 0, 1, ' ', tabwriter.TabIndent)
	for i, t := range inv {
		if i%2 == 0 {
			fmt.Fprintf(w, "  %s\t\t\t\t", t)
		} else {
			fmt.Fprintf(w, "%s\n", t)
		}
	}
	w.Flush()

	return s + string(buf.Bytes())
}

// tradingPostHandler input handler for the trading post
func (g *Game) tradingPostHandler() func(termbox.Event) {
	g.renderSplash(tradingPost(g.currentState.C.Inventory()))
	return func(e termbox.Event) {
		switch e.Key {
		case termbox.KeyEsc: // Exit
			g.inputHandler = g.defaultHandler
			g.render(display(g.currentState))
		}
	}
}
