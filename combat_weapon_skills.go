package main

import (
	"soulsrl/data"
	"soulsrl/geometry"
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
		tickToOccur = b.currentTick+1 // it's for action occurrence, as mob delay is already set.
	}
	acting.stamina -= skill.StaminaCost

	patternCoords := skill.Pattern.GetListOfCoordsWhenAppliedAtRect(acting.x, acting.y, acting.size, tx, ty, tsize)
	toHitRoll := weapon.RollToHitDice(rnd)
	damageRoll := weapon.RollDamageDice(rnd)
	damageRoll = skill.WeaponDamageAmountPercent * damageRoll / 100
	for _, coord := range patternCoords {
		if !b.containsCoords(coord[0], coord[1]) {
			continue
		}
		b.actions = append(b.actions, &action{
			tickToOccur: tickToOccur,
			owner:       acting,
			damageRoll:  damageRoll,
			toHitRoll:   toHitRoll,
			x:           coord[0],
			y:           coord[1],
		})
	}
	if skill.HasMoveTo {
		// calculate movement vector
		vx, vy := (tx+tsize/2)-acting.x, (ty+tsize/2)-acting.y
		if vx != 0 {
			vx /= geometry.Abs(vx)
		}
		if vy != 0 {
			vy /= geometry.Abs(vy)
		}
		for i := 0; i < acting.size; i++ {
			b.actions = append(b.actions, &action{
				actionType:  ACTIONTYPE_MOVE,
				tickToOccur: tickToOccur,
				owner:       acting,
				vx:          vx, vy: vy,
			})
		}
	}
}
