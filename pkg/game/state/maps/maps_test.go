package maps

import "testing"

func TestMapsSmoke(t *testing.T) {
	m := New()

	for i, row := range m.home {
		for j, r := range row {
			if ru := r.Rune(); ru != '.' {
				t.Error("(%v,%v): not an empty rune: %v", i, j, ru)
			}
		}
	}
}
