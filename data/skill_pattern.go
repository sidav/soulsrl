package data

import (
	"soulsrl/geometry"
	"soulsrl/geometry/line"
)

const ( // how exactly skill coords will be determined
	spTypeNearbyRect = iota
	spTypeRelativeCoordinates
	spTypeTargetRectAnywhere
	spTypeLine
	spTypeSelf
)

type SkillPattern struct {
	patternApplicationType int
	Name                   string
	AttackRelativeCoords   [][]int
	MovementRelativeCoords [][]int

	// for helping ai calculations, used ONLY in them
	ReachInUnitSizes int
}

type PatternMovementVector struct {
	Vx, Vy int
	JumpOver bool
}

func getRelativeCoordsByVector(coordsFacingRight [][]int, vx, vy int) [][]int {
	var coords [][]int
	for _, coord := range coordsFacingRight {
		rotatedX, rotatedY := geometry.GetVectorRotatedLikeVector(coord[0], coord[1], vx, vy)
		coords = append(coords, []int{rotatedX, rotatedY})
	}
	return coords
}

func getScaledRelativeCoordsByVector(coordsFacingRight [][]int, vx, vy, size int) [][]int {
	rotatedCoords := getRelativeCoordsByVector(coordsFacingRight, vx, vy)
	var coords [][]int
	for _, coord := range rotatedCoords {
		squareForThis := geometry.MoveSquareByVector(coord[0], coord[1], 0, 0, size)
		for _, sqcoord := range squareForThis {
			coords = append(coords, sqcoord)
		}
	}
	return coords
}

func (sp *SkillPattern) getListOfCoordsWhenApplied(actorSize, vx, vy int) [][]int {
	return getScaledRelativeCoordsByVector(sp.AttackRelativeCoords, vx, vy, actorSize)
}

func (sp *SkillPattern) GetListOfCoordsWhenAppliedAtRect(actorX, actorY, actorSize, targetX, targetY, targetSize int) ([][]int, [][]int) {
	actorCenterX, actorCenterY := actorX+actorSize/2, actorY+actorSize/2
	targetCenterX, targetCenterY := targetX+targetSize/2, targetY+targetSize/2
	vx, vy := line.GetNextStepForLine(actorCenterX, actorCenterY, targetCenterX, targetCenterY)
	var attackCoords[][]int
	switch sp.patternApplicationType {
	case spTypeNearbyRect:
		coverX, coverY := geometry.GetCoordsOfClosestCoordToRectFromRect(actorX, actorY, actorSize, actorSize,
			targetX, targetY, targetSize, targetSize)
		found, x, y := geometry.FindCoordsForNeighbouringSquareOfSameSizeContainingCoords(actorX, actorY, actorSize,
			coverX, coverY)
		if !found {
			panic("Not found... Y U NO FOUND?")
		}
		attackCoords = geometry.RectToCoords(x, y, actorSize, actorSize)
	case spTypeRelativeCoordinates:
		coords := getScaledRelativeCoordsByVector(sp.AttackRelativeCoords, vx, vy, actorSize)
		for _, c := range coords {
			c[0] += actorX
			c[1] += actorY
		}
		attackCoords = coords
	default:
		panic("Y U NO IMPLEMENT")
	}
	// now collect movement VECTORS (not coordinates)
	moveVectors := make([][]int, 0)
	if len(sp.MovementRelativeCoords) > 0 {
		movCoords := getRelativeCoordsByVector(sp.MovementRelativeCoords, vx, vy)
		moveLength := actorSize
		for i := 0; i < moveLength; i++ {
			moveVectors = append(moveVectors, movCoords[0])
		}
	}
	return attackCoords, moveVectors
}

const (
	APATTERN_SIMPLE_STRIKE = iota

	APATTERN_STRIKE_STEP_BACK

	APATTERN_RIGHT_SLASH
	APATTERN_SLASH
	APATTERN_BIG_SLASH

	APATTERN_LUNGE
	APATTERN_JUMP_LUNGE

	APATTERN_TWO_SIDES
)

var AttackPatternsTable = map[int]*SkillPattern{
	APATTERN_SIMPLE_STRIKE: {
		patternApplicationType: spTypeNearbyRect,
		Name:                   "Strike",
		AttackRelativeCoords: [][]int{
			{1, 0},
		},
		ReachInUnitSizes: 1,
	},
	APATTERN_STRIKE_STEP_BACK: {
		patternApplicationType: spTypeNearbyRect,
		Name:                   "Strike and step back",
		AttackRelativeCoords: [][]int{
			{1, 0},
		},
		MovementRelativeCoords: [][]int{
			{-1, 0},
		},
		ReachInUnitSizes: 1,
	},
	APATTERN_RIGHT_SLASH: {
		patternApplicationType: spTypeRelativeCoordinates,
		Name:                   "Right Slash",
		AttackRelativeCoords: [][]int{
			{1, 0},
			{1, 1},
		},
		ReachInUnitSizes: 1,
	},
	APATTERN_SLASH: {
		patternApplicationType: spTypeRelativeCoordinates,
		Name:                   "Full Slash",
		AttackRelativeCoords: [][]int{
			{1, -1},
			{1, 0},
			{1, 1},
		},
		ReachInUnitSizes: 1,
	},
	APATTERN_BIG_SLASH: {
		patternApplicationType: spTypeRelativeCoordinates,
		Name:                   "Big Slash",
		AttackRelativeCoords: [][]int{
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
		Name:                   "Lunge",
		AttackRelativeCoords: [][]int{
			{1, 0},
			{2, 0},
		},
		ReachInUnitSizes: 2,
	},
	APATTERN_JUMP_LUNGE: {
		patternApplicationType: spTypeRelativeCoordinates,
		Name:                   "Jump Lunge",
		AttackRelativeCoords: [][]int{
			{2, 0},
		},
		MovementRelativeCoords: [][]int{
			{1, 0},
		},
		ReachInUnitSizes: 2,
	},
	APATTERN_TWO_SIDES: {
		patternApplicationType: spTypeRelativeCoordinates,
		Name:                   "Two-side strike",
		AttackRelativeCoords: [][]int{
			{1, 0},
			{-1, 0},
		},
		ReachInUnitSizes: 1,
	},
}
