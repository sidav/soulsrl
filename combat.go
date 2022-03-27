package main

import (
	"soulsrl/geometry"
	"soulsrl/geometry/line"
)

const (
	TILE_FLOOR = iota
	TILE_WALL

	TICKS_IN_COMBAT_TURN = 2
)

type battlefield struct {
	tiles       [][]int
	player      *mob
	mobs        []*mob
	actions     []*action
	currentTick int
}

func newBattlefield() *battlefield {
	b := &battlefield{}
	bfW := rnd.RandInRange(3, 12)*2 + 1
	bfH := rnd.RandInRange(3, 7)*2 + 1
	b.tiles = make([][]int, bfW)
	for i := range b.tiles {
		b.tiles[i] = make([]int, bfH)
	}

	//b.mobs = append(b.mobs, newMob("Giant"))
	//b.mobs = append(b.mobs, newMob("Swordmaster"))
	//b.mobs[1].x, b.mobs[1].y = bfW-1, bfH-1
	// fmt.Printf("Distance is %d\n", geometry.DistanceBetweenSquares(b.mobs[0].x, b.mobs[0].y, b.mobs[0].size, b.mobs[1].x, b.mobs[1].y, b.mobs[1].size))

	totalWalls := bfW * bfH * 7 / 100
	for i := 0; i < totalWalls; {
		x, y := rnd.RandInRange(0, bfW-1), rnd.RandInRange(0, bfH-1)
		if b.getMobPresentAt(x, y) == nil {
			b.tiles[x][y] = TILE_WALL
			i++
		}
	}

	return b
}

func (b *battlefield) addMobAtRandomEmptyPlace(m *mob) {
	size := m.size
	var goodCoords [][]int
	for x := 0; x < len(b.tiles)-size; x++ {
		for y := 0; y < len(b.tiles[0])-size; y++ {
			if b.areAllTilesInRectPassable(x, y, size) && b.getMobInSquareOtherThan(x, y, size, nil) == nil {
				goodCoords = append(goodCoords, []int{x, y})
			}
		}
	}
	if len(goodCoords) > 0 {
		ind := rnd.Rand(len(goodCoords))
		m.x = goodCoords[ind][0]
		m.y = goodCoords[ind][1]
		b.mobs = append(b.mobs, m)
	}
}

func (b *battlefield) getVectorForVisibleFromMobToMobIfExists(from, to *mob) (int, int) {
	fx, fy := from.getCentralCoord()
	tx, ty := to.getCentralCoord()
	lineToCheck := line.GetLine(fx, fy, tx, ty)
	for _, l := range lineToCheck {
		if b.tiles[l.X][l.Y] == TILE_WALL {
			return 0, 0
		}
	}
	return lineToCheck[1].X-lineToCheck[0].X, lineToCheck[1].Y - lineToCheck[0].Y
}

func (b *battlefield) containsCoords(x, y int) bool {
	return geometry.RectContainsCoords(0, 0, len(b.tiles), len(b.tiles[0]), x, y)
}

func (b *battlefield) areAllTilesInRectPassable(x, y, w int) bool {
	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			cx, cy := x+i, y+j
			if !b.containsCoords(cx, cy) {
				return false
			}
			if b.tiles[cx][cy] != TILE_FLOOR {
				return false
			}
		}
	}
	return true
}

func (b *battlefield) getMobPresentAt(x, y int) *mob {
	for _, m := range b.mobs {
		if m.containsCoords(x, y) {
			return m
		}
	}
	return nil
}

func (b *battlefield) getMobInSquareOtherThan(x, y, w int, otherThan *mob) *mob {
	for _, m := range b.mobs {
		if m == otherThan {
			continue
		}
		if geometry.DoTwoSquaresOverlap(x, y, w, m.x, m.y, m.size) {
			return m
		}
	}
	return nil
}

func (b *battlefield) getActionPresentAt(x, y int) *action {
	for _, a := range b.actions {
		if a.x == x && a.y == y {
			return a
		}
	}
	return nil
}

func (b *battlefield) getListOfVectorsToPassableCoordsForMob(m *mob) [][]int {
	var coords [][]int
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if b.areAllTilesInRectPassable(m.x+x, m.y+y, m.size) {
				coords = append(coords, []int{x, y})
			}
		}
	}
	return coords
}

func (b *battlefield) tryMoveMobByVector(m *mob, vx, vy int, spendTime bool) bool {
	mobAtCoords := b.getMobInSquareOtherThan(m.x+vx, m.y+vy, m.size, m)
	if mobAtCoords == nil && b.areAllTilesInRectPassable(m.x+vx, m.y+vy, m.size) {
		m.x += vx
		m.y += vy
		if spendTime {
			m.nextTickToAct = b.currentTick + TICKS_IN_COMBAT_TURN
		}
		return true
	}
	return false
}

func (b *battlefield) tryRollMobByVector(m *mob, vx, vy int) bool {
	rolled := false
	if m.stamina < 3 {
		return false
	}
	for i := 1; i <= 2; i++ {
		mobAtCoords := b.getMobInSquareOtherThan(m.x+vx, m.y+vy, m.size, m)
		if  mobAtCoords == nil && b.areAllTilesInRectPassable(m.x+vx, m.y+vy, m.size) {
			m.x += vx
			m.y += vy
			rolled = true
		}
	}
	if rolled {
		m.stamina -= 3
		m.nextTickToAct = b.currentTick + TICKS_IN_COMBAT_TURN
	}
	return rolled
}
