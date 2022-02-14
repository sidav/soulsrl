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
	bfW := 11 //  rnd.RandInRange(7, 10)
	bfH := 9  // rnd.RandInRange(7, 10)
	b.tiles = make([][]int, bfW)
	for i := range b.tiles {
		b.tiles[i] = make([]int, bfH)
	}

	//for i := 0; i < 5*bfW*bfH/100; i++ {
	//	x, y := rnd.RandInRange(1, bfW-2), rnd.RandInRange(1, bfH-2)
	//	b.tiles[x][y] = TILE_WALL
	//}
	b.mobs = append(b.mobs, newMob(bfW/2, bfH/2))
	return b
}

func (b *battlefield) isRectFullyPassable(x, y, w int) bool {
	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			cx, cy := x+i, y+j
			if cx < 0 || cx >= len(b.tiles) || cy < 0 || cy >= len(b.tiles[cx]) {
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
