package stats

type Stats struct {
	MaxSpells    uint   // max number of spells
	Spells       uint   // number of spells remaining
	Ac           uint   // armor class
	Wc           uint   // weapon class
	Level        uint   // current character level
	Exp          uint   // experience gained
	Title        string // current title based on level
	MaxHP        uint   // maximum health points
	Hp           uint   // remaining health points
	Str          uint   // strength
	Intelligence uint   // intelligence
	Wisdom       uint   // wisdom
	Con          uint   // constitution
	Dex          uint   // dexterity
	Cha          uint   // charisma
	Loc          string // current location
	Gold         uint   // current gold being held
}
