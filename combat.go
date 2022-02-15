package main

import (
	"soulsrl/data"
	"soulsrl/geometry"
)

const (
	TILE_FLOOR = iota
	TILE_WALL

	TICKS_IN_COMBAT_TURN = 2
)

type battlefield struct {
	tiles       [][]int
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

	b.mobs = append(b.mobs, newMob("Giant", 0, 0))
	b.mobs = append(b.mobs, newMob("Swordmaster", bfW-1, bfH-1))
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

func (b *battlefield) applyAttackPattern(acting *mob, ap *data.AttackPattern, vectorX, vectorY int) {
	tickToOccur := b.currentTick + ap.GetDurationForTurnTicks(TICKS_IN_COMBAT_TURN)
	patternCoords := ap.GetScaledRelativeCoordsByVector(vectorX, vectorY, acting.size)
	for _, coord := range patternCoords {
		b.actions = append(b.actions, &action{
			tickToOccur: tickToOccur,
			owner:       acting,
			x:           acting.x + coord[0],
			y:           acting.y + coord[1],
		})
	}
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
