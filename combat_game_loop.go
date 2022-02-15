package main

import "time"

func (b *battlefield) combatGameLoop() {
	nextTickToGiveControls := b.currentTick
	for !exitGame {
		io.renderBattlefield(b)
		if b.currentTick >= nextTickToGiveControls {
			key := io.readKey()
			if key == "ESCAPE" {
				exitGame = true
			}
			if key == "SPACE" {
				nextTickToGiveControls = b.currentTick + 1
			}
			if key == "ENTER" {
				nextTickToGiveControls = b.currentTick + 100
			}
		} else {
			time.Sleep(50 * time.Millisecond)
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
}
