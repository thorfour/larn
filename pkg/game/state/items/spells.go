package items

type spell struct {
	code string
	name string
	desc string
}

var spellLevel = []int{1, 4, 9, 14, 18, 22, 26, 29, 32, 35, 37, 37, 37, 37, 37}

var spellIndex = []spell{
	{"pro", "protection", "generates a +2 protection field"},
	{"mle", "magic missile", "creates and hurls a magic missile equivalent to a +1 magic arrow"},
	{"dex", "dexterity", "adds +2 to the caster's dexterity"},
	{"sle", "sleep", "causes some monsters to go to sleep"},
	{"chm", "charm monster", "some monsters may be awed at your magnificence"},
	{"ssp", "sonic spear", "causes your hands to emit a screeching sound toward what they point"},
	{"web", "web", "causes strands of sticky thread to entangle an enemy"},
	{"str", "strength", "adds +2 to the caster's strength for a short term"},
	{"enl", "englightenment", "the caster becomes aware of things in the vicinity"},
	{"hel", "healing", "restores some hp to the caster"},
	{"cbl", "cure blindness", "restores sight to one so unfortunate as to be blinded"},
	{"cre", "create monster", "creates a monster near the caster appropriate for the location"},
	{"pha", "phantasmal forces", "creates illusions, and if believed, monsters die"},
	{"inv", "invisibility", "the caster becomes invisible"},
	{"bal", "fireball", "makes a ball of fire that burns on what it hits"},
	{"cld", "cold", "sends forth a cone of cold which freezes what it touches"},
	{"ply", "polymorph", "you can find out what this does for yourself"},
	{"can", "cancellation", "negates the ability of a monster to use its special abilities"},
	{"has", "haste self", "speeds up the caster's movements"},
	{"ckl", "cloud kill", "creates a fog of poisonous gas which kills all that is within it"},
	{"vpr", "vaporize rock", "this changes rock to air"},
	{"dry", "dehydration", "dries up water in the immediate vicinity"},
	{"lit", "lightning", "your finger will emit a lightning bolt when this spell is cast"},
	{"drl", "drain life", "subtracts hit points from both you and a monster"},
	{"glo", "invulnerability", "this globe helps to protect the player from physical attack"},
	{"flo", "flood", "this creates an avalanche of H2O to flood the immediate chamber"},
	{"fgr", "finger of death", "this is a holy spell and calls upon your god to back you up"},
	{"sca", "scare monster", "terrifies the monster so that hopefully it won't hit the magic user"},
	{"hld", "hold monster", "the monster is frozen in its tracks if this is successful"},
	{"stp", "time stop", "all movement in the caverns ceases for a limited duration"},
	{"tel", "teleport away", "moves a particular monster around in the dungeon (hopefully away from you)"},
	{"mfi", "magic fire", "this causes a curtain of fire to appear all around you"},
	{"sph", "sphere of annihilation", "anything caught in this sphere is instantly killed.  Warning -- dangerous"},
	{"gen", "genocide", "eliminates a species of monster from the game -- use sparingly"},
	{"sum", "summon demon", "summons a demon who hopefully helps you out"},
	{"wtw", "walk through walls", "allows the player to walk through walls for a short period of time"},
	{"alt", "alter reality", "god only knows what this will do"},
	{"per", "permanence", "makes a character spell permanent, i. e. protection, strength, etc."},
}

var spellLUT = map[string]spell{
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
func SpellFromIndex(i int) spell { return spellIndex[i] }
