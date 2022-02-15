package main

import (
	"soulsrl/data"
	"soulsrl/geometry"
	"strings"
)

type mob struct {
	// topleft coord
	x, y, size    int
	dirX, dirY    int
	nextTickToAct int

	name string

	rightHand *data.Item
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

func newMob(name string, x, y int) *mob {
	u := &mob{
		x:    x,
		y:    y,
		name: name,
	}
	switch strings.ToLower(name) {
	case "giant":
		u.size = 3
		u.rightHand = &data.Item{AsWeapon: &data.Weapon{Code: data.WEAPON_LONGSWORD}}
	case "swordmaster":
		u.size = 1
		u.rightHand = &data.Item{AsWeapon: &data.Weapon{Code: data.WEAPON_SPEAR}}
	}

	return u
}
