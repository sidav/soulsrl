package data

import (
	"soulsrl/geometry"
)

type AttackPattern struct {
	RelativeCoords       [][]int
	durationInTurnTenths int
}

func (ap *AttackPattern) GetDurationForTurnTicks(ticks int) int {
	return ticks * ap.durationInTurnTenths / 10
}

func (ap *AttackPattern) GetRelativeCoordsByVector(vx, vy int) [][]int {
	var coords [][]int
	for _, coord := range ap.RelativeCoords {
		rotatedX, rotatedY := geometry.GetVectorRotatedLikeVector(coord[0], coord[1], vx, vy)
		coords = append(coords, []int{rotatedX, rotatedY})
	}
	return coords
}

func (ap *AttackPattern) GetScaledRelativeCoordsByVector(vx, vy, size int) [][]int {
	rotatedCoords := ap.GetRelativeCoordsByVector(vx, vy)
	var coords [][]int
	for _, coord := range rotatedCoords {
		squareForThis := geometry.MoveSquareByVector(coord[0], coord[1], 0, 0, size)
		for _, sqcoord := range squareForThis {
			coords = append(coords, sqcoord)
		}
	}
	return coords
}

const (
	APATTERN_SIMPLE_STRIKE = iota
	APATTERN_RIGHT_SLASH
	APATTERN_SLASH
	APATTERN_BIG_SLASH
	APATTERN_LUNGE
	APATTERN_TWO_SIDES
)

var PatternsTable = map[int]*AttackPattern{
	APATTERN_SIMPLE_STRIKE: {
		RelativeCoords: [][]int{
			{1, 0},
		},
		durationInTurnTenths: 10,
	},
	APATTERN_RIGHT_SLASH: {
		RelativeCoords: [][]int{
			{1, 0},
			{1, 1},
		},
		durationInTurnTenths: 10,
	},
	APATTERN_SLASH: {
		RelativeCoords: [][]int{
			{1, -1},
			{1, 0},
			{1, 1},
		},
		durationInTurnTenths: 20,
	},
	APATTERN_BIG_SLASH: {
		RelativeCoords: [][]int{
			{0, -1},
			{1, -1},
			{1, 0},
			{1, 1},
			{0, 1},
		},
		durationInTurnTenths: 30,
	},
	APATTERN_LUNGE: {
		RelativeCoords: [][]int{
			{1, 0},
			{2, 0},
		},
		durationInTurnTenths: 20,
	},
	APATTERN_TWO_SIDES: {
		RelativeCoords: [][]int{
			{1, 0},
			{-1, 0},
		},
		durationInTurnTenths: 20,
	},
}
