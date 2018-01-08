package monster

import "math/rand"

// slice to generate monsters at a given level
var monstLevel = []int{5, 11, 17, 22, 27, 33, 39, 42, 46, 50, 53, 56, 59}

// MonsterFromLevel generates a monster for a dungeon level
func MonsterFromLevel(lev int) int {
	if lev < 1 {
		lev = 1
	}
	if lev > 12 {
		lev = 12
	}
	tmp := Waterlord
	if lev < 5 {
		for tmp == Waterlord { // use waterlord for sentinel since they can only spawn from fountains
			tmp = rand.Intn(monstLevel[lev-1]) + 1
		}
	} else {
		for tmp == Waterlord {
			tmp = rand.Intn(monstLevel[lev-1]-monstLevel[lev-4]) + monstLevel[lev-4] + 1
		}
	}

	// TODO don't return a genocided monster

	return tmp
}
