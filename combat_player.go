package main

import (
	"soulsrl/data"
	"strconv"
)

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
		if skillNumber < len(b.player.rightHand.AsWeapon.GetData().AttackPatterns) {
			skill := b.player.rightHand.AsWeapon.GetData().AttackPatterns[skillNumber]
			log.AppendMessagef("Using %s", skill.Pattern.Name)
			selected, x, y := b.selectHowToUseSkill(skill)
			if selected {
				b.applyWeaponSkill(b.player, skill, x, y)
			}
		}
	}

	// everything below is for testing, safe to delete
	if key == "p" {
		b.addMobAtRandomEmptyPlace(b.player)
		log.AppendMessage("Player dropped.")
	}
	if key == "u" {
		p := newMob("giant")
		log.AppendMessage("Big mob dropped")
		b.addMobAtRandomEmptyPlace(p)
	}
	if key == "i" {
		p := newMob("beast")
		log.AppendMessage("Medium mob dropped")
		b.addMobAtRandomEmptyPlace(p)
	}
	if key == "o" {
		p := newMob("swordmaster")
		log.AppendMessage("Small mob dropped.")
		b.addMobAtRandomEmptyPlace(p)
	}
	if key == "ENTER" {
		b.player.nextTickToAct = b.currentTick + 100
	}
}

func (b *battlefield) selectHowToUseSkill(ws *data.WeaponSkill) (bool, int, int) {
	log.AppendMessagef("Select direction for %s", ws.Pattern.Name)
	io.renderBattlefield(b, [][]int{})
	selected := false
	x, y := 0, 0
	for !selected {
		key := io.readKey()
		if key == "ESCAPE" {
			break
		}
		if key == "ENTER" && (x != 0 || y != 0) {
			selected = true
			break
		}
		x, y = readKeyToVector(key)
		var potentialCoords [][]int
		if x != 0 || y != 0 {
			potentialCoords = ws.Pattern.GetListOfCoordsWhenApplied(b.player.size, x, y)
			for i := range potentialCoords {
				potentialCoords[i][0] += b.player.x
				potentialCoords[i][1] += b.player.y
			}
		}
		io.renderBattlefield(b, potentialCoords)
	}
	return selected, x, y
}
