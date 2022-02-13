package main

type mob struct {
	x, y, dirX, dirY int
	nextTickToAct    int
	ap               *attackPattern
}

func newMob(x, y int) *mob {
	u := &mob{
		x: x,
		y: y,
		dirX: 1,
		dirY: 0,
	}

	u.ap = &attackPattern{relativeCoords:
	[][]int{
		{1, -1},
		{1, 0},
		{1, 1},
		{2, -1},
		{2, 0},
		{2, 1},
	},
		ticksToPerform: 10,
	}

	return u
}
