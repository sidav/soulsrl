package main

import "time"

func (b *battlefield) combatGameLoop() {
	io.renderBattlefield(b)
	if b.currentTick % 100 == 0 {
		key := io.readKey()
		if key == "ESCAPE" {
			exitGame = true
		}
	} else {
		time.Sleep(50*time.Millisecond)
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
