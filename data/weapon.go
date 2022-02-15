package data

type weapon struct {
	code int
}

type weaponData struct {
	name               string
	dnum, dval, dmod   int
	attackPatternCodes []int
}

var weaponsTable = []*weaponData{
	{
		name: "Broken Sword",
		dnum: 1,
		dval: 3,
		dmod: 0,
		attackPatternCodes: []int{
			APATTERN_SIMPLE_STRIKE,
			APATTERN_RIGHT_SLASH,
		},
	},
	{
		name: "Short Sword",
		dnum: 2,
		dval: 6,
		dmod: 0,
		attackPatternCodes: []int{
			APATTERN_SIMPLE_STRIKE,
			APATTERN_RIGHT_SLASH,
		},
	},
	{
		name: "Long Sword",
		dnum: 2,
		dval: 6,
		dmod: 2,
		attackPatternCodes: []int{
			APATTERN_SIMPLE_STRIKE,
			APATTERN_SLASH,
		},
	},
	{
		name: "Spear",
		dnum: 1,
		dval: 6,
		dmod: 3,
		attackPatternCodes: []int{
			APATTERN_SIMPLE_STRIKE,
			APATTERN_LUNGE,
		},
	},
}
