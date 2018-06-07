package character

const (
	aa1 = " mighty evil master"
	aa2 = "apprentice demi-god"
	aa3 = "  minor demi-god   "
	aa4 = "  major demi-god   "
	aa5 = "    minor deity    "
	aa6 = "    major deity    "
	aa7 = "  novice guardian  "
	aa8 = "apprentice guardian"
	aa9 = "    The Creator    "
)

// titles[LEVEL-1] = title
var titles = []string{
	"  novice explorer  ",
	"apprentice explorer",
	" practiced explorer",
	"   expert explorer ",
	"  novice adventurer",
	"     adventurer    ",
	"apprentice conjurer",
	"     conjurer      ",
	"  master conjurer  ",
	"  apprentice mage  ",
	"        mage       ",
	"  experienced mage ",
	"     master mage   ",
	" apprentice warlord",
	"   novice warlord  ",
	"   expert warlord  ",
	"   master warlord  ",
	" apprentice gorgon ",
	"       gorgon      ",
	"  practiced gorgon ",
	"   master gorgon   ",
	"    demi-gorgon    ",
	"    evil master    ",
	" great evil master ",
	aa1, aa1, aa1, /* -27 */
	aa1, aa1, aa1, /* -30 */
	aa1, aa1, aa1, /* -33 */
	aa1, aa1, aa1, /* -36 */
	aa1, aa1, aa1, /* -39 */
	aa2, aa2, aa2, /* -42 */
	aa2, aa2, aa2, /* -45 */
	aa2, aa2, aa2, /* -48 */
	aa3, aa3, aa3, /* -51 */
	aa3, aa3, aa3, /* -54 */
	aa3, aa3, aa3, /* -57 */
	aa4, aa4, aa4, /* -60 */
	aa4, aa4, aa4, /* -63 */
	aa4, aa4, aa4, /* -66 */
	aa5, aa5, aa5, /* -69 */
	aa5, aa5, aa5, /* -72 */
	aa5, aa5, aa5, /* -75 */
	aa6, aa6, aa6, /* -78 */
	aa6, aa6, aa6, /* -81 */
	aa6, aa6, aa6, /* -84 */
	aa7, aa7, aa7, /* -87 */
	aa8, aa8, aa8, /* -90 */
	aa8, aa8, aa8, /* -93 */
	"  earth guardian   ", "   air guardian    ", "   fire guardian   ", /* -96 */
	"  water guardian   ", "  time guardian    ", " ethereal guardian ", /* -99 */
	aa9, aa9, aa9, /* -102 */
}

// MEG skill multiplier
const MEG = 1000000

var skill = []int{
	0, 10, 20, 40, 80, 160, 320, 640, 1280, 2560, 5120, /*  1-11 */
	10240, 20480, 40960, 100000, 200000, 400000, 700000, 1 * MEG, /* 12-19 */
	2 * MEG, 3 * MEG, 4 * MEG, 5 * MEG, 6 * MEG, 8 * MEG, 10 * MEG, /* 20-26 */
	12 * MEG, 14 * MEG, 16 * MEG, 18 * MEG, 20 * MEG, 22 * MEG, 24 * MEG, 26 * MEG, 28 * MEG, /* 27-35 */
	30 * MEG, 32 * MEG, 34 * MEG, 36 * MEG, 38 * MEG, 40 * MEG, 42 * MEG, 44 * MEG, 46 * MEG, /* 36-44 */
	48 * MEG, 50 * MEG, 52 * MEG, 54 * MEG, 56 * MEG, 58 * MEG, 60 * MEG, 62 * MEG, 64 * MEG, /* 45-53 */
	66 * MEG, 68 * MEG, 70 * MEG, 72 * MEG, 74 * MEG, 76 * MEG, 78 * MEG, 80 * MEG, 82 * MEG, /* 54-62 */
	84 * MEG, 86 * MEG, 88 * MEG, 90 * MEG, 92 * MEG, 94 * MEG, 96 * MEG, 98 * MEG, 100 * MEG, /* 63-71 */
	105 * MEG, 110 * MEG, 115 * MEG, 120 * MEG, 125 * MEG, 130 * MEG, 135 * MEG, 140 * MEG, /* 72-79 */
	145 * MEG, 150 * MEG, 155 * MEG, 160 * MEG, 165 * MEG, 170 * MEG, 175 * MEG, 180 * MEG, /* 80-87 */
	185 * MEG, 190 * MEG, 195 * MEG, 200 * MEG, 210 * MEG, 220 * MEG, 230 * MEG, 240 * MEG, /* 88-95 */
	250 * MEG, 260 * MEG, 270 * MEG, 280 * MEG, 290 * MEG, 300 * MEG, /* 96-101*/
}
