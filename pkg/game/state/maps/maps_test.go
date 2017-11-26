package maps

import "testing"

func TestMapsSmoke(t *testing.T) {
	m := New()

	for i, col := range m.home {
		for j, r := range col {
			if ru := r.(Empty).Rune(); ru != '.' {
				t.Error("(%v,%v): not an empty rune: %v", i, j, ru)
			}
		}
	}
}
