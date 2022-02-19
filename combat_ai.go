package main

import (
	"soulsrl/data"
	"soulsrl/geometry"
)

type mobAi struct {
	dirX, dirY                                                int
	changeDirPercent, attackPercent, changeDirInCombatPercent int
}

func initDefaultAi() *mobAi {
	return &mobAi{
		changeDirPercent:         10,
		attackPercent:            50,
		changeDirInCombatPercent: 5,
	}
}

func (b *battlefield) actAsMob(m *mob) {
	newx, newy := m.x+m.ai.dirX, m.y+m.ai.dirY
	if m.ai.dirX == 0 && m.ai.dirY == 0 || !b.containsCoords(newx, newy) || rnd.PercentChance(m.ai.changeDirPercent) {
		if b.tryRotateToEnemy(m) {
			return 
		}
		coordsList := b.getListOfVectorsToPassableCoordsForMob(m)
		if len(coordsList) > 0 {
			selected := coordsList[rnd.Rand(len(coordsList))]
			m.ai.dirX, m.ai.dirY = selected[0], selected[1]
			m.nextTickToAct = b.currentTick + TICKS_IN_COMBAT_TURN
			return
		}
	}

	if rnd.PercentChance(m.ai.attackPercent) {
		if b.tryAttackAsMob(m) {
			return
		}
	} else {
		// move by coords
		moved := b.tryMoveMobByVector(m, m.ai.dirX, m.ai.dirY)
		if !moved {
			if rnd.PercentChance(m.ai.changeDirInCombatPercent) {
				m.ai.dirX, m.ai.dirY = 0, 0 // so it will be changed later
				return
			}
		}
	}
}

func (b *battlefield) tryRotateToEnemy(m *mob) bool {
	for _, anotherMob := range b.mobs {
		if anotherMob == m {
			continue
		}
		vx, vy := b.getVectorForVisibleFromMobToMobIfExists(m, anotherMob)
		if vx != 0 || vy != 0 {
			m.ai.dirX, m.ai.dirY = vx, vy
			return true
		}
	}
	return false
}

func (b *battlefield) tryAttackAsMob(m *mob) bool {
	for _, anotherMob := range b.mobs {
		if anotherMob == m {
			continue
		}
		var applicableAttacks []*data.WeaponSkill
		for _, wskill := range m.rightHand.AsWeapon.GetData().AttackPatterns {
			ap := wskill.Pattern
			attackReach := ap.ReachInUnitSizes * m.size
			if geometry.DistanceBetweenSquares(m.x, m.y, m.size, anotherMob.x, anotherMob.y, anotherMob.size) <= attackReach {
				applicableAttacks = append(applicableAttacks, wskill)
			}
		}
		if len(applicableAttacks) > 0 {
			ap := applicableAttacks[rnd.Rand(len(applicableAttacks))]
			b.applyWeaponSkill(m, ap, anotherMob.x, anotherMob.y, anotherMob.size)
		}
	}
	return false
}
