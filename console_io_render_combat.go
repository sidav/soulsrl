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
func (c *consoleIO) renderBattlefield(b *battlefield) {
	c.screen.Clear()
	c.putColorTaggedString("COMBAT: ", 0, 0)
	bfW, bfH := len(b.tiles), len(b.tiles[0])

	// render outline:
	c.setStyle(tcell.ColorWhite, tcell.ColorDarkRed)
	c.drawRect(0, 1, bfW+1, bfH+1)
	// render the battlefield itself
	for x := range b.tiles {
		for y := range b.tiles[x] {
			switch b.tiles[x][y] {
			case TILE_WALL:
				c.style = c.style.Background(tcell.ColorDarkRed)
				c.putChar(' ', x+bf_x_offset, y+bf_y_offset)
			case TILE_FLOOR:
				c.resetStyle()
				c.putChar('.', x+bf_x_offset, y+bf_y_offset)
			}
		}
	}
	for _, e := range b.events {
		c.setStyle(tcell.ColorBlack, tcell.ColorDarkMagenta)
		c.putChar(' ', e.x+bf_x_offset, e.y+bf_y_offset)
	}
	c.resetStyle()
	for _, e := range b.units {
		c.renderMobAtCoords(e, e.x+bf_x_offset, e.y+bf_y_offset)
	}
	c.resetStyle()
	c.putUncoloredString(fmt.Sprintf("TICK %d", b.currentTick), bfW+bf_x_offset+1, 1)
	//c.putChar('@', b.player.x+bf_x_offset, b.player.y+bf_y_offset)
	//c.renderPlayerBattlefieldUI(bf_x_offset+bfW+1, b)
	//c.renderLogAt(log, 0, bf_y_offset+bfH+1)
	c.screen.Show()
}
//
func (c *consoleIO) renderMobAtCoords(e *mob, x, y int) {
	var view []string
	switch e.size {
	case 0, 1:
		view = []string{"@"}
	case 2:
		view = []string{
			"@@",
			"@@",
		}
	case 3:
		view = []string{
			" o ",
			"/|0",
			"/ \\",
		}
	}
	for i := 0; i < e.size; i++ {
		for j := 0; j < e.size; j++ {
			c.putUncoloredString(string(view[j][i]), x+i, y+j)
		}
	}
	// render dir, safe to remove
	cx, cy := e.getCentralCoord()
	c.putChar('X', cx+e.dirX+bf_x_offset, cy+e.dirY+bf_y_offset)
}

//func (c *consoleIO) renderPlayerBattlefieldUI(xCoord int, b *battlefield) {
//	var lines = []string{
//		fmt.Sprintf("HP: %d/%d", b.player.hitpoints, b.player.getMaxHp()),
//		fmt.Sprintf("1) Prim Wpn: %s", b.player.primaryWeapon.GetName()),
//		"   |x to swap|   ",
//		fmt.Sprintf("2) Scnd Wpn: %s", b.player.secondaryWeapon.GetName()),
//		fmt.Sprintf("3) Itm: %dx %s",
//			b.player.currentConsumable.AsConsumable.Amount,
//			b.player.currentConsumable.GetName()),
//		"",
//		"ENEMIES:",
//	}
//	enemiesLinesStart := len(lines)
//	for i := range b.enemies {
//		lines = append(lines, fmt.Sprintf("  %s (%s)",
//			b.enemies[i].getName(),
//			getAttackDescriptionString(b.player.primaryWeapon, b.enemies[i]),
//		))
//	}
//	for i := range lines {
//		c.putColorTaggedString(lines[i], xCoord, i)
//	}
//	// render enemies for those enemy lines
//	for i, e := range b.enemies {
//		c.renderEnemyAtCoords(e, b.currentTick, xCoord, enemiesLinesStart+i)
//	}
//}
//
//func (c *consoleIO) getCharForEnemy(heads int) rune {
//	if heads < 10 {
//		return rune(strconv.Itoa(heads)[0])
//	} else if heads < 16 {
//		return []rune{'A', 'B', 'C', 'D', 'E', 'F'}[heads-10]
//	}
//	return '?'
//}
