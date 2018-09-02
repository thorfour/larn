package state

func castSpell(s string) error {

	switch s {
	case "pro":
	case "mle":
	case "dex":
	case "sle":
	case "chm":
	case "ssp":
	case "web":
	case "str":
	case "enl":
	case "hel":
	case "cbl":
	case "cre":
	case "pha":
	case "inv":
	case "bal":
	case "cld":
	case "ply":
	case "can":
	case "has":
	case "ckl":
	case "vpr":
	case "dry":
	case "lit":
	case "drl":
	case "glo":
	case "flo":
	case "fgr":
	case "sca":
	case "hld":
	case "stp":
	case "tel":
	case "mfi":
	case "sph":
	case "gen":
	case "sum":
	case "wtw":
	case "alt":
	case "per":
	default:
		return ErrSpellNotKnown
	}
}
