package data

import (
	"soulsrl/geometry"
	"soulsrl/geometry/line"
)

const ( // how exactly skill coords will be determined
	spTypeNearbyRect = iota
	spTypeRelativeCoordinates
	spTypeTargetRect
	spTypeLine
	spTypeSelf
)

type SkillPattern struct {
	patternApplicationType int
	Name                   string
	RelativeCoords         [][]int

	// for helping ai calculations
	ReachInUnitSizes int
}

func (sp *SkillPattern) GetRelativeCoordsByVector(vx, vy int) [][]int {
	var coords [][]int
	for _, coord := range sp.RelativeCoords {
		rotatedX, rotatedY := geometry.GetVectorRotatedLikeVector(coord[0], coord[1], vx, vy)
		coords = append(coords, []int{rotatedX, rotatedY})
	}
	return coords
}

func (sp *SkillPattern) getScaledRelativeCoordsByVector(vx, vy, size int) [][]int {
	rotatedCoords := sp.GetRelativeCoordsByVector(vx, vy)
	var coords [][]int
	for _, coord := range rotatedCoords {
		squareForThis := geometry.MoveSquareByVector(coord[0], coord[1], 0, 0, size)
		for _, sqcoord := range squareForThis {
			coords = append(coords, sqcoord)
		}
	}
	return coords
}

func (sp *SkillPattern) GetListOfCoordsWhenApplied(actorSize, vx, vy int) [][]int {
	return sp.getScaledRelativeCoordsByVector(vx, vy, actorSize)
}

func (sp *SkillPattern) GetListOfCoordsWhenAppliedAtRect(actorX, actorY, actorSize, targetX, targetY, targetSize int) [][]int {
	actorCenterX, actorCenterY := actorX+actorSize/2, actorY+actorSize/2
	targetCenterX, targetCenterY := targetX+targetSize/2, targetY+targetSize/2
	vx, vy := line.GetNextStepForLine(actorCenterX, actorCenterY, targetCenterX, targetCenterY)
	return sp.getScaledRelativeCoordsByVector(vx, vy, actorSize)
}

const (
	APATTERN_SIMPLE_STRIKE = iota
	APATTERN_RIGHT_SLASH
	APATTERN_SLASH
	APATTERN_BIG_SLASH
	APATTERN_LUNGE
	APATTERN_TWO_SIDES
)

var AttackPatternsTable = map[int]*SkillPattern{
	APATTERN_SIMPLE_STRIKE: {
		patternApplicationType: spTypeRelativeCoordinates,
		Name: "Strike",
		RelativeCoords: [][]int{
			{1, 0},
		},
		ReachInUnitSizes: 1,
	},
	APATTERN_RIGHT_SLASH: {
		patternApplicationType: spTypeRelativeCoordinates,
		Name: "Right Slash",
		RelativeCoords: [][]int{
			{1, 0},
			{1, 1},
		},
		ReachInUnitSizes: 1,
	},
	APATTERN_SLASH: {
		patternApplicationType: spTypeRelativeCoordinates,
		Name: "Full Slash",
		RelativeCoords: [][]int{
			{1, -1},
			{1, 0},
			{1, 1},
		},
		ReachInUnitSizes: 1,
	},
	APATTERN_BIG_SLASH: {
		patternApplicationType: spTypeRelativeCoordinates,
		Name: "Big Slash",
		RelativeCoords: [][]int{
			{0, -1},
			{1, -1},
			{1, 0},
			{1, 1},
			{0, 1},
		},
		ReachInUnitSizes: 1,
	},
	APATTERN_LUNGE: {
		patternApplicationType: spTypeRelativeCoordinates,
		Name: "Lunge",
		RelativeCoords: [][]int{
			{1, 0},
			{2, 0},
		},
		ReachInUnitSizes: 2,
	},
	APATTERN_TWO_SIDES: {
		patternApplicationType: spTypeRelativeCoordinates,
		Name: "Two-side strike",
		RelativeCoords: [][]int{
			{1, 0},
			{-1, 0},
		},
		ReachInUnitSizes: 1,
	},
}
