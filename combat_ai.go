package main

func (b *battlefield) actAsMob(m *mob) {
	if m.dirX == 0 && m.dirY == 0 || rnd.OneChanceFrom(10) {
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
		if b.isRectFullyPassable(m.x+m.dirX, m.y+m.dirY, m.size) {
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
