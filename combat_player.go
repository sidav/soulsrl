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
	moved := b.movePlayerOrDefaultHit(key)
	if moved {
		return
	}
	if key == "v" {
		selected, x, y := b.selectDirection("Select dodge direction")
		if selected {
			rolled := b.tryRollMobByVector(b.player, x, y)
			if !rolled {
				log.AppendMessagef("Can't dodge!")
				return
			}
		}
	}

	if '1' <= rune(key[0]) && rune(key[0]) <= '9' {
		skillNumber, _ := strconv.Atoi(key)
		skillNumber-- // because numeration is from 0 in code
		if skillNumber < len(b.player.rightHand.AsWeapon.GetData().AttackPatterns) {
			skill := b.player.rightHand.AsWeapon.GetData().AttackPatterns[skillNumber]
			log.AppendMessagef("Using %s", skill.Pattern.Name)
			selected, x, y := b.selectHowToUseSkill(skill, key)
			if selected {
				b.applyWeaponSkill(b.player, b.player.rightHand.AsWeapon,
					skill, b.player.x+x*b.player.size, b.player.y+y*b.player.size, b.player.size)
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

func (b *battlefield) movePlayerOrDefaultHit(key string) bool {
	dirx, diry := readKeyToVector(key)
	if !(dirx == 0 && diry == 0) {
		moved := b.tryMoveMobByVector(b.player, dirx, diry)
		if moved {
			return true
		}

		mobAtCoords := b.getMobPresentAt(b.player.x+b.player.size*dirx, b.player.y+b.player.size*diry)
		if mobAtCoords != nil {
			b.applyWeaponSkill(b.player, b.player.rightHand.AsWeapon,
				b.player.rightHand.AsWeapon.GetData().AttackPatterns[0],
				b.player.x+dirx*b.player.size, b.player.y+diry*b.player.size, b.player.size)
			return true
		}
	}
	return false
}

func (b *battlefield) selectDirection(text string) (bool, int, int) {
	log.AppendMessagef(text)
	io.renderBattlefield(b, [][]int{})
	x, y := 0, 0
	selected := false
	for !selected {
		key := io.readKey()
		if key == "ESCAPE" {
			break
		}
		x, y = readKeyToVector(key)
		if x != 0 || y != 0 {
			selected = true
			break
		}
	}
	return selected, x, y
}

func (b *battlefield) selectHowToUseSkill(ws *data.WeaponSkill, confirmButton string) (bool, int, int) {
	log.AppendMessagef("Select direction for %s", ws.Pattern.Name)
	io.renderBattlefield(b, [][]int{})
	selected := false
	x, y := 0, 0
	for !selected {
		key := io.readKey()
		if key == "ESCAPE" {
			break
		}
		if (key == "ENTER" || key == confirmButton) && (x != 0 || y != 0) {
			selected = true
			break
		}
		x, y = readKeyToVector(key)
		var potentialCoords [][]int
		if x != 0 || y != 0 {
			potentialCoords = ws.Pattern.GetListOfCoordsWhenAppliedAtRect(b.player.x, b.player.y, b.player.size,
				b.player.x+x*b.player.size, b.player.y+y*b.player.size, b.player.size)
			//for i := range potentialCoords {
			//	potentialCoords[i][0] += b.player.x
			//	potentialCoords[i][1] += b.player.y
			//}
		}
		io.renderBattlefield(b, potentialCoords)
	}
	return selected, x, y
}
