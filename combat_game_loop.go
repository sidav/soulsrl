package main

import (
	"soulsrl/data"
	"strconv"
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

var fuck = 0
func (b *battlefield) workPlayerInput() {

	key := io.readKey()

	if key == "ESCAPE" {
		exitGame = true
	}
	if key == " " {
		b.player.nextTickToAct = b.currentTick + 1
	}
	dirx, diry := readKeyToVector(key)
	if !(dirx == 0 && diry == 0) {
		b.tryMoveMobByVector(b.player, dirx, diry)
	}

	if '1' <= rune(key[0]) && rune(key[0]) <= '9' {
		skillNumber, _ := strconv.Atoi(key)
		skillNumber--  // because numeration is from 0 in code
		if skillNumber < len(b.player.rightHand.AsWeapon.GetData().AttackPatternCodes) {
			ap := data.AttackPatternsTable[b.player.rightHand.AsWeapon.GetData().AttackPatternCodes[skillNumber]]
			log.AppendMessagef("Using %s", ap.Name)
		}
	}

	// everything below is for testing, safe to delete
	if key == "p" {
		b.addMobAtRandomEmptyPlace(b.player)
		log.AppendMessage("Player dropped.")
	}
	if key == "o" {
		p := newMob("giant")
		log.AppendMessage("Big mob dropped")
		b.addMobAtRandomEmptyPlace(p)
	}
	if key == "l" {
		p := newMob("swordmaster")
		log.AppendMessage("Small mob dropped.")
		b.addMobAtRandomEmptyPlace(p)
	}
	if key == "ENTER" {
		b.player.nextTickToAct = b.currentTick + 100
	}
}
