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

func (m *mob) containsCoords(x, y int) bool {
	return m.x <= x && x < m.x+m.size && m.y <= y && y < m.y+m.size
}

// rewrite, too hardcoded
func newMob(x, y int) *mob {
	u := &mob{
		x:    x,
		y:    y,
		size: 1,
		dirX: 1,
		dirY: 0,
	}

	u.ap = patternsTable[APATTERN_SLASH]

	return u
}
