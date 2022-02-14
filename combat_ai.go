package main

import "soulsrl/line"

func (b *battlefield) actAsMob(m *mob) {
	newx, newy := m.x+m.dirX, m.y+m.dirY
	if m.dirX == 0 && m.dirY == 0 || !b.containsCoords(newx, newy) || rnd.OneChanceFrom(10) {
		m.dirX, m.dirY = rnd.RandomUnitVectorInt(true)
		m.nextTickToAct = b.currentTick + TICKS_IN_COMBAT_TURN
		return
	}

	if rnd.OneChanceFrom(5) && b.tryAttackAsMob(m) {
		return
	} else {
		// move by coords
		mobAtCoords := b.getMobInSquareOtherThan(m.x+m.dirX, m.y+m.dirY, m.size, m)
		if b.isRectFullyPassable(m.x+m.dirX, m.y+m.dirY, m.size) && mobAtCoords == nil {
			m.x += m.dirX
			m.y += m.dirY
			m.nextTickToAct = b.currentTick + TICKS_IN_COMBAT_TURN
		} else {
			m.dirX, m.dirY = rnd.RandomUnitVectorInt(true)
			m.nextTickToAct = b.currentTick + TICKS_IN_COMBAT_TURN
			return
		}
	}
}

func (b *battlefield) tryAttackAsMob(m *mob) bool {
	mcx, mcy := m.getCentralCoord()
	for _, anotherMob := range b.mobs {
		if anotherMob == m {
			continue
		}
		amcx, amcy := anotherMob.getCentralCoord()
		if orthogonalDistance(mcx, mcy, amcx, amcy) <= m.size+m.size/2+anotherMob.size+anotherMob.size/2 {
			m.dirX, m.dirY = line.GetNextStepForLine(mcx, mcy, amcx, amcy)
			b.applyAttackPattern(m, m.ap, m.dirX, m.dirY)
			m.nextTickToAct = b.currentTick + m.ap.ticksToPerform
			return true
		}
	}
	return false
}
