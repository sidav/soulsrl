package main

type attackPattern struct {
	relativeCoords [][]int
	ticksToPerform int
}

func (ap *attackPattern) getRelativeCoordsByVector(vx, vy int) [][]int {
	var coords [][]int
	for _, coord := range ap.relativeCoords {
		rotatedX, rotatedY := getVectorRotatedLikeVector(coord[0], coord[1], vx, vy)
		coords = append(coords, []int{rotatedX, rotatedY})
	}
	return coords
}

func (ap *attackPattern) getScaledRelativeCoordsByVector(vx, vy, size int) [][]int {
	rotatedCoords := ap.getRelativeCoordsByVector(vx, vy)
	var coords [][]int
	for _, coord := range rotatedCoords {
		squareForThis := getSquareByVectorFromSquareOfSameSize(coord[0], coord[1], 0, 0, size)
		for _, sqcoord := range squareForThis {
			coords = append(coords, sqcoord)
		}
	}
	return coords
}

const (
	APATTERN_SLASH = iota
	APATTERN_LUNGE
)

var patternsTable = map[int]*attackPattern {
	APATTERN_SLASH: {
		relativeCoords: [][]int{
			{1, -1},
			{1, 0},
			{1, 1},
		},
		ticksToPerform: 10,
	},
	APATTERN_LUNGE: {
		relativeCoords: [][]int{
			{1, 0},
			{2, 0},
		},
		ticksToPerform: 10,
	},
}
