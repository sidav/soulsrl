package main

import (
	"time"
)

func (b *battlefield) combatGameLoop() {
	b.player = newMob("player")
	for !exitGame {
		io.renderBattlefield(b)
		for !exitGame && b.player.nextTickToAct <= b.currentTick {
			b.workPlayerInput()
			io.renderBattlefield(b)
		}
		time.Sleep(40 * time.Millisecond)

		for _, e := range b.mobs {
			e.wasAlreadyAffectedByActionBy = nil
			if e.nextTickToAct > b.currentTick || e.ai == nil {
				continue
			}
			b.actAsMob(e)
		}

		b.applyActions()
		b.currentTick++
	}
}
