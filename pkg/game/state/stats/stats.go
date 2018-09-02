package stats

type Stats struct {
	MaxSpells    uint            // max number of spells
	Spells       uint            // number of spells remaining
	Ac           int             // armor class
	Wc           int             // weapon class
	Level        uint            // current character level
	Exp          uint            // experience gained
	Title        string          // current title based on level
	MaxHP        uint            // maximum health points
	Hp           uint            // remaining health points
	Str          uint            // strength
	StrExtra     int             // extra strength
	MoreDmg      int             // more damage
	Intelligence uint            // intelligence
	Wisdom       uint            // wisdom
	Con          uint            // constitution
	Dex          uint            // dexterity
	Cha          uint            // charisma
	Loc          string          // current location
	Gold         uint            // current gold being held
	Special      map[int]bool    // Special stats for if the character is holding special items
	KnownSpells  map[string]bool // Known spells
}

// RaiseMaxHP raises the max HP and the HP by n
func (s *Stats) RaiseMaxHP(n uint) {
	s.MaxHP += n
	s.Hp += n
}

// GainHP adds HP, handles not exceeding max HP
func (s *Stats) GainHP(n uint) {
	if s.Hp+n < s.MaxHP {
		s.Hp += n
	} else {
		s.Hp = s.MaxHP
	}
}
