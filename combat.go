package main

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
	bfW := 25 //  rnd.RandInRange(7, 10)
	bfH := 15  // rnd.RandInRange(7, 10)
	b.tiles = make([][]int, bfW)
	for i := range b.tiles {
		b.tiles[i] = make([]int, bfH)
	}

	b.mobs = append(b.mobs, newMob(0, 0))
	b.mobs = append(b.mobs, newMob(bfW-1, bfH-1))
	b.mobs[0].size = 3

	totalWalls := bfW*bfH * 10/100
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
	return rectContainsCoords(0, 0, len(b.tiles), len(b.tiles[0]), x, y)
}

func (b *battlefield) isRectFullyPassable(x, y, w int) bool {
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

func (b *battlefield) getMobInSquare(x, y, w int) *mob {
	// TODO: CHECK WHY IT DOESN'T WORK
	for _, m := range b.mobs {
		if doTwoSquaresOverlap(x, y, w, m.x, m.y, m.size) {
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

func (b *battlefield) applyAttackPattern(u *mob, ap *attackPattern, vectorX, vectorY int) {
	tickToOccur := b.currentTick + ap.ticksToPerform
	patternCoords := ap.getScaledRelativeCoordsByVector(vectorX, vectorY, u.size)
	for _, coord := range patternCoords {
		b.actions = append(b.actions, &action{
			tickToOccur: tickToOccur,
			owner:       u,
			x:           u.x + coord[0],
			y:           u.y + coord[1],
		})
	}
}
