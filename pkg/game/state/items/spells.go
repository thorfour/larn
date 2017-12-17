package items

type Spell struct {
	Level int
	Code  string
	Name  string
	Desc  string
}

var spellLevel = []int{1, 4, 9, 14, 18, 22, 26, 29, 32, 35, 37, 37, 37, 37, 37}

var spellIndex = []Spell{
	{0, "pro", "protection", "generates a +2 protection field"},
	{1, "mle", "magic missile", "creates and hurls a magic missile equivalent to a +1 magic arrow"},
	{2, "dex", "dexterity", "adds +2 to the caster's dexterity"},
	{3, "sle", "sleep", "causes some monsters to go to sleep"},
	{4, "chm", "charm monster", "some monsters may be awed at your magnificence"},
	{5, "ssp", "sonic spear", "causes your hands to emit a screeching sound toward what they point"},
	{6, "web", "web", "causes strands of sticky thread to entangle an enemy"},
	{7, "str", "strength", "adds +2 to the caster's strength for a short term"},
	{8, "enl", "englightenment", "the caster becomes aware of things in the vicinity"},
	{9, "hel", "healing", "restores some hp to the caster"},
	{10, "cbl", "cure blindness", "restores sight to one so unfortunate as to be blinded"},
	{11, "cre", "create monster", "creates a monster near the caster appropriate for the location"},
	{12, "pha", "phantasmal forces", "creates illusions, and if believed, monsters die"},
	{13, "inv", "invisibility", "the caster becomes invisible"},
	{14, "bal", "fireball", "makes a ball of fire that burns on what it hits"},
	{15, "cld", "cold", "sends forth a cone of cold which freezes what it touches"},
	{16, "ply", "polymorph", "you can find out what this does for yourself"},
	{17, "can", "cancellation", "negates the ability of a monster to use its special abilities"},
	{18, "has", "haste self", "speeds up the caster's movements"},
	{19, "ckl", "cloud kill", "creates a fog of poisonous gas which kills all that is within it"},
	{20, "vpr", "vaporize rock", "this changes rock to air"},
	{21, "dry", "dehydration", "dries up water in the immediate vicinity"},
	{22, "lit", "lightning", "your finger will emit a lightning bolt when this spell is cast"},
	{23, "drl", "drain life", "subtracts hit points from both you and a monster"},
	{24, "glo", "invulnerability", "this globe helps to protect the player from physical attack"},
	{25, "flo", "flood", "this creates an avalanche of H2O to flood the immediate chamber"},
	{26, "fgr", "finger of death", "this is a holy spell and calls upon your god to back you up"},
	{27, "sca", "scare monster", "terrifies the monster so that hopefully it won't hit the magic user"},
	{28, "hld", "hold monster", "the monster is frozen in its tracks if this is successful"},
	{29, "stp", "time stop", "all movement in the caverns ceases for a limited duration"},
	{30, "tel", "teleport away", "moves a particular monster around in the dungeon (hopefully away from you)"},
	{31, "mfi", "magic fire", "this causes a curtain of fire to appear all around you"},
	{32, "sph", "sphere of annihilation", "anything caught in this sphere is instantly killed.  Warning -- dangerous"},
	{33, "gen", "genocide", "eliminates a species of monster from the game -- use sparingly"},
	{34, "sum", "summon demon", "summons a demon who hopefully helps you out"},
	{35, "wtw", "walk through walls", "allows the player to walk through walls for a short period of time"},
	{36, "alt", "alter reality", "god only knows what this will do"},
	{37, "per", "permanence", "makes a character spell permanent, i. e. protection, strength, etc."},
}

var spellLUT = map[string]Spell{
	"pro": spellIndex[0],
	"mle": spellIndex[1],
	"dex": spellIndex[2],
	"sle": spellIndex[3],
	"chm": spellIndex[4],
	"ssp": spellIndex[5],
	"web": spellIndex[6],
	"str": spellIndex[7],
	"enl": spellIndex[8],
	"hel": spellIndex[9],
	"cbl": spellIndex[10],
	"cre": spellIndex[11],
	"pha": spellIndex[12],
	"inv": spellIndex[13],
	"bal": spellIndex[14],
	"cld": spellIndex[15],
	"ply": spellIndex[16],
	"can": spellIndex[17],
	"has": spellIndex[18],
	"ckl": spellIndex[19],
	"vpr": spellIndex[20],
	"dry": spellIndex[21],
	"lit": spellIndex[22],
	"drl": spellIndex[23],
	"glo": spellIndex[24],
	"flo": spellIndex[25],
	"fgr": spellIndex[26],
	"sca": spellIndex[27],
	"hld": spellIndex[28],
	"stp": spellIndex[29],
	"tel": spellIndex[30],
	"mfi": spellIndex[31],
	"sph": spellIndex[32],
	"gen": spellIndex[33],
	"sum": spellIndex[34],
	"wtw": spellIndex[35],
	"alt": spellIndex[36],
	"per": spellIndex[37],
}

// LearnSpell returns the spell that corresponds to the index i
func SpellFromIndex(i int) Spell { return spellIndex[i] }

// SpellFromCode returns a spell from the 3 letter code
func SpellFromCode(c string) Spell { return spellLUT[c] }
