package main

import (
	"soulsrl/data"
	"soulsrl/geometry"
)

func (b *battlefield) willWeaponSkillReachSquare(acting *mob, skill data.WeaponSkill, tx, ty, tsize int) bool {
	patternCoords := skill.Pattern.GetListOfCoordsWhenAppliedAtRect(acting.x, acting.y, acting.size, tx, ty, tsize)
	for _, c := range patternCoords {
		if geometry.RectContainsCoords(tx, ty, tsize, tsize, c[0], c[1]) {
			return true
		}
	}
	return false
}

func (b *battlefield) applyWeaponSkill(acting *mob, weaponSkill *data.WeaponSkill, tx, ty, tsize int) {
	tickToOccur := b.currentTick + weaponSkill.GetDurationForTurnTicks(TICKS_IN_COMBAT_TURN)
	acting.nextTickToAct = tickToOccur
	patternCoords := weaponSkill.Pattern.GetListOfCoordsWhenAppliedAtRect(acting.x, acting.y, acting.size, tx, ty, tsize)
	for _, coord := range patternCoords {
		if !b.containsCoords(coord[0], coord[1]) {
			continue
		}
		b.actions = append(b.actions, &action{
			tickToOccur: tickToOccur,
			owner:       acting,
			x:           coord[0],
			y:           coord[1],
		})
	}
}
