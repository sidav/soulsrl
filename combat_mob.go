package main

type mob struct {
	// topleft coord
	x, y, size    int
	dirX, dirY    int
	nextTickToAct int
	ap            *attackPattern
}

func (m *mob) getCentralCoord() (int, int) {
	if m.size == 0 || m.size == 1 {
		return m.x, m.y
	}
	return m.x + m.size/2, m.y + m.size/2
}

func newMob(x, y int) *mob {
	u := &mob{
		x:    x,
		y:    y,
		size: 1,
		dirX: 1,
		dirY: 0,
	}

	u.ap = &attackPattern{relativeCoords:
	[][]int{
		{1, -1},
		{1, 0},
		{1, 1},
	},
		ticksToPerform: 10,
	}

	return u
}
