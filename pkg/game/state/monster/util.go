package monster

import "math/rand"

const (
	_ = iota
	Bat
	Gnome
	Hobgoblin
	Jackal
	Kobold
	Orc
	Snake
	Centipede
	Jaculi
	Troglodyte
	Ant
	Eye
	Leprechaun
	Nymph
	Quasit
	Rustmonster
	Zombie
	Assassinbug
	Bugbear
	Hellhound
	Icelizard
	Centaur
	Troll
	Yeti
	Whitedragon
	Elf
	Cube
	Metamorph
	Vortex
	Ziller
	Violetfungi
	Wraith
	Forvalaka
	Lamanobe
	Osequip
	Rothe
	Xorn
	Vampire
	Invisiblestalker
	Poltergeist
	Disenchantress
	Shamblingmound
	Yellowmold
	Umberhulk
	Gnomeking
	Mimic
	Waterlord
	Bronzedragon
	Greendragon
	Purpleworm
	Xvart
	Spiritnaga
	Silverdragon
	Platinumdragon
	Greenurchin
	Reddragon
	DemonlordI
	DemonlordII
	DemonlordIII
	DemonlordIV
	DemonlordV
	DemonlordVI
	DemonlordVII
	Demonprince = 64 // requires special spawn
)

// slice to generate monsters at a given level
var monstLevel = []int{5, 11, 17, 22, 27, 33, 39, 42, 46, 50, 53, 56, 59}

// FromLevel generates a monster for a dungeon level
func FromLevel(lev int) int {
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

	// don't generate demon lords or higher
	if tmp >= DemonlordI {
		tmp = 0
	}

	// Don't return a genocided monster
	for monsterData[tmp].Genocided != 0 {
		tmp++
		if tmp == DemonlordI {
			tmp = 0
		}
	}

	return tmp
}

type MonsterType struct {
	MonsterRune  rune   // the monsters displayable rune
	Name         string // the monsters displayable name
	Lvl          int
	Armor        int
	Dmg          int
	Attack       int
	Defense      int
	Genocided    int
	Intelligence int
	Gold         int
	Hitpoints    int
	Experience   int
}

