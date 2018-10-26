package items

import (
	termbox "github.com/nsf/termbox-go"
)

const (
	Pro = iota
	Mle
	Dex
	Sle
	Chm
	Ssp
	Web
	Str
	Enl
	Hel
	Cbl
	Cre
	Pha
	Inv
	Bal
	Cld
	Ply
	Can
	Has
	Ckl
	Vpr
	Dry
	Lit
	Drl
	Glo
	Flo
	Fgr
	Sca
	Hld
	Stp
	Tel
	Mfi
	Sph
	Gen
	Sum
	Wtw
	Alt
	Per
)

type Spell struct {
	Id   int
	Code string
	Name string
	Desc string
}

var spellLevel = []int{1, 4, 9, 14, 18, 22, 26, 29, 32, 35, 37, 37, 37, 37, 37}

var spellIndex = []Spell{
	{Pro, "pro", "protection", "generates a +2 protection field"},
	{Mle, "mle", "magic missile", "creates and hurls a magic missile equivalent to a +1 magic arrow"},
	{Dex, "dex", "dexterity", "adds +2 to the caster's dexterity"},
	{Sle, "sle", "sleep", "causes some monsters to go to sleep"},
	{Chm, "chm", "charm monster", "some monsters may be awed at your magnificence"},
	{Ssp, "ssp", "sonic spear", "causes your hands to emit a screeching sound toward what they point"},
	{Web, "web", "web", "causes strands of sticky thread to entangle an enemy"},
	{Str, "str", "strength", "adds +2 to the caster's strength for a short term"},
	{Enl, "enl", "englightenment", "the caster becomes aware of things in the vicinity"},
	{Hel, "hel", "healing", "restores some hp to the caster"},
	{Cbl, "cbl", "cure blindness", "restores sight to one so unfortunate as to be blinded"},
	{Cre, "cre", "create monster", "creates a monster near the caster appropriate for the location"},
	{Pha, "pha", "phantasmal forces", "creates illusions, and if believed, monsters die"},
	{Inv, "inv", "invisibility", "the caster becomes invisible"},
	{Bal, "bal", "fireball", "makes a ball of fire that burns on what it hits"},
	{Cld, "cld", "cold", "sends forth a cone of cold which freezes what it touches"},
	{Ply, "ply", "polymorph", "you can find out what this does for yourself"},
	{Can, "can", "cancellation", "negates the ability of a monster to use its special abilities"},
	{Has, "has", "haste self", "speeds up the caster's movements"},
	{Ckl, "ckl", "cloud kill", "creates a fog of poisonous gas which kills all that is within it"},
	{Vpr, "vpr", "vaporize rock", "this changes rock to air"},
	{Dry, "dry", "dehydration", "dries up water in the immediate vicinity"},
	{Lit, "lit", "lightning", "your finger will emit a lightning bolt when this spell is cast"},
	{Drl, "drl", "drain life", "subtracts hit points from both you and a monster"},
	{Glo, "glo", "invulnerability", "this globe helps to protect the player from physical attack"},
	{Flo, "flo", "flood", "this creates an avalanche of H2O to flood the immediate chamber"},
	{Fgr, "fgr", "finger of death", "this is a holy spell and calls upon your god to back you up"},
	{Sca, "sca", "scare monster", "terrifies the monster so that hopefully it won't hit the magic user"},
	{Hld, "hld", "hold monster", "the monster is frozen in its tracks if this is successful"},
	{Stp, "stp", "time stop", "all movement in the caverns ceases for a limited duration"},
	{Tel, "tel", "teleport away", "moves a particular monster around in the dungeon (hopefully away from you)"},
	{Mfi, "mfi", "magic fire", "this causes a curtain of fire to appear all around you"},
	{Sph, "sph", "sphere of annihilation", "anything caught in this sphere is instantly killed.  Warning -- dangerous"},
	{Gen, "gen", "genocide", "eliminates a species of monster from the game -- use sparingly"},
	{Sum, "sum", "summon demon", "summons a demon who hopefully helps you out"},
	{Wtw, "wtw", "walk through walls", "allows the player to walk through walls for a short period of time"},
	{Alt, "alt", "alter reality", "god only knows what this will do"},
	{Per, "per", "permanence", "makes a character spell permanent, i. e. protection, strength, etc."},
}

var spellLUT = map[string]Spell{
	"pro": spellIndex[Pro],
	"mle": spellIndex[Mle],
	"dex": spellIndex[Dex],
	"sle": spellIndex[Sle],
	"chm": spellIndex[Chm],
	"ssp": spellIndex[Ssp],
	"web": spellIndex[Web],
	"str": spellIndex[Str],
	"enl": spellIndex[Enl],
	"hel": spellIndex[Hel],
	"cbl": spellIndex[Cbl],
	"cre": spellIndex[Cre],
	"pha": spellIndex[Pha],
	"inv": spellIndex[Inv],
	"bal": spellIndex[Bal],
	"cld": spellIndex[Cld],
	"ply": spellIndex[Ply],
	"can": spellIndex[Can],
	"has": spellIndex[Has],
	"ckl": spellIndex[Ckl],
	"vpr": spellIndex[Vpr],
	"dry": spellIndex[Dry],
	"lit": spellIndex[Lit],
	"drl": spellIndex[Drl],
	"glo": spellIndex[Glo],
	"flo": spellIndex[Flo],
	"fgr": spellIndex[Fgr],
	"sca": spellIndex[Sca],
	"hld": spellIndex[Hld],
	"stp": spellIndex[Stp],
	"tel": spellIndex[Tel],
	"mfi": spellIndex[Mfi],
	"sph": spellIndex[Sph],
	"gen": spellIndex[Gen],
	"sum": spellIndex[Sum],
	"wtw": spellIndex[Wtw],
	"alt": spellIndex[Alt],
	"per": spellIndex[Per],
}

// SpellFromIndex returns the spell that corresponds to the index i
func SpellFromIndex(i int) Spell { return spellIndex[i] }

// SpellFromCode returns a spell from the 3 letter code
func SpellFromCode(c string) Spell { return spellLUT[c] }

// ProjectileSpell is a type of spell that cast a projectile
type ProjectileSpell struct {
	R rune
}

// Rune implements the io.Runeable interface
func (p *ProjectileSpell) Rune() rune {
	return p.R
}

// Fg implements the io.Runeable interface
func (p *ProjectileSpell) Fg() termbox.Attribute { return termbox.ColorDefault }

// Bg implements the io.Runeable interface
func (p *ProjectileSpell) Bg() termbox.Attribute { return termbox.ColorDefault }
