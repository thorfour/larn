package game

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/io"
)

const (
	internalKeyBufferSize = 10
)

// data holds all current game stat information
type Data struct {
	saveFile string // path to the save file for this game
	userID   uint64 // unique userID

	// input channel from keyboard
	input chan termbox.Key
}

// New initializes a game state
func New() *Data {
	d := new(Data)
	d.input = make(chan termbox.Key, internalKeyBufferSize)
	return d
}

// Start is the entrypoint to running a new game, should not return without a request from the user
func (d *Data) Start() error {
	if err := termbox.Init(); err != nil {
		return fmt.Errorf("termbox failed to initialize: %v", err)
	}
	defer termbox.Close()

	// Start a listener for user input
	go io.KeyboardListener(d.input)

	// Game logic
	return d.run()
}

// run is the main game handler loop
func (d *Data) run() error {
	for {
		select {
		case e := <-d.input:
			fmt.Println(e)
		}
	}

	return nil
}
