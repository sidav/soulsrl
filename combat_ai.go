package main

func (b *battlefield) actAsMob(m *mob) {
	newx, newy := m.x+m.dirX, m.y+m.dirY
	if m.dirX == 0 && m.dirY == 0 || !b.containsCoords(newx, newy) || rnd.OneChanceFrom(10) {
		m.dirX, m.dirY = rnd.RandomUnitVectorInt(true)
		m.nextTickToAct = b.currentTick + TICKS_IN_COMBAT_TURN
		return
	}

	if rnd.OneChanceFrom(10) {
		// attack coords
		b.applyAttackPattern(m, m.ap, m.dirX, m.dirY)
		m.nextTickToAct = b.currentTick + m.ap.ticksToPerform
	} else {
		// move by coords
		if b.isRectFullyPassable(m.x+m.dirX, m.y+m.dirY, m.size) && b.getMobInSquare(m.x+m.dirX, m.y+m.dirY, m.size) == nil {
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
