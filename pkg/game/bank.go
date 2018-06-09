package game

import (
	"bytes"
	"fmt"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/golang/glog"
	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/game/state/items"
)

// account is the number of gold pieces the user has in the bank
var account int

func bankSplash() string {
	return `        Welcome to the First National Bank of Larn.`
}

func bankPage(gold int, stones []*items.Gem) string {
	pg := bankSplash() + "\n\n"
	var b []byte
	buf := bytes.NewBuffer(b)
	w := tabwriter.NewWriter(buf, 5, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "Gemstone\t\t\tAppraisal\t\tGemstone\t\t\tAppraisal")

	// Display all stones
	for i, s := range stones {
		switch i % 2 {
		case 0:
			fmt.Fprintf(w, "%s\t\t\t%v\t\t", s.String(), s.Value)
		case 1:
			fmt.Fprintf(w, "%s\t\t\t%v\t\t\n", s.String(), s.Value)
		}

	}

	// Pad out the rest
	for i := 0; i < 16-len(stones); i++ {
		fmt.Fprintln(w)
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
	g.renderSplash(bankPage(int(g.currentState.C.Stats.Gold), g.currentState.C.Gems()))
	return func(e termbox.Event) {
		switch e.Key {
		case termbox.KeyEsc: // Exit
			g.inputHandler = g.defaultHandler
			g.render(display(g.currentState))
		default:
			switch e.Ch {
			case 'd': // deposit into bank
				g.renderSplash(bankPage(int(g.currentState.C.Stats.Gold), g.currentState.C.Gems()) + howmuch())
				g.inputHandler = g.accountHandler(true)
				// TODO switch to a deposit handler
			case 'w': // witdraw from the bank
				g.renderSplash(bankPage(int(g.currentState.C.Stats.Gold), g.currentState.C.Gems()) + howmuch())
				g.inputHandler = g.accountHandler(false)
				// TODO switch to withdraw handler
			case 's': // sell a stone
				g.renderSplash(bankPage(int(g.currentState.C.Stats.Gold), g.currentState.C.Gems()))
				// TODO switch to sell handler
			}
		}
	}
}

func (g *Game) accountHandler(deposit bool) func(termbox.Event) {
	var amt string
	return func(e termbox.Event) {
		if e.Ch == '*' { // Short circuit for a deposit/withdraw all action
			if deposit {
				amt = fmt.Sprintf("%v", g.currentState.C.Stats.Gold)
			} else {
				amt = fmt.Sprintf("%v", account)
			}
			e.Key = termbox.KeyEnter // To enter the next switch statement to deposit/withdraw
		}
		switch e.Key {
		case termbox.KeyEsc: // Exit
			g.inputHandler = g.defaultHandler
			g.render(display(g.currentState))
		case termbox.KeyEnter: // Deposit/Withdraw
			n, err := strconv.Atoi(amt)
			if err != nil {
				glog.Errorf("unable to convert bank input to number: %s", amt)
			}
			if deposit {
				if g.currentState.C.Stats.Gold < uint(n) {
					g.renderSplash(bankPage(int(g.currentState.C.Stats.Gold), g.currentState.C.Gems()) + howmuch() + fmt.Sprintf(" %s\n", amt) + "  You don't have that much")
					time.Sleep(time.Millisecond * 700)
				} else {
					account += n
					g.currentState.C.Stats.Gold -= uint(n)
				}
			} else {
				if account < n {
					g.renderSplash(bankPage(int(g.currentState.C.Stats.Gold), g.currentState.C.Gems()) + howmuch() + fmt.Sprintf(" %s\n", amt) + "  You don't have that much in the bank!")
					time.Sleep(time.Millisecond * 700)
				} else {
					account -= n
					g.currentState.C.Stats.Gold += uint(n)
				}
			}
			g.inputHandler = g.bankHandler()
		default:
			switch e.Ch {
			case '0':
				fallthrough
			case '1':
				fallthrough
			case '2':
				fallthrough
			case '3':
				fallthrough
			case '4':
				fallthrough
			case '5':
				fallthrough
			case '6':
				fallthrough
			case '7':
				fallthrough
			case '8':
				fallthrough
			case '9':
				amt = amt + string(e.Ch)
				g.renderSplash(bankPage(int(g.currentState.C.Stats.Gold), g.currentState.C.Gems()) + howmuch() + fmt.Sprintf(" %s", amt))
			}
		}
	}
}
