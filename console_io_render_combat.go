package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"soulsrl/data"
)

const (
	bf_x_offset = 1
	bf_y_offset = 2
)

//
func (c *consoleIO) renderBattlefield(b *battlefield) {
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
			if invertColorOnAction {
				c.setColorForAction(b, b.getActionPresentAt(x, y))
			} else {
				c.setStyle(fgColor, bgColor)
			}
			c.putChar(char, x, y)
		}
	}
	for _, e := range b.actions {
		if e.tickToOccur == b.currentTick {
			c.setStyle(tcell.ColorYellow, tcell.ColorBlack)
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
			if b.getActionPresentAt(x+i, y+j) != nil {
				c.setColorForAction(b, b.getActionPresentAt(x+i, y+j))
			} else {
				c.setStyle(tcell.ColorDarkRed, tcell.ColorBlack)
			}
			c.putUncoloredString(string(view[j][i]), x+i, y+j)
		}
	}
	// render dir, safe to remove
	if e.ai != nil {
		cx, cy := e.getCentralCoord()
		if b.getActionPresentAt(cx+e.ai.dirX, cy+e.ai.dirY) != nil {
			c.setColorForAction(b, b.getActionPresentAt(cx+e.ai.dirX, cy+e.ai.dirY))
		} else {
			c.setStyle(tcell.ColorDarkMagenta, tcell.ColorBlack)
		}
		c.putChar('X', cx+e.ai.dirX, cy+e.ai.dirY)
	}
}

func (c *consoleIO) renderBattlefieldUI(b *battlefield, xcoord int) {
	c.putUncoloredString(fmt.Sprintf("TICK: %d", b.currentTick), xcoord, 0)
	c.putUncoloredString(fmt.Sprintf("LIFE: %d/%d", b.player.hitpoints, 10), xcoord, 1)
	c.putUncoloredString(fmt.Sprintf("STMN: %d/%d", b.player.stamina, 10), xcoord, 2)
	c.putUncoloredString(fmt.Sprintf("STNC: STEADY"), xcoord, 2)
	currLine := 4
	for i, code := range b.player.rightHand.AsWeapon.GetData().AttackPatternCodes {
		ap := data.AttackPatternsTable[code]
		c.putUncoloredString(fmt.Sprintf("%d) %s", i+1, ap.Name), xcoord, currLine)
		currLine++
	}
	currLine++
	for _, mob := range b.mobs {
		if mob == b.player {
			continue
		}
		c.putUncoloredString(fmt.Sprintf("%s: hp %d/%d stm %d/%d", mob.name, mob.hitpoints, 10, mob.stamina, 10),
			xcoord, currLine)
		currLine++
	}
}

func (c *consoleIO) setColorForAction(b *battlefield, act *action) {
	if act != nil {
		if act.tickToOccur-b.currentTick > TICKS_IN_COMBAT_TURN {
			c.setStyle(tcell.ColorBlack, tcell.ColorYellow)
		} else {
			c.setStyle(tcell.ColorBlack, tcell.ColorRed)
		}
	}
}
