package game

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/io"
)

const (
	internalKeyBufferSize = 10
	borderRune            = rune('#')
	borderWidth           = 67
	borderHeight          = 17
)

// data holds all current game stat information
type Data struct {
	saveFile   string // path to the save file for this game
	userID     uint64 // unique userID
	difficulty uint32 // current game difficulty

	// input channel from keyboard
	input chan termbox.Event
}

// New initializes a game state
func New() *Data {
	d := new(Data)
	d.input = make(chan termbox.Event, internalKeyBufferSize)
	return d
}

// Start is the entrypoint to running a new game, should not return without a request from the user
func (d *Data) Start() error {
	if err := termbox.Init(); err != nil {
		return fmt.Errorf("termbox failed to initialize: %v", err)
	}
	defer termbox.Close()

	io.RenderWelcome(welcome)

	// Start a listener for user input
	go io.KeyboardListener(d.input)

	// Wait for first key stroke to bypass welcome
	<-d.input

	// Game logic
	return d.run()
}

// run is the main game handler loop
func (d *Data) run() error {
	for {
		switch e := <-d.input; e.Ch {
		case 'H': // run left
		case 'J': // run down
		case 'K': // run up
		case 'L': // run right
		case 'Y': // run northwest
		case 'U': // run northeast
		case 'B': // run southwest
		case 'N': // run southeast
		case 'h': // move left
		case 'j': // move down
		case 'k': // move up
		case 'l': // move right
		case 'y': // move northwest
		case 'u': // move northeast
		case 'b': // move southwest
		case 'n': // move southeast
		case '^': // identify a trap
		case 'd': // drop an item
		case 'v': // print program version
		case '?': // help screen
		case 'g': // give present pack weight
		case 'i': // inventory your pockets
		case 'A': // create diagnostic file
		case '.': // stay here
		case 'Z': // teleport yourself
		case 'c': // cast a spell
		case 'r': // read a scroll
		case 'q': // quaff a potion
		case 'W': // wear armor
		case 'T': // take off armor
		case 'w': // wield a weapon
		case 'P': // give tax status
		case 'D': // list all items found
		case 'e': // eat something
		case 'S': // save the game and quit
			fallthrough
		case 'Q': // quit the game
			return nil
		}
	}

	return nil
}
