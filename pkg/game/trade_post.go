package game

import (
	"bytes"
	"fmt"
	"text/tabwriter"
	"time"

	termbox "github.com/nsf/termbox-go"
)

// MaxDisplay the max number of inventory items to display at once
const MaxDisplay = 26

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
		if i >= MaxDisplay { // only display up to max
			break
		}
		if i%2 == 0 {
			fmt.Fprintf(w, "  %s\t\t\t\t", t)
		} else {
			fmt.Fprintf(w, "%s\n", t)
		}
	}
	w.Flush()

	s = s + string(buf.Bytes())

	// Pad out with empty lines
	for i := len(inv); i < MaxDisplay/2; i++ {
		s = s + "\n"
	}

	return s + "  What item do you want to sell us? [Press escape to leave]"
}

// tradingPostHandler input handler for the trading post
func (g *Game) tradingPostHandler() func(termbox.Event) {
	g.renderSplash(tradingPost(g.currentState.C.Inventory()))
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
				fallthrough
			case 'i':
				fallthrough
			case 'j':
				fallthrough
			case 'k':
				fallthrough
			case 'l':
				fallthrough
			case 'm':
				fallthrough
			case 'n':
				fallthrough
			case 'o':
				fallthrough
			case 'p':
				fallthrough
			case 'q':
				fallthrough
			case 'r':
				fallthrough
			case 's':
				fallthrough
			case 't':
				fallthrough
			case 'u':
				fallthrough
			case 'v':
				fallthrough
			case 'w':
				fallthrough
			case 'x':
				fallthrough
			case 'y':
				fallthrough
			case 'z':
				g.handleSellingInv(e.Ch)
			}
		}
	}
}

func (g *Game) handleSellingInv(r rune) {
	//  Check if they have the item
	if !g.currentState.C.HasItem(r) {
		g.renderSplash(tradingPost(g.currentState.C.Inventory()) + fmt.Sprintf("\n\n  You don't have item %s!", string(r)))
		time.Sleep(time.Millisecond * 700)
		g.renderSplash(tradingPost(g.currentState.C.Inventory()))
		return
	}
	//  check if the item is identified
	// TODO

	//  offer to buy the item
}
