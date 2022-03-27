package data

import "github.com/sidav/sidavgorandom/fibrandom"

type Armor struct {
	Code int
}

func (a *Armor) GetData() *armorData {
	return armorsTable[a.Code]
}

type armorData struct {
	Name               string
	ArmorClass         int
	BaseToBlockPercent int
}

func (a *Armor) RollBlock(rnd *fibrandom.FibRandom) int {
	data := a.GetData()
	totalBlock := 0
	for i := 0; i < data.ArmorClass; i++ {
		rand := rnd.RandInRange(1, 100)
		if rand < data.BaseToBlockPercent {
			totalBlock++
		}
	}
	return totalBlock
}

const (
	ARMOR_BASIC = iota
	ARMOR_LEATHER
	ARMOR_PLATE
	ARMOR_HIGH_AC
)

var armorsTable = map[int]*armorData{
	ARMOR_BASIC: {
		Name:               "Clothes",
		ArmorClass:         1,
		BaseToBlockPercent: 70,
	},
	ARMOR_LEATHER: {
		Name:               "Leather armor",
		ArmorClass:         3,
		BaseToBlockPercent: 30,
	},
	ARMOR_PLATE: {
		Name:               "Plate mail",
		ArmorClass:         4,
		BaseToBlockPercent: 30,
	},
	ARMOR_HIGH_AC: {
		Name:               "Duelist armor",
		ArmorClass:         3,
		BaseToBlockPercent: 50,
	},
}
