package game

type stats struct {
	maxSpells    uint   // max number of spells
	spells       uint   // number of spells remaining
	ac           uint   // armor class
	wc           uint   // weapon class
	level        uint   // current character level
	exp          uint   // experience gained
	title        string // current title based on level
	maxHP        uint   // maximum health points
	hp           uint   // remaining health points
	str          uint   // strength
	intelligence uint   // intelligence
	wisdom       uint   // wisdom
	con          uint   // constitution
	dex          uint   // dexterity
	cha          uint   // charisma
	loc          string // current location
	gold         uint   // current gold being held
}

type character struct {
	armor  []item // Currently worn armor
	weapon []item // Currently wielded weapon(s)
	stats
}
