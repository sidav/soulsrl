package main

import (
	"fmt"
	"github.com/gdamore/tcell"
)

const (
	bf_x_offset = 1
	bf_y_offset = 2
)

//
func (c *consoleIO) renderBattlefield(b *battlefield, potentialActionCoords [][]int) {
	c.makeActionsMap(b)
	c.potentialActionsCoords = potentialActionCoords

	c.screen.Clear()
	c.putColorTaggedString("COMBAT: ", 0, 0)
	bfW, bfH := len(b.tiles), len(b.tiles[0])

	// render outline:
	c.setStyle(tcell.ColorWhite, tcell.ColorDarkBlue)
	c.drawRect(0, 1, bfW+1, bfH+1)

	// render the battlefield itself
	c.setOffsets(bf_x_offset, bf_y_offset)
	char := '?'
	fgColor := tcell.ColorBlack
	bgColor := tcell.ColorDarkMagenta
	invertColorOnAction := false
	for x := range b.tiles {
		for y := range b.tiles[x] {
			c.resetStyle()
			switch b.tiles[x][y] {
			case TILE_WALL:
				bgColor = tcell.ColorDarkBlue
				char = ' '
				invertColorOnAction = false
			case TILE_FLOOR:
				fgColor = tcell.ColorGray
				bgColor = tcell.ColorBlack
				invertColorOnAction = true
				char = '.'
			}
			c.setStyle(fgColor, bgColor)
			if invertColorOnAction {
				c.setColorForActionIfPresentAt(x, y)
			}
			c.putChar(char, x, y)
		}
	}
	for _, e := range b.actions {
		if e.tickToOccur == b.currentTick {
			c.setStyle(tcell.ColorRed, tcell.ColorBlack)
			c.putChar('*', e.x, e.y)
		}
	}
	c.resetStyle()
	for _, e := range b.mobs {
		c.renderMobAtCoords(b, e, e.x, e.y)
	}
	c.resetStyle()
	//c.putChar('@', b.player.x+bf_x_offset, b.player.y+bf_y_offset)
	c.setOffsets(0, 0)
	c.renderBattlefieldUI(b, bfW+2)
	c.renderLogAt(log, 0, bfH+3)
	c.screen.Show()
}

//
func (c *consoleIO) renderMobAtCoords(b *battlefield, e *mob, x, y int) {
	var view []string
	switch e.size {
	case 0, 1:
		view = []string{"@"}
	case 2:
		view = []string{
			"@|",
			"\\/",
		}
	case 3:
		view = []string{
			" o ",
			"/|0",
			"/ \\",
		}
	}
	c.resetStyle()
	for i := 0; i < e.size; i++ {
		for j := 0; j < e.size; j++ {
			c.setStyle(tcell.ColorDarkRed, tcell.ColorBlack)
			c.setColorForActionIfPresentAt(x+i, y+j)
			c.putUncoloredString(string(view[j][i]), x+i, y+j)
		}
	}
	// render dir, safe to remove
	if e.ai != nil {
		cx, cy := e.getCentralCoord()
		c.setStyle(tcell.ColorDarkMagenta, tcell.ColorBlack)
		c.setColorForActionIfPresentAt(cx+e.ai.dirX, cy+e.ai.dirY)
		c.putChar('X', cx+e.ai.dirX, cy+e.ai.dirY)
	}
}

func (c *consoleIO) renderBattlefieldUI(b *battlefield, xcoord int) {
	c.putUncoloredString(fmt.Sprintf("TICK: %d", b.currentTick), xcoord, 0)
	c.putUncoloredString(fmt.Sprintf("LIFE: %d/%d", b.player.hitpoints, b.player.getMaxStamina()), xcoord, 1)
	c.putUncoloredString(fmt.Sprintf("STMN: %d/%d", b.player.stamina, b.player.getMaxStamina()), xcoord, 2)
	c.putUncoloredString(fmt.Sprintf("STNC: STEADY"), xcoord, 2)
	currLine := 4
	for i, ap := range b.player.rightHand.AsWeapon.GetData().AttackPatterns {
		c.putUncoloredString(fmt.Sprintf("%d) %s", i+1, ap.Pattern.Name), xcoord, currLine)
		currLine++
	}
	currLine++
	for _, mob := range b.mobs {
		if mob == b.player {
			continue
		}
		c.putUncoloredString(fmt.Sprintf("%s: hp %d/%d stm %d/%d", mob.name, mob.hitpoints, mob.getMaxHitpoints(),
			mob.stamina, mob.getMaxStamina()),
			xcoord, currLine)
		currLine++
	}
}

func (c *consoleIO) makeActionsMap(b *battlefield) {
	actsmap := make([][]int, len(b.tiles))
	for i := range b.tiles {
		actsmap[i] = make([]int, len(b.tiles[i]))
	}
	for _, a := range b.actions {
		if b.containsCoords(a.x, a.y) {
			number := 2
			if a.tickToOccur - b.currentTick > TICKS_IN_COMBAT_TURN {
				number = 2
			}
			if a.tickToOccur - b.currentTick <= TICKS_IN_COMBAT_TURN {
				number = 1
			}
			actsmap[a.x][a.y] = number
		}
	}
	c.battlefieldActionsMap = actsmap
}

func (c *consoleIO) setColorForActionIfPresentAt(x, y int) {
	for _, coord := range c.potentialActionsCoords {
		if coord[0] == x && coord[1] == y {
			c.setStyle(tcell.ColorBlack, tcell.ColorDarkMagenta)
			return
		}
	}
	if x < 0 || y < 0 || x >= len(c.battlefieldActionsMap) || y >= len(c.battlefieldActionsMap[x]) {
		return
	}
	switch c.battlefieldActionsMap[x][y] {
	case 0:
		return
	case 1:
		c.setStyle(tcell.ColorBlack, tcell.ColorRed)
	case 2:
		c.setStyle(tcell.ColorBlack, tcell.ColorYellow)
	default:
		panic("no color!")
	}
}
