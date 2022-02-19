package main

import (
	"time"
)

func (b *battlefield) combatGameLoop() {
	b.player = newMob("player")
	for !exitGame {
		b.LoopThroughMobs()
		b.applyActions()

		io.renderBattlefield(b, [][]int{})
		for !exitGame && b.player.nextTickToAct <= b.currentTick {
			b.workPlayerInput()
			io.renderBattlefield(b, [][]int{})
		}
		time.Sleep(40 * time.Millisecond)

		for _, e := range b.mobs {
			e.wasAlreadyAffectedByActionBy = nil
			if e.nextTickToAct > b.currentTick || e.ai == nil {
				continue
			}
			b.actAsMob(e)
		}

		b.cleanupActions()

		b.currentTick++
	}
}

func (b *battlefield) LoopThroughMobs() {
	for i := 0; i < len(b.mobs); i++ {
		mob := b.mobs[i]
		if mob.hitpoints <= 0 {
			b.mobs[i] = b.mobs[len(b.mobs)-1]
			b.mobs = append(b.mobs[:len(b.mobs)-1])
			i--
			continue
		}
		if b.currentTick % (TICKS_IN_COMBAT_TURN * 2) == 0 && mob.nextTickToAct <= b.currentTick {
			mob.stamina++
			if mob.stamina > mob.getMaxStamina() {
				mob.stamina = mob.getMaxStamina()
			}
		}
	}
}
