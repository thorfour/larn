package game

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/game/data"
	"github.com/thorfour/larn/pkg/game/state"
	"github.com/thorfour/larn/pkg/game/state/items"
	"github.com/thorfour/larn/pkg/game/state/maps"
	"github.com/thorfour/larn/pkg/game/state/types"
	"github.com/thorfour/larn/pkg/io"
)

type action int

const (
	dropAction action = iota
	wieldAction
	wearAction
	takeOffAction
	readAction
	quaffAction
)

var (
	Quit = fmt.Errorf("%s", "Quit")
	Save = fmt.Errorf("%s", "Save")
)

const (
	internalKeyBufferSize = 10
	borderRune            = rune('#')
	borderWidth           = 67
	borderHeight          = 17
	invMaxDisplay         = 5 // number of inventory items to display on a page at a time
)

// Game holds all current game information
type Game struct {
	settings     *data.Settings
	currentState *state.State

	// input channel from keyboard
	input chan termbox.Event

	// inputHandler is the function that handles input from the keyboard
	inputHandler func(e termbox.Event)

	// Indicates if the game has hit an error
	err error
}

// SaveFilePresent returns true if a save file exists, and the name of the file
func saveFilePresent() (bool, string) {
	// TODO
	return false, ""
}

// New initializes a game state
func New(s *data.Settings) *Game {
	glog.V(1).Infof("Creating new game with %v difficulty", s.Difficulty)
	g := new(Game)
	g.settings = s
	g.inputHandler = g.defaultHandler
	g.input = make(chan termbox.Event, internalKeyBufferSize)

	if ok, saveFile := saveFilePresent(); ok {
		// TODO handle loading a save game
		fmt.Println(saveFile)
		return g
	}

	// Generate starting game state
	g.currentState = state.New(g.settings.Difficulty)

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
		g.renderSplash(welcome)

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
			if g.err == Save || g.err == Quit { // Save or Quit aren't errors to return
				return nil
			}
			return g.err
		}

		// Check for player death
		if g.GameOver() {
			return nil // TODO show scores
		}

		// Get next input
		e := <-g.input

		// Handle the next input event
		g.inputHandler(e)
	}
}

func (g *Game) defaultWrapper() func(termbox.Event) {
	return g.defaultHandler
}

func (g *Game) defaultHandler(e termbox.Event) {

	switch e.Ch {
	case 'H': // run left
		g.runAction(types.Left)
	case 'J': // run down
		g.runAction(types.Down)
	case 'K': // run up
		g.runAction(types.Up)
	case 'L': // run right
		g.runAction(types.Right)
	case 'Y': // run northwest
		g.runAction(types.UpLeft)
	case 'U': // run northeast
		g.runAction(types.UpRight)
	case 'B': // run southwest
		g.runAction(types.DownLeft)
	case 'N': // run southeast
		g.runAction(types.DownRight)
	case 'h': // move left
		g.currentState.Move(types.Left)
		g.render(display(g.currentState))
	case 'j': // move down
		g.currentState.Move(types.Down)
		g.render(display(g.currentState))
	case 'k': // move up
		g.currentState.Move(types.Up)
		g.render(display(g.currentState))
	case 'l': // move right
		g.currentState.Move(types.Right)
		g.render(display(g.currentState))
	case 'y': // move northwest
		g.currentState.Move(types.UpLeft)
		g.render(display(g.currentState))
	case 'u': // move northeast
		g.currentState.Move(types.UpRight)
		g.render(display(g.currentState))
	case 'b': // move southwest
		g.currentState.Move(types.DownLeft)
		g.render(display(g.currentState))
	case 'n': // move southeast
		g.currentState.Move(types.DownRight)
		g.render(display(g.currentState))
	case ',': // Pick up the item
		g.currentState.PickUp()
		g.render(display(g.currentState))
	case '^': // identify a trap
		g.currentState.IdentTrap()
		g.render(display(g.currentState))
	case 'd': // drop an item
		g.inputHandler = g.itemAction(dropAction)
	case 'v': // print program version
	case '?': // help screen
		g.inputHandler = g.help()
	case 'g': // give present pack weight
	case 'i': // inventory your pockets
		g.inputHandler = g.inventoryWrapper(g.defaultWrapper)
	case 'A': // create diagnostic file
	case '.': // stay here
	case 'Z': // teleport yourself
	case 'c': // cast a spell
		g.inputHandler = g.cast()
	case 'r': // read a scroll/book
		g.inputHandler = g.itemAction(readAction)
	case 'q': // quaff a potion
		g.inputHandler = g.itemAction(quaffAction)
	case 'W': // wear armor
		g.inputHandler = g.itemAction(wearAction)
	case 'T': // take off armor
		g.inputHandler = g.itemAction(takeOffAction)
	case 'w': // wield a weapon
		g.inputHandler = g.itemAction(wieldAction)
	case 'P': // give tax status
	case 'D': // list all items found
	case 'e': // eat something
	case 'S': // save the game and quit
		g.err = Save
		return
	case 'Q': // quit the game
		g.err = Quit // Set the error to quit
		return
	case 'E': // Enter the building
		g.inputHandler = g.enterAction()
	}
}

