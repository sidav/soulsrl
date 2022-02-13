package main

func (b *battlefield) combatGameLoop() {
	key := io.readKey()
	if key == "ESCAPE" {
		exitGame = true
	}

	for _, e := range b.units {
		if e.nextTickToAct > b.currentTick {
			continue
		}
		e.dirX, e.dirY = stupidRotateVector45(e.dirX, e.dirY)
		e.nextTickToAct = b.currentTick + e.ap.ticksToPerform + 5
		b.applyAttackPattern(e, e.ap, e.dirX, e.dirY)
	}

	b.applyEvents()
	b.currentTick++
}
