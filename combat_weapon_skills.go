package main

import (
	"soulsrl/data"
)

//func (b *battlefield) willWeaponSkillReachSquare(acting *mob, skill data.WeaponSkill, tx, ty, tsize int) bool {
//	patternCoords := skill.Pattern.GetListOfCoordsWhenAppliedAtRect(acting.x, acting.y, acting.size, tx, ty, tsize)
//	for _, c := range patternCoords {
//		if geometry.RectContainsCoords(tx, ty, tsize, tsize, c[0], c[1]) {
//			return true
//		}
//	}
//	return false
//}

func (b *battlefield) applyWeaponSkill(acting *mob, weapon *data.Weapon, skill *data.WeaponSkill, tx, ty, tsize int) {
	if acting.stamina < skill.StaminaCost {
		return
	}
	tickToOccur := b.currentTick + skill.GetDurationForTurnTicks(TICKS_IN_COMBAT_TURN)
	acting.nextTickToAct = tickToOccur
	if skill.IsInstant {
		tickToOccur = b.currentTick + 1 // it's for action occurrence, as mob delay is already set.
	}
	acting.stamina -= skill.StaminaCost

	patternCoords, moveVectors := skill.Pattern.GetListOfCoordsWhenAppliedAtRect(acting.x, acting.y, acting.size, tx, ty, tsize)
	for _, coord := range patternCoords {
		if !b.containsCoords(coord[0], coord[1]) {
			continue
		}
		b.actions = append(b.actions, &action{
			tickToOccur:         tickToOccur,
			owner:               acting,
			rolledDamage:        weapon.RollDamage(rnd),
			damageAmountPercent: skill.WeaponDamageAmountPercent,
			x:                   coord[0],
			y:                   coord[1],
		})
	}
	for _, vect := range moveVectors {
		actionType := ACTIONTYPE_MOVE
		if skill.Pattern.MovementIsJumpOver {
			actionType = ACTIONTYPE_JUMPOVER
		}
		b.actions = append(b.actions, &action{
			vx:          vect[0],
			vy:          vect[1],
			tickToOccur: tickToOccur,
			actionType:  actionType,
			hidden:      true,
			owner:       acting,
		})
	}
}