//  renderSplash renders a pre-arranged splash screen
func (g *Game) renderSplash(s string) {
	if g.err != nil {
		return
	}
	g.err = io.RenderNew(s)
}

func (g *Game) render(display [][]io.Runeable) {
	if g.err != nil {
		return
	}

	g.err = io.RenderNewGrid(display)
}

func (g *Game) renderCharacter(c types.Coordinate) {
	if g.err != nil {
		return
	}

	g.err = io.RenderCell(c.X, c.Y, '&', termbox.ColorGreen, termbox.ColorGreen)
}

func (g *Game) runAction(d types.Direction) {
	for moved := g.currentState.Move(d); moved; moved = g.currentState.Move(d) {
		g.render(display(g.currentState))
	}
}

// inventoryWrapper returns a truncated input handler, used after a user requests an inventory display
// it will render the first inventory list, and subsequent calls the the function it returns will render the remaining pages
func (g *Game) inventoryWrapper(callback func() func(termbox.Event)) func(termbox.Event) {
	offset := 0
	s := g.currentState.Inventory()

	generateInv := func() []string {
		var inv []string
		inv = append(inv, "") // empty string at the top
		for i := 0; i < invMaxDisplay && offset < len(s); i++ {
			inv = append(inv, fmt.Sprintf("%v", s[offset]))
			offset++
		}
		inv = append(inv, g.currentState.TimeStr())             // add the elapsed time
		inv = append(inv, "   --- press space to continue ---") // add the help string at the bottom
		return inv
	}

	g.render(overlay(display(g.currentState), convert(generateInv())))

	return func(e termbox.Event) {
		switch e.Key {
		case termbox.KeyEsc: // Escape key
			g.inputHandler = callback()
			g.render(display(g.currentState))
		case termbox.KeySpace: // Space key
			if offset < len(s) { // Render next page
				g.render(overlay(display(g.currentState), convert(generateInv())))
				return
			}
			// No more pages to display, remove the overlay
			g.inputHandler = callback()
			g.render(display(g.currentState))
		default:
			glog.V(6).Infof("Receive invalid input: %s", string(e.Ch))
			return
		}
	}
}

// itemAction is a subroutine for a player to interact with his inventory
func (g *Game) itemAction(a action) func(termbox.Event) {
	glog.V(2).Infof("item action requested")

	switch a {
	case wieldAction:
		g.currentState.Log("What do you want to wield (- for nothing) [* for all] ?")
	case dropAction:
		g.currentState.Log("What do you want to drop [* for all] ?")
	case wearAction:
		g.currentState.Log("What do you want to wear [* for all] ?")
	case takeOffAction:
		if err := g.currentState.C.TakeOff(); err != nil {
			g.currentState.Log("You aren't wearing anything")
		} else {
			g.currentState.Log("Your armor is off")
		}
		g.render(display(g.currentState))
		return g.defaultHandler
	case readAction:
		g.currentState.Log("What do you want to read [* for all] ?")
	case quaffAction:
		g.currentState.Log("What do you want to quaff [space to view] ?")
	default:
		glog.Fatal("unknown item action %v", a)
	}

	g.render(display(g.currentState))

	// Capute the input character for the item action
	return func(e termbox.Event) {
		g.inputHandler = g.defaultHandler

		switch e.Key {
		case termbox.KeyEsc: // abort
			g.currentState.Log("aborted")
		default:
			if e.Ch == '*' {
				g.inputHandler = g.inventoryWrapper(g.itemActionWrapper(a))
				return
			}
			switch e.Ch {
			case '-': // drop nothing
			default: // try and act on something
				var err error
				switch a {
				case wieldAction:
					err = g.currentState.C.Wield(e.Ch)
				case wearAction:
					err = g.currentState.C.Wear(e.Ch)
				case dropAction:
					var item items.Item
					item, err = g.currentState.Drop(e.Ch)
					if err == nil {
						g.currentState.Log("You drop:")
						g.currentState.Log(fmt.Sprintf("%s) %s", string(e.Ch), item.String()))
					}
				case readAction:
					err = g.currentState.Read(e.Ch)
				case quaffAction:
					var callback func() bool
					callback, err = g.currentState.Quaff(e.Ch)
					g.render(display(g.currentState))
					if callback != nil {
						for callback() { // render until callback returns false
							g.render(display(g.currentState))
						}
					}
				}
				if err != nil {
					g.currentState.Log(err.Error())
				}
			}
		}
		g.render(display(g.currentState))
	}
}

