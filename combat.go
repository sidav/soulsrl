package main

const (
	TILE_FLOOR = iota
	TILE_WALL

	TICKS_IN_COMBAT_TURN = 10
)

type battlefield struct {
	tiles       [][]int
	units       []*mob
	events      []*event
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

	for i := 0; i < 5*bfW*bfH/100; i++ {
		x, y := rnd.RandInRange(1, bfW-2), rnd.RandInRange(1, bfH-2)
		b.tiles[x][y] = TILE_WALL
	}
	b.units = append(b.units, newMob(bfW/2, 2))
	return b
}

func (b *battlefield) applyEvents() {
	for i := 0; i < len(b.events); i++ {
		if b.currentTick >= b.events[i].tickToOccur {
			b.events[i] = b.events[len(b.events)-1]
			b.events = append(b.events[:len(b.events)-1])
			i -= 1
			continue
		}
	}
}

func (b *battlefield) applyAttackPattern(u *mob, ap *attackPattern, vectorX, vectorY int) {
	tickToOccur := b.currentTick + ap.ticksToPerform
	for _, coord := range ap.relativeCoords {
		rotatedX, rotatedY := getVectorRotatedLikeVector(coord[0], coord[1], vectorX, vectorY)
		b.events = append(b.events, &event{
			tickToOccur: tickToOccur,
			owner:       u,
			x: u.x + rotatedX,
			y: u.y + rotatedY,
		})
	}
}
