package main

import (
	"soulsrl/data"
	"soulsrl/geometry"
	"strings"
)

type mob struct {
	ai *mobAi
	// topleft coord
	x, y, size    int
	nextTickToAct int

	hitpoints, stamina int
	stats              *mobStats

	name string

	// weapons/shields/catalysts
	rightHand *data.Item

	//armor
	body *data.Item

	// temp vars
	wasAlreadyAffectedByActionBy *mob
}

func (m *mob) getCentralCoord() (int, int) {
	if m.size == 0 || m.size == 1 {
		return m.x, m.y
	}
	return m.x + m.size/2, m.y + m.size/2
}

func (m *mob) containsCoords(x, y int) bool {
	return geometry.SquareContainsCoords(m.x, m.y, m.size, x, y)
}

var mobsTable = map[string]mob{
	"player": {
		size: 1,
		name:      "Player",
		rightHand: &data.Item{AsWeapon: &data.Weapon{Code: data.WEAPON_RAPIER}},
		body: &data.Item{AsArmor: &data.Armor{Code: data.ARMOR_LEATHER}},
		stats: &mobStats{
			vitality:  100,
			endurance: 10,
			dexterity: 10,
			strength:  10,
		},
	},
	"giant": {
		size: 3,
		name:      "Giant",
		rightHand: &data.Item{AsWeapon: &data.Weapon{Code: data.WEAPON_GIANTS_LONGSWORD}},
		body: &data.Item{AsArmor: &data.Armor{Code: data.ARMOR_LEATHER}},
		stats: &mobStats{
			vitality:  15,
			endurance: 5,
			dexterity: 5,
			strength:  15,
		},
	},
	"beast": {
		size: 2,
		name:      "Undead Serpent",
		rightHand: &data.Item{AsWeapon: &data.Weapon{Code: data.WEAPON_GIANTS_LONGSWORD}},
		body: &data.Item{AsArmor: &data.Armor{Code: data.ARMOR_LEATHER}},
		stats: &mobStats{
			vitality:  25,
			endurance: 10,
			dexterity: 25,
			strength:  5,
		},
	},
	"spearman": {
		size: 1,
		name:      "Undead Spearman",
		rightHand: &data.Item{AsWeapon: &data.Weapon{Code: data.WEAPON_SPEAR}},
		body: &data.Item{AsArmor: &data.Armor{Code: data.ARMOR_LEATHER}},
		stats: &mobStats{
			vitality:  10,
			endurance: 7,
			dexterity: 7,
			strength:  7,
		},
	},
	"swordmaster": {
		size: 1,
		name:      "Undead Swordmaster",
		rightHand: &data.Item{AsWeapon: &data.Weapon{Code: data.WEAPON_SHORTSWORD}},
		body: &data.Item{AsArmor: &data.Armor{Code: data.ARMOR_HIGH_AC}},
		stats: &mobStats{
			vitality:  12,
			endurance: 10,
			dexterity: 5,
			strength:  7,
		},
	},
}

func newMob(name string) *mob {
	code := strings.ToLower(name)
	mob := mobsTable[code]
	if code != "player" {
		mob.ai = initDefaultAi()
	}
	mob.hitpoints = mob.getMaxHitpoints()
	mob.stamina = mob.getMaxStamina()
	return &mob
}
