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

	name string

	rightHand *data.Item

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

func newMob(name string) *mob {
	u := &mob{
		name: name,
	}
	switch strings.ToLower(name) {
	case "player":
		u.size = 1
		u.rightHand = &data.Item{AsWeapon: &data.Weapon{Code: data.WEAPON_SHORTSWORD}}
		return u
	case "giant":
		u.size = 3
		u.rightHand = &data.Item{AsWeapon: &data.Weapon{Code: data.WEAPON_LONGSWORD}}
	case "swordmaster":
		u.size = 1
		u.rightHand = &data.Item{AsWeapon: &data.Weapon{Code: data.WEAPON_SPEAR}}
	}
	u.ai = initDefaultAi()
	return u
}
