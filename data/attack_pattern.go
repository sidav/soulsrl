package data

import (
	"soulsrl/geometry"
)

type AttackPattern struct {
	Name                  string
	RelativeCoords        [][]int
	durationInTurnLengths int

	// for helping ai calculations
	ReachInUnitSizes int
}

func (ap *AttackPattern) GetDurationForTurnTicks(ticks int) int {
	return ticks * ap.durationInTurnLengths / 10
}

func (ap *AttackPattern) GetRelativeCoordsByVector(vx, vy int) [][]int {
	var coords [][]int
	for _, coord := range ap.RelativeCoords {
		rotatedX, rotatedY := geometry.GetVectorRotatedLikeVector(coord[0], coord[1], vx, vy)
		coords = append(coords, []int{rotatedX, rotatedY})
	}
	return coords
}

func (ap *AttackPattern) getScaledRelativeCoordsByVector(vx, vy, size int) [][]int {
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

func (ap *AttackPattern) GetListOfCoordsWhenApplied(actorSize, vx, vy int) [][]int {
	return ap.getScaledRelativeCoordsByVector(vx, vy, actorSize)
}

const (
	APATTERN_SIMPLE_STRIKE = iota
	APATTERN_RIGHT_SLASH
	APATTERN_SLASH
	APATTERN_BIG_SLASH
	APATTERN_LUNGE
	APATTERN_TWO_SIDES
)

var AttackPatternsTable = map[int]*AttackPattern{
	APATTERN_SIMPLE_STRIKE: {
		Name: "Strike",
		RelativeCoords: [][]int{
			{1, 0},
		},
		ReachInUnitSizes:      1,
		durationInTurnLengths: 10,
	},
	APATTERN_RIGHT_SLASH: {
		Name: "Right Slash",
		RelativeCoords: [][]int{
			{1, 0},
			{1, 1},
		},
		ReachInUnitSizes:      1,
		durationInTurnLengths: 10,
	},
	APATTERN_SLASH: {
		Name: "Full Slash",
		RelativeCoords: [][]int{
			{1, -1},
			{1, 0},
			{1, 1},
		},
		ReachInUnitSizes:      1,
		durationInTurnLengths: 20,
	},
	APATTERN_BIG_SLASH: {
		Name: "Big Slash",
		RelativeCoords: [][]int{
			{0, -1},
			{1, -1},
			{1, 0},
			{1, 1},
			{0, 1},
		},
		ReachInUnitSizes:      1,
		durationInTurnLengths: 30,
	},
	APATTERN_LUNGE: {
		Name: "Lunge",
		RelativeCoords: [][]int{
			{1, 0},
			{2, 0},
		},
		ReachInUnitSizes:      2,
		durationInTurnLengths: 20,
	},
	APATTERN_TWO_SIDES: {
		Name: "Two-side strike",
		RelativeCoords: [][]int{
			{1, 0},
			{-1, 0},
		},
		ReachInUnitSizes:      1,
		durationInTurnLengths: 20,
	},
}
