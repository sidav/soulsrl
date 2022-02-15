package data

type Weapon struct {
	Code int
}

func (w *Weapon) GetData() *weaponData {
	return weaponsTable[w.Code]
}

type weaponData struct {
	Name               string
	dnum, dval, dmod   int
	AttackPatternCodes []int
}

const (
	WEAPON_BROKENSWORD = iota
	WEAPON_SHORTSWORD
	WEAPON_LONGSWORD
	WEAPON_SPEAR
)

var weaponsTable = map[int]*weaponData{
	WEAPON_BROKENSWORD: {
		Name: "Broken Sword",
		dnum: 1,
		dval: 3,
		dmod: 0,
		AttackPatternCodes: []int{
			APATTERN_SIMPLE_STRIKE,
			APATTERN_RIGHT_SLASH,
		},
	},
	WEAPON_SHORTSWORD: {
		Name: "Short Sword",
		dnum: 2,
		dval: 6,
		dmod: 0,
		AttackPatternCodes: []int{
			APATTERN_SIMPLE_STRIKE,
			APATTERN_RIGHT_SLASH,
		},
	},
	WEAPON_LONGSWORD: {
		Name: "Long Sword",
		dnum: 2,
		dval: 6,
		dmod: 2,
		AttackPatternCodes: []int{
			APATTERN_SIMPLE_STRIKE,
			APATTERN_SLASH,
		},
	},
	WEAPON_SPEAR: {
		Name: "Spear",
		dnum: 1,
		dval: 6,
		dmod: 3,
		AttackPatternCodes: []int{
			APATTERN_SIMPLE_STRIKE,
			APATTERN_LUNGE,
		},
	},
}
