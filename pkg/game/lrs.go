package game

import "fmt"

var taxes int

func tax() string {
	if taxes == 0 {
		return "  You do not owe us any taxes"
	}
	return fmt.Sprintf("  You owe us %v gold pieces", taxes)
}

func gold(gold uint) string {
	if gold == 0 {
		return "  You have no gold pieces"
	}

	return fmt.Sprintf("  You have %v gold pieces", gold)
}

func lrsPage(g uint) string {
	s := "\n\n\n\n  Welcome to the Larn Revenue Service distrcit office"
	s += "\n\n" + tax()
	s += "\n\n" + gold(g)
	s += "\n\n\n\n\n\n\n\n\n\n\n" + "  How can we help you? [(p) pay taxes, or escape]"

	return s
}
