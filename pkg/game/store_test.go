package game

import (
	"fmt"
	"testing"
)

var divider = `---------------------------------------------------------------`

func TestRender(t *testing.T) {
	fmt.Println(dndstorepage(0, 100))
	fmt.Println(divider)
	fmt.Println(bankPage(nil))
}
