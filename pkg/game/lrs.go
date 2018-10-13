package game

import (
	"fmt"
	"strconv"
	"time"

	termbox "github.com/nsf/termbox-go"
	log "github.com/sirupsen/logrus"
)

func tax(taxes int) string {
	if taxes == 0 {
		return "  You do not owe any taxes."
	}
	return fmt.Sprintf("  You presently owe %v gp in taxes.", taxes)
}

func gold(gold uint) string {
	if gold == 0 {
		return "  You have no gold pieces"
	}

	return fmt.Sprintf("  You have %v gold pieces", gold)
}

func lrsPage(taxes int, g uint) string {
	s := "\n\n\n\n  Welcome to the Larn Revenue Service distrcit office"
	s += "\n\n" + tax(taxes)
	s += "\n\n" + gold(g)
	s += "\n\n\n\n\n\n\n\n\n\n\n" + "  How can we help you? [(p) pay taxes, or escape]"

	return s
}

func (g *Game) lrsHandler() func(termbox.Event) {
	g.renderSplash(lrsPage(g.currentState.Taxes, g.currentState.C.Stats.Gold))
	return func(e termbox.Event) {
		switch e.Key {
		case termbox.KeyEsc: // Exit
			g.inputHandler = g.defaultHandler
			g.render(display(g.currentState))
		default:
			switch e.Ch {
			case 'p':
				g.renderSplash(lrsPage(g.currentState.Taxes, g.currentState.C.Stats.Gold) + "\n  How much? ")
				g.inputHandler = g.payTaxesHandler()
			}
		}
	}
}

func (g *Game) payTaxesHandler() func(termbox.Event) {
	var amt string
	return func(e termbox.Event) {
		switch e.Key {
		case termbox.KeyEsc: // Exit
			g.inputHandler = g.defaultHandler
			g.render(display(g.currentState))
		case termbox.KeyEnter: // Execute payment
			amount, err := strconv.Atoi(amt)
			if err != nil {
				log.WithField("amount", amt).Error("unable to convert tax input to number")
			}
			if g.currentState.C.Stats.Gold < uint(amount) {
				g.renderSplash(lrsPage(g.currentState.Taxes, g.currentState.C.Stats.Gold) + "\n  How much? " + amt + "\n  You don't have that much.")
				time.Sleep(time.Millisecond * 700) // blink the message
			} else { // pay the taxes
				g.currentState.C.Stats.Gold -= uint(amount)
				g.currentState.Taxes -= amount
			}
			g.inputHandler = g.lrsHandler()
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
				g.renderSplash(lrsPage(g.currentState.Taxes, g.currentState.C.Stats.Gold) + "\n  How much? " + amt)
			}
		}
	}
}
