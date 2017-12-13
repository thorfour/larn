package maps

import (
	"testing"
)

// TestTreasureRooms ensures treasure rooms don't cause panic
func TestTreasureRooms(t *testing.T) {

	// Create a map
	m := newLevel(1)

	defer func() {
		if r := recover(); r != nil {
			t.Fatal("treasureRoom caused a panic")
		}
	}()

	// Generate treasure rooms
	for i := 0; i < 1000; i++ {
		treasureRoom(m)
	}
}
