package data

type weapon struct {
	code int
}

type weaponData struct {
	name             string
	dnum, dval, dmod int
}

var weaponsTable = []*weaponData{
	{
		name: "Broken Sword",
		dnum: 1,
		dval: 3,
		dmod: 0,
	},
	{
		name: "Short Sword",
		dnum: 2,
		dval: 6,
		dmod: 0,
	},
}
