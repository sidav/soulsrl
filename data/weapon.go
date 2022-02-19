package data

import "github.com/sidav/sidavgorandom/fibrandom"

type Weapon struct {
	Code int
}

func (w *Weapon) GetData() *weaponData {
	return weaponsTable[w.Code]
}

func (w *Weapon) RollDamageDice(rnd *fibrandom.FibRandom) int {
	data := w.GetData()
	return rnd.RollDice(data.dnum, data.dval, data.dmod)
}

type WeaponSkill struct {
	Pattern                   *SkillPattern
	DurationInTurnLengths     int
	WeaponDamageAmountPercent int
}

func (ws *WeaponSkill) GetDurationForTurnTicks(ticksPerTurn int) int {
	return ticksPerTurn * ws.DurationInTurnLengths / 10
}

type weaponData struct {
	Name             string
	dnum, dval, dmod int
	AttackPatterns   []*WeaponSkill
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
		AttackPatterns: []*WeaponSkill{
			{
				Pattern:               AttackPatternsTable[APATTERN_SIMPLE_STRIKE],
				DurationInTurnLengths: 10,
				WeaponDamageAmountPercent: 100,
			},
		},
	},
	WEAPON_SHORTSWORD: {
		Name: "Short Sword",
		dnum: 2,
		dval: 6,
		dmod: 0,
		AttackPatterns: []*WeaponSkill{
			{
				Pattern:               AttackPatternsTable[APATTERN_SIMPLE_STRIKE],
				DurationInTurnLengths: 10,
				WeaponDamageAmountPercent: 100,
			},
			{
				Pattern:               AttackPatternsTable[APATTERN_RIGHT_SLASH],
				DurationInTurnLengths: 15,
				WeaponDamageAmountPercent: 75,
			},
		},
	},
	WEAPON_LONGSWORD: {
		Name: "Long Sword",
		dnum: 2,
		dval: 6,
		dmod: 2,
		AttackPatterns: []*WeaponSkill{
			{
				Pattern:               AttackPatternsTable[APATTERN_SIMPLE_STRIKE],
				DurationInTurnLengths: 10,
				WeaponDamageAmountPercent: 100,
			},
			{
				Pattern:               AttackPatternsTable[APATTERN_SLASH],
				DurationInTurnLengths: 25,
				WeaponDamageAmountPercent: 50,
			},
		},
	},
	WEAPON_SPEAR: {
		Name: "Spear",
		dnum: 1,
		dval: 6,
		dmod: 3,
		AttackPatterns: []*WeaponSkill{
			{
				Pattern:               AttackPatternsTable[APATTERN_SIMPLE_STRIKE],
				DurationInTurnLengths: 10,
				WeaponDamageAmountPercent: 100,
			},
			{
				Pattern:               AttackPatternsTable[APATTERN_LUNGE],
				DurationInTurnLengths: 25,
				WeaponDamageAmountPercent: 150,
			},
		},
	},
}
