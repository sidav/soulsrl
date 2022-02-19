package data

type Armor struct {
	Code int
}

func (a *Armor) GetData() *armorData {
	return armorsTable[a.Code]
}

type armorData struct {
	Name            string
	ArmorClass      int
	DamageReduction int
}

const (
	ARMOR_BASIC = iota
	ARMOR_LEATHER
	ARMOR_PLATE
	ARMOR_HIGH_AC
)

var armorsTable = map[int]*armorData{
	ARMOR_BASIC: {
		Name:            "Clothes",
		ArmorClass:      1,
		DamageReduction: 1,
	},
	ARMOR_LEATHER: {
		Name:            "Leather armor",
		ArmorClass:      3,
		DamageReduction: 3,
	},
	ARMOR_PLATE: {
		Name:            "Plate mail",
		ArmorClass:      4,
		DamageReduction: 11,
	},
	ARMOR_HIGH_AC: {
		Name:            "Duelist armor",
		ArmorClass:      8,
		DamageReduction: 2,
	},
}
