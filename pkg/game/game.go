package game

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/game/data"
	"github.com/thorfour/larn/pkg/game/state"
	"github.com/thorfour/larn/pkg/game/state/character"
	"github.com/thorfour/larn/pkg/io"
)

const (
	internalKeyBufferSize = 10
	borderRune            = rune('#')
	borderWidth           = 67
	borderHeight          = 17
)

// Game holds all current game information
type Game struct {
	settings     data.Settings
	currentState *state.State

	// input channel from keyboard
	input chan termbox.Event

	// Indicates if the game has hit an error
	err error
}

// SaveFilePresent returns true if a save file exists, and the name of the file
func saveFilePresent() (bool, string) {
	// TODO
	return false, ""
}

// New initializes a game state
func New() *Game {
	g := new(Game)
	g.input = make(chan termbox.Event, internalKeyBufferSize)

	if ok, saveFile := saveFilePresent(); ok {
		// TODO handle loading a save game
		fmt.Println(saveFile)
		return g
	}

	// TODO setup game settings
	// g.settings

	// Generate starting game state
	g.currentState = state.New()

	return g
}

// Start is the entrypoint to running a new game, should not return without a request from the user
func (g *Game) Start() error {
	if err := termbox.Init(); err != nil {
		return fmt.Errorf("termbox failed to initialize: %v", err)
	}
	defer termbox.Close()

	// Start a listener for user input
	go io.KeyboardListener(g.input)

	// If the game wasn't from a save file, display the welcome screen
	if !g.settings.FromSaveFile {
		g.renderWelcome()

		// Wait for first key stroke to bypass welcome
		<-g.input
	}

	// Render the game
	g.render(display(g.currentState))

	// Game logic
	return g.run()
}

// run is the main game handler loop
func (g *Game) run() error {
	for {
		// Check for a game error
		if g.err != nil {
			return g.err
		}
		switch e := <-g.input; e.Ch {
		case 'H': // run left
			g.runAction(character.Left)
		case 'J': // run down
			g.runAction(character.Down)
		case 'K': // run up
			g.runAction(character.Up)
		case 'L': // run right
			g.runAction(character.Right)
		case 'Y': // run northwest
			g.runAction(character.UpLeft)
		case 'U': // run northeast
			g.runAction(character.UpRight)
		case 'B': // run southwest
			g.runAction(character.DownLeft)
		case 'N': // run southeast
			g.runAction(character.DownRight)
		case 'h': // move left
			g.currentState.Move(character.Left)
			g.render(display(g.currentState))
		case 'j': // move down
			g.currentState.Move(character.Down)
			g.render(display(g.currentState))
		case 'k': // move up
			g.currentState.Move(character.Up)
			g.render(display(g.currentState))
		case 'l': // move right
			g.currentState.Move(character.Right)
			g.render(display(g.currentState))
		case 'y': // move northwest
			g.currentState.Move(character.UpLeft)
			g.render(display(g.currentState))
		case 'u': // move northeast
			g.currentState.Move(character.UpRight)
			g.render(display(g.currentState))
		case 'b': // move southwest
			g.currentState.Move(character.DownLeft)
			g.render(display(g.currentState))
		case 'n': // move southeast
			g.currentState.Move(character.DownRight)
			g.render(display(g.currentState))
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
		case 'E': // Enter the building
			g.currentState.Enter()
			g.render(display(g.currentState))
		}
	}

	return nil
}

//  renderWelcome generates the welcome to larn message
func (g *Game) renderWelcome() {
	if g.err != nil {
		return
	}
	g.err = io.RenderWelcome(welcome)
}

func (g *Game) render(display [][]io.Runeable) {
	if g.err != nil {
		return
	}

	g.err = io.RenderNewGrid(display)
}

func (g *Game) renderCharacter(c character.Coordinate) {
	if g.err != nil {
		return
	}

	g.err = io.RenderCell(c.X, c.Y, '&', termbox.ColorGreen, termbox.ColorGreen)
}

func (g *Game) runAction(d character.Direction) {
	for moved := g.currentState.Move(d); moved; moved = g.currentState.Move(d) {
		g.render(display(g.currentState))
	}
}
