package data

type Armor struct {
	code int
}

func (a *Armor) GetData() *weaponData {
	return weaponsTable[a.code]
}

type armorData struct {
	Name            string
	ArmorClass      int
	DamageReduction int
}

var armorsTable = []*armorData{
	{
		Name:            "Clothes",
		ArmorClass:      1,
		DamageReduction: 1,
	},
	{
		Name:            "Leather armor",
		ArmorClass:      3,
		DamageReduction: 3,
	},
	{
		Name:            "Plate mail",
		ArmorClass:      11,
		DamageReduction: 6,
	},
}
