package main

import (
	"soulsrl/data"
	"soulsrl/geometry"
	"soulsrl/geometry/line"
)

func (b *battlefield) actAsMob(m *mob) {
	newx, newy := m.x+m.dirX, m.y+m.dirY
	if m.dirX == 0 && m.dirY == 0 || !b.containsCoords(newx, newy) || rnd.OneChanceFrom(10) {
		coordsList := b.getListOfVectorsToPassableCoordsForMob(m)
		if len(coordsList) > 0 {
			selected := coordsList[rnd.Rand(len(coordsList))]
			m.dirX, m.dirY = selected[0], selected[1]
			m.nextTickToAct = b.currentTick + TICKS_IN_COMBAT_TURN
			return
		}
	}

	if rnd.OneChanceFrom(5) && b.tryAttackAsMob(m) {
		return
	} else {
		// move by coords
		mobAtCoords := b.getMobInSquareOtherThan(m.x+m.dirX, m.y+m.dirY, m.size, m)
		if b.areAllTilesInRectPassable(m.x+m.dirX, m.y+m.dirY, m.size) && mobAtCoords == nil {
			m.x += m.dirX
			m.y += m.dirY
			m.nextTickToAct = b.currentTick + TICKS_IN_COMBAT_TURN
		} else {
			m.dirX, m.dirY = 0, 0 // so it will be changed later
			return
		}
	}
}

func (b *battlefield) tryAttackAsMob(m *mob) bool {
	for _, anotherMob := range b.mobs {
		if anotherMob == m {
			continue
		}
		var applicableAttacks []int
		for _, apc := range m.rightHand.AsWeapon.GetData().AttackPatternCodes {
			ap := data.AttackPatternsTable[apc]
			attackReach := ap.ReachInUnitSizes * m.size
			if geometry.DistanceBetweenSquares(m.x, m.y, m.size, anotherMob.x, anotherMob.y, anotherMob.size) <= m.size*attackReach {
				applicableAttacks = append(applicableAttacks, apc)
			}
		}
		if len(applicableAttacks) > 0 {
			mcx, mcy := m.getCentralCoord()
			amcx, amcy := anotherMob.getCentralCoord()
			ap := data.AttackPatternsTable[applicableAttacks[rnd.Rand(len(applicableAttacks))]]
			m.dirX, m.dirY = line.GetNextStepForLine(mcx, mcy, amcx, amcy)
			b.applyAttackPattern(m, ap, m.dirX, m.dirY)
			m.nextTickToAct = b.currentTick + ap.GetDurationForTurnTicks(TICKS_IN_COMBAT_TURN)
			return true
		}
	}
	return false
}
