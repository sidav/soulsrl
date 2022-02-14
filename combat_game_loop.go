package main

func (b *battlefield) combatGameLoop() {
	io.renderBattlefield(b)
	key := io.readKey()
	if key == "ESCAPE" {
		exitGame = true
	}

	for _, e := range b.mobs {
		if e.nextTickToAct > b.currentTick {
			continue
		}
		b.actAsMob(e)
	}

	b.applyActions()
	b.currentTick++
}
