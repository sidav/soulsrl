package geometry

import "soulsrl/geometry/line"

func SquareContainsCoords(sx, sy, w, x, y int) bool {
	return RectContainsCoords(sx, sy, w, w, x, y)
}

func RectToCoords(x, y, w, h int) [][]int {
	var coords [][]int
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			coords = append(coords, []int{i, j})
		}
	}
	return coords
}

func RectContainsCoords(rx, ry, w, h, x, y int) bool {
	return rx <= x && x < rx+w && ry <= y && y < ry+h
}

func DoTwoSquaresOverlap(x1, y1, w1, x2, y2, w2 int) bool {
	return DoTwoCellRectsOverlap(x1, y1, w1, w1, x2, y2, w2, w2)
}

func DoTwoCellRectsOverlap(x1, y1, w1, h1, x2, y2, w2, h2 int) bool {
	// WARNING:
	// ALL "-1"s HERE ARE BECAUSE OF WE ARE IN CELLS SPACE
	// I.E. A SINGLE CELL IS 1x1 RECTANGLE
	// SO RECTS (0, 0, 1x1) AND (1, 0, 1x1) ARE NOT OVERLAPPING IN THIS SPACE (BUT SHOULD IN EUCLIDEAN OF COURSE)
	right1 := x1 + w1 - 1
	bottom1 := y1 + h1 - 1
	right2 := x2 + w2 - 1
	bottom2 := y2 + h2 - 1
	return !(x2 > right1 ||
		right2 < x1 ||
		y2 > bottom1 ||
		bottom2 < y1)
}

func DistanceBetweenSquares(x1, y1, w1, x2, y2, w2 int) int {
	return DistanceBetweenRectangles(x1, y1, w1, w1, x2, y2, w2, w2)
}

func GetCoordsOfClosestCoordToRectFromRect(x1, y1, w1, h1, x2, y2, w2, h2 int) (int, int) {
	linePoints := line.GetLine(x1+w1/2, y1+h1/2, x2+w2/2, y2+h2/2)
	for _, point := range linePoints {
		if RectContainsCoords(x2, y2, w2, h2, point.X, point.Y) {
			return point.X, point.Y
		}
	}
	panic("Y U NO CONTAIN")
}

func DistanceBetweenRectangles(x1, y1, w1, h1, x2, y2, w2, h2 int) int {
	w1--
	h1--
	w2--
	h2--
	left := x2+w2 < x1
	right := x1+w1 < x2
	bottom := y2+h2 < y1
	top := y1+h1 < y2
	if top && left {
		return OrthogonalDistance(x1, y1+h1, x2+w2, y2)
	}
	if left && bottom {
		return OrthogonalDistance(x1, y1, x2+w2, y2+h2)
	}
	if bottom && right {
		return OrthogonalDistance(x1+w1, y1, x2, y2+h2)
	}
	if right && top {
		return OrthogonalDistance(x1+w1, y1+h1, x2, y2)
	}
	if left {
		return x1 - x2 - w2
	}
	if right {
		return x2 - x1 - w1
	}
	if bottom {
		return y1 - y2 - h2
	}
	if top {
		return y2 - y1 - h1
	}
	return 0 // intersect
}

func MoveSquareByVector(vx, vy, topleftx, toplefty, size int) [][]int {
	var rectCoords [][]int
	newtopleftx := topleftx + vx*size
	newtoplefty := toplefty + vy*size
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			rectCoords = append(rectCoords, []int{newtopleftx + x, newtoplefty + y})
		}
	}
	return rectCoords
}

// returns true as first result if found anything
func FindCoordsForNeighbouringSquareOfSameSizeContainingCoords(sqx, sqy, size, contX, contY int) (bool, int, int) {
	if SquareContainsCoords(sqx, sqy, size, contX, contY) {
		return false, 0, 0
	}
	if DistanceBetweenSquares(sqx, sqy, size, contX, contY, size) > size {
		return false, 0, 0
	}
	var foundX, foundY, foundDistanceToCenter int
	foundDistanceToCenter = 999 // foundDist should be 1 anyway

	// stupid algorithm of brute force
	for tryx := contX -(size-1); tryx <= contX; tryx++ {
		for tryy := contY -(size-1); tryy <= contY; tryy++ {
			// fmt.Printf("Trying %d, %d\n", tryx, tryy)
			distBetwSquares := DistanceBetweenSquares(sqx, sqy, size, tryx, tryy, size)
			if distBetwSquares != 1 {
				// fmt.Printf(" Dist is %d, skipping...\n", distBetwSquares)
				continue
			}
			centerX, centerY := tryx+size/2, tryy+size/2
			dist := sqDistance(contX, contY, centerX, centerY)
			// fmt.Printf(" Center is at %d,%d; comparing %d with %d...\n", centerX, centerY, dist, foundDistanceToCenter)
			if dist < foundDistanceToCenter {
				// fmt.Printf("   ...distance is picked.\n")
				foundX = tryx
				foundY = tryy
				foundDistanceToCenter = dist
			}
		}
	}
	return true, foundX, foundY
}
