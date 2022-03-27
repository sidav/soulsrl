package data

import "github.com/sidav/sidavgorandom/fibrandom"

type Weapon struct {
	Code int
}

func (w *Weapon) GetData() *weaponData {
	return weaponsTable[w.Code]
}

func (w *Weapon) RollDamage(rnd *fibrandom.FibRandom) int {
	data := w.GetData()
	totalHits := 0
	for i := 0; i < data.attackRating; i++ {
		rand := rnd.RandInRange(1, 100)
		if rand < data.baseToHitPercent {
			totalHits++
		}
	}
	return totalHits
}

type WeaponSkill struct {
	Pattern                   *SkillPattern
	DurationInTurnLengths     int
	WeaponDamageAmountPercent int
	StaminaCost               int
	IsInstant                 bool
}

func (ws *WeaponSkill) GetDurationForTurnTicks(ticksPerTurn int) int {
	return ticksPerTurn * ws.DurationInTurnLengths / 10
}

type weaponData struct {
	Name             string
	attackRating     int
	baseToHitPercent int
	AttackPatterns   []*WeaponSkill
}

const (
	WEAPON_BROKENSWORD = iota
	WEAPON_DEBUGGER
	WEAPON_RAPIER
	WEAPON_SHORTSWORD
	WEAPON_LONGSWORD
	WEAPON_GIANTS_LONGSWORD
	WEAPON_SPEAR
)

var weaponsTable = map[int]*weaponData{
	WEAPON_DEBUGGER: {
		Name:     "Sword of Holy Debug",
		attackRating: 8,
		baseToHitPercent: 30,
		AttackPatterns: []*WeaponSkill{
			{
				Pattern:                   AttackPatternsTable[APATTERN_RIGHT_STRIKE_SIDESTEP],
				DurationInTurnLengths:     15,
				WeaponDamageAmountPercent: 75,
				StaminaCost:               2,
			},
			{
				Pattern:                   AttackPatternsTable[APATTERN_STRIKE_STEP_BACK],
				DurationInTurnLengths:     15,
				WeaponDamageAmountPercent: 75,
				StaminaCost:               2,
			},
			{
				Pattern:                   AttackPatternsTable[APATTERN_JUMP_LUNGE],
				DurationInTurnLengths:     10,
				IsInstant:                 true,
				WeaponDamageAmountPercent: 150,
				StaminaCost:               5,
			},
			{
				Pattern:                   AttackPatternsTable[APATTERN_STRIKE_JUMP_OVER],
				DurationInTurnLengths:     10,
				IsInstant:                 true,
				WeaponDamageAmountPercent: 150,
				StaminaCost:               5,
			},
		},
	},
	WEAPON_BROKENSWORD: {
		Name:     "Broken Sword",
		attackRating: 5,
		baseToHitPercent: 30,
		AttackPatterns: []*WeaponSkill{
			{
				Pattern:                   AttackPatternsTable[APATTERN_SIMPLE_STRIKE],
				DurationInTurnLengths:     10,
				WeaponDamageAmountPercent: 100,
				StaminaCost:               1,
			},
		},
	},
	WEAPON_RAPIER: {
		Name:     "Rapier",
		attackRating: 4,
		baseToHitPercent: 70,
		AttackPatterns: []*WeaponSkill{
			{
				Pattern:                   AttackPatternsTable[APATTERN_SIMPLE_STRIKE],
				DurationInTurnLengths:     10,
				WeaponDamageAmountPercent: 100,
				StaminaCost:               1,
			},
			{
				Pattern:                   AttackPatternsTable[APATTERN_STRIKE_STEP_BACK],
				DurationInTurnLengths:     15,
				WeaponDamageAmountPercent: 75,
				StaminaCost:               2,
			},
			{
				Pattern:                   AttackPatternsTable[APATTERN_RIGHT_STRIKE_SIDESTEP],
				DurationInTurnLengths:     15,
				WeaponDamageAmountPercent: 75,
				StaminaCost:               2,
			},
		},
	},
	WEAPON_SHORTSWORD: {
		Name:     "Short Sword",
		attackRating: 5,
		baseToHitPercent: 40,
		AttackPatterns: []*WeaponSkill{
			{
				Pattern:                   AttackPatternsTable[APATTERN_SIMPLE_STRIKE],
				DurationInTurnLengths:     10,
				WeaponDamageAmountPercent: 100,
				StaminaCost:               1,
			},
			{
				Pattern:                   AttackPatternsTable[APATTERN_RIGHT_SLASH],
				DurationInTurnLengths:     15,
				WeaponDamageAmountPercent: 75,
				StaminaCost:               2,
			},
		},
	},
	WEAPON_LONGSWORD: {
		Name:     "Long Sword",
		attackRating: 7,
		baseToHitPercent: 30,
		AttackPatterns: []*WeaponSkill{
			{
				Pattern:                   AttackPatternsTable[APATTERN_SIMPLE_STRIKE],
				DurationInTurnLengths:     10,
				WeaponDamageAmountPercent: 100,
				StaminaCost:               1,
			},
			{
				Pattern:                   AttackPatternsTable[APATTERN_SLASH],
				DurationInTurnLengths:     25,
				WeaponDamageAmountPercent: 50,
				StaminaCost:               3,
			},
		},
	},
	WEAPON_GIANTS_LONGSWORD: {
		Name:     "Giant Long Sword",
		attackRating: 15,
		baseToHitPercent: 30,
		AttackPatterns: []*WeaponSkill{
			{
				Pattern:                   AttackPatternsTable[APATTERN_SIMPLE_STRIKE],
				DurationInTurnLengths:     15,
				WeaponDamageAmountPercent: 100,
				StaminaCost:               1,
			},
			{
				Pattern:                   AttackPatternsTable[APATTERN_SLASH],
				DurationInTurnLengths:     35,
				WeaponDamageAmountPercent: 50,
				StaminaCost:               3,
			},
		},
	},
	WEAPON_SPEAR: {
		Name:     "Spear",
		attackRating: 5,
		baseToHitPercent: 40,
		AttackPatterns: []*WeaponSkill{
			{
				Pattern:                   AttackPatternsTable[APATTERN_SIMPLE_STRIKE],
				DurationInTurnLengths:     10,
				WeaponDamageAmountPercent: 100,
				StaminaCost:               1,
			},
			{
				Pattern:                   AttackPatternsTable[APATTERN_LUNGE],
				DurationInTurnLengths:     25,
				WeaponDamageAmountPercent: 150,
				StaminaCost:               2,
			},
			{
				Pattern:                   AttackPatternsTable[APATTERN_JUMP_LUNGE],
				DurationInTurnLengths:     10,
				IsInstant:                 true,
				WeaponDamageAmountPercent: 150,
				StaminaCost:               5,
			},
		},
	},
}