var monsterData = map[int]MonsterType{
	Bat:              {'B', "bat", 1, 0, 1, 0, 0, 0, 3, 0, 1, 1},
	Gnome:            {'G', "gnome", 1, 10, 1, 0, 0, 0, 8, 30, 2, 2},
	Hobgoblin:        {'H', "hobgoblin", 1, 14, 2, 0, 0, 0, 5, 25, 3, 2},
	Jackal:           {'J', "jackal", 1, 17, 1, 0, 0, 0, 4, 0, 1, 1},
	Kobold:           {'K', "kobold", 1, 20, 1, 0, 0, 0, 7, 10, 1, 1},
	Orc:              {'O', "orc", 2, 12, 1, 0, 0, 0, 9, 40, 4, 2},
	Snake:            {'S', "snake", 2, 15, 1, 0, 0, 0, 3, 0, 3, 1},
	Centipede:        {'c', "giant centipede", 2, 14, 0, 4, 0, 0, 3, 0, 1, 2},
	Jaculi:           {'j', "jaculi", 2, 20, 1, 0, 0, 0, 3, 0, 2, 1},
	Troglodyte:       {'t', "troglodyte", 2, 10, 2, 0, 0, 0, 5, 80, 4, 3},
	Ant:              {'A', "giant ant", 2, 8, 1, 4, 0, 0, 4, 0, 5, 5},
	Eye:              {'E', "floating eye", 3, 8, 1, 0, 0, 0, 3, 0, 5, 2},
	Leprechaun:       {'L', "leprechaun", 3, 3, 0, 8, 0, 0, 3, 1500, 13, 45},
	Nymph:            {'N', "nymph", 3, 3, 0, 14, 0, 0, 9, 0, 18, 45},
	Quasit:           {'Q', "quasit", 3, 5, 3, 0, 0, 0, 3, 0, 10, 15},
	Rustmonster:      {'R', "rust monster", 3, 4, 0, 1, 0, 0, 3, 0, 18, 25},
	Zombie:           {'Z', "zombie", 3, 12, 2, 0, 0, 0, 3, 0, 6, 7},
	Assassinbug:      {'a', "assassin bug", 4, 9, 3, 0, 0, 0, 3, 0, 20, 15},
	Bugbear:          {'b', "bugbear", 4, 5, 4, 15, 0, 0, 5, 40, 20, 35},
	Hellhound:        {'h', "hell hound", 4, 5, 2, 2, 0, 0, 6, 0, 16, 35},
	Icelizard:        {'i', "ice lizard", 4, 11, 2, 10, 0, 0, 6, 50, 16, 25},
	Centaur:          {'C', "centaur", 4, 6, 4, 0, 0, 0, 10, 40, 24, 45},
	Troll:            {'T', "troll", 5, 4, 5, 0, 0, 0, 9, 80, 50, 300},
	Yeti:             {'Y', "yeti", 5, 6, 4, 0, 0, 0, 5, 50, 35, 100},
	Whitedragon:      {'d', "white dragon", 5, 2, 4, 5, 0, 0, 16, 500, 55, 1000},
	Elf:              {'e', "elf", 5, 8, 1, 0, 0, 0, 15, 50, 22, 35},
	Cube:             {'g', "gelatinous cube", 5, 9, 1, 0, 0, 0, 3, 0, 22, 45},
	Metamorph:        {'m', "metamorph", 6, 7, 3, 0, 0, 0, 3, 0, 30, 40},
	Vortex:           {'v', "vortex", 6, 4, 3, 0, 0, 0, 3, 0, 30, 55},
	Ziller:           {'z', "ziller", 6, 15, 3, 0, 0, 0, 3, 0, 30, 35},
	Violetfungi:      {'F', "violet fungi", 6, 12, 3, 0, 0, 0, 3, 0, 38, 100},
	Wraith:           {'W', "wraith", 6, 3, 1, 6, 0, 0, 3, 0, 30, 325},
	Forvalaka:        {'f', "forvalaka", 6, 2, 5, 0, 0, 0, 7, 0, 50, 280},
	Lamanobe:         {'l', "lama nobe", 7, 7, 3, 0, 0, 0, 6, 0, 35, 80},
	Osequip:          {'o', "osequip", 7, 4, 3, 16, 0, 0, 4, 0, 35, 100},
	Rothe:            {'r', "rothe", 7, 15, 5, 0, 0, 0, 3, 100, 50, 250},
	Xorn:             {'X', "xorn", 7, 0, 6, 0, 0, 0, 13, 0, 60, 300},
	Vampire:          {'V', "vampire", 7, 3, 4, 6, 0, 0, 17, 0, 50, 1000},
	Invisiblestalker: {' ', "invisible stalker", 7, 3, 6, 0, 0, 0, 5, 0, 50, 350},
	Poltergeist:      {'p', "poltergeist", 8, 1, 4, 0, 0, 0, 3, 0, 50, 450},
	Disenchantress:   {'q', "disenchantress", 8, 3, 0, 9, 0, 0, 3, 0, 50, 500},
	Shamblingmound:   {'s', "shambling mound", 8, 2, 5, 0, 0, 0, 6, 0, 45, 400},
	Yellowmold:       {'y', "yellow mold", 8, 12, 4, 0, 0, 0, 3, 0, 35, 250},
	Umberhulk:        {'U', "umber hulk", 8, 3, 7, 11, 0, 0, 14, 0, 65, 600},
	Gnomeking:        {'k', "gnome king", 9, -1, 10, 0, 0, 0, 18, 2000, 100, 3000},
	Mimic:            {'M', "mimic", 9, 5, 6, 0, 0, 0, 8, 0, 55, 99},
	Waterlord:        {'w', "water lord", 9, -10, 15, 7, 0, 0, 20, 0, 150, 15000},
	Bronzedragon:     {'D', "bronze dragon", 9, 2, 9, 3, 0, 0, 16, 300, 80, 4000},
	Greendragon:      {'D', "green dragon", 9, 3, 8, 10, 0, 0, 15, 200, 70, 2500},
	Purpleworm:       {'P', "purple worm", 9, -1, 11, 0, 0, 0, 3, 100, 120, 15000},
	Xvart:            {'x', "xvart", 9, -2, 12, 0, 0, 0, 13, 0, 90, 1000},
	Spiritnaga:       {'n', "spirit naga", 10, -20, 12, 12, 0, 0, 23, 0, 95, 20000},
	Silverdragon:     {'D', "silver dragon", 10, -1, 12, 3, 0, 0, 20, 700, 100, 10000},
	Platinumdragon:   {'D', "platinum dragon", 10, -5, 15, 13, 0, 0, 22, 1000, 130, 24000},
	Greenurchin:      {'u', "green urchin", 10, -3, 12, 0, 0, 0, 3, 0, 85, 5000},
	Reddragon:        {'D', "red dragon", 10, -2, 13, 3, 0, 0, 19, 800, 110, 14000},
	DemonlordI:       {' ', "type I demon lord", 12, -30, 18, 0, 0, 0, 20, 0, 140, 50000},
	DemonlordII:      {' ', "type II demon lord", 13, -30, 18, 0, 0, 0, 21, 0, 160, 75000},
	DemonlordIII:     {' ', "type III demon lord", 14, -30, 18, 0, 0, 0, 22, 0, 180, 100000},
	DemonlordIV:      {' ', "type IV demon lord", 15, -35, 20, 0, 0, 0, 23, 0, 200, 125000},
	DemonlordV:       {' ', "type V demon lord", 16, -40, 22, 0, 0, 0, 24, 0, 220, 150000},
	DemonlordVI:      {' ', "type VI demon lord", 17, -45, 24, 0, 0, 0, 25, 0, 240, 175000},
	DemonlordVII:     {' ', "type VII demon lord", 18, -70, 27, 6, 0, 0, 26, 0, 260, 200000},
	Demonprince:      {' ', "demon prince", 25, -127, 30, 6, 0, 0, 28, 0, 345, 300000},
}

// NameFromID returns the monsters name from a monster ID
func NameFromID(id int) string { return monsterData[id].Name }

// Genocided returns true if the monster has been genocided
func Genocided(id int) bool { return monsterData[id].Genocided == 1 }

// Random returns a random non-genocided monster ID
func Random() int {
	var id int
	for id = rand.Intn(Reddragon) + 1; Genocided(id); id = rand.Intn(Reddragon) + 1 {
	}

	return id
}