// itemActionWrapper wraps the itemAction functon for inventory callbacks
func (g *Game) itemActionWrapper(a action) func() func(termbox.Event) {
	return func() func(termbox.Event) {
		return g.itemAction(a)
	}
}

func (g *Game) cast() func(termbox.Event) {
	if g.currentState.C.Stats.Spells <= 0 {
		g.currentState.Log("You don't have any spells!")
		g.render(display(g.currentState))
		return g.defaultHandler
	}

	g.currentState.Log("Enter your spell: ")
	g.render(display(g.currentState))

	var spell []byte

	// Next 3 inputs count towards casting a spell
	return func(e termbox.Event) {
		switch e.Key {
		case termbox.KeyEsc: // abort
			g.currentState.Log("aborted")
			g.inputHandler = g.defaultHandler
		default:
			spell = append(spell, byte(e.Ch))
			if len(spell) == 3 { // Spell complete
				glog.V(2).Infof("Spell: %s", string(spell))
				callback, err := g.currentState.Cast(string(spell))
				if err != nil {
					g.currentState.Log(err.Error())
					g.inputHandler = g.defaultHandler
					g.render(display(g.currentState))
					return
				}

				// If there was a callback func passed, that means the player is casting a projectile.
				// Obtian the direction the player would like to cast it, before using the callback to render
				// the animation
				if callback != nil {
					g.inputHandler = g.directionalSpellHandler(callback)
				}
			}
		}
		g.render(display(g.currentState))
	}
}

func (g *Game) help() func(termbox.Event) {

	// Display the first help screen
	i := 0
	g.renderSplash(help[i])
	i++

	return func(e termbox.Event) {
		switch e.Key {
		case termbox.KeyEnter: // exit
			fallthrough
		case termbox.KeyEsc: // abort
			g.inputHandler = g.defaultHandler
			g.render(display(g.currentState))
		case termbox.KeySpace:
			if i >= len(help) { // run out of help menus
				g.inputHandler = g.defaultHandler
				g.render(display(g.currentState))
				return
			}
			g.renderSplash(help[i])
			i++
		}
	}
}

// GameOver returns true if the game has ended
func (g *Game) GameOver() bool {
	if DEBUG {
		return false
	}
	return g.currentState.C.Stats.Hp == 0
}

// enterAction to handle a user entering a dungeon or store
func (g *Game) enterAction() func(termbox.Event) {

	switch g.currentState.Enter() {
	case maps.DndLvl:
		return g.dndStoreHandler()
	case maps.BankLvl:
		return g.bankHandler()
	case maps.LrsLvl:
		return g.lrsHandler()
	case maps.HomeLvl:
		return g.homeHandler()
	case maps.CollegeLvl:
		return g.collegeHandler()
	case maps.TradeLvl:
		return g.tradingPostHandler()
	default:
		g.render(display(g.currentState))
		return g.defaultHandler
	}
}

func (g *Game) directionalSpellHandler(cb func(types.Direction) bool) func(termbox.Event) {

	g.currentState.Log("What Direction? ")
	g.render(display(g.currentState))

	return func(e termbox.Event) {
		var d types.Direction
		switch e.Ch {
		case 'b':
			d = types.DownLeft
		case 'n':
			d = types.DownRight
		case 'y':
			d = types.UpLeft
		case 'u':
			d = types.UpRight
		case 'h':
			d = types.Left
		case 'k':
			d = types.Up
		case 'l':
			d = types.Right
		case 'j':
			d = types.Down
		default: // keep waiting for a valid direction to be entered
			return
		}

		// Continue to call the callback function until animation complete
		for cb(d) {
			g.render(display(g.currentState))
			time.Sleep(30 * time.Millisecond) // small sleep so user can see the spell animation
		}
		g.render(display(g.currentState))
		g.inputHandler = g.defaultHandler
	}
}
