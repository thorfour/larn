package io

import termbox "github.com/nsf/termbox-go"

// KeyboardListener listens for keyboard input and passes the event to the channel. Does not return
func KeyboardListener(k chan termbox.Event) {
	termbox.SetInputMode(termbox.InputEsc) // Treat the Esc key as an Esc key

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey: // Send all keys to the event channel
			k <- ev
		case termbox.EventError: // Unexpected error occured
			panic(ev.Err)
		}

	}
}
