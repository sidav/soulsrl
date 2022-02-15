package geometry

func GetVectorRotatedLikeVector(vx, vy, rx, ry int) (int, int) {
	// rotates vx, vy relative to (1, 0) until this (1, 0) is (rx, ry)
	x, y := 1, 0
	for x != rx || y != ry {
		x, y = StupidRotateVector45(x, y)
		vx, vy = StupidRotateVector45(vx, vy)
	}
	return vx, vy
	//atan := math.Atan2(float64(ry), float64(rx))
	//return rotateIntVector(vx, vy, int(atan*180/3.14159265))
}

//func rotateIntVector(x, y, angle int) (int, int) {
//	rads := float64(angle) * 3.14159265358979323 / 180
//	sin := math.Sin(rads)
//	cos := math.Cos(rads)
//	fx := float64(x)*cos - float64(y)*sin
//	fy := float64(x)*sin + float64(y)*cos
//	x = int(math.Round(fx))
//	y = int(math.Round(fy))
//	return x, y
//}

// doesn't work well when (x != 0 && y != 0 && coords are not diagonal)
func StupidRotateVector45(x, y int) (int, int) {
	initialLen := max(abs(x), abs(y))
	t := x
	x = x - y
	y = t + y
	return sign(x) * initialLen, sign(y) * initialLen
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func sign(x int) int {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return -1
	}
	return 1
}

func OrthogonalDistance(x1, y1, x2, y2 int) int {
	return max(abs(x1-x2), abs(y1-y2))
}

func sqDistance(x1, y1, x2, y2 int) int {
	return abs(x1-x2)*abs(x1-x2) + abs(y1-y2)*abs(y1-y2)
}

func ScaleCoords(x, y, scale int) [][]int {
	var scaled [][]int
	for nx := 0; nx < scale; nx++ {
		for ny := 0; ny < scale; ny++ {
			scaled = append(scaled, []int{x + nx, y + ny})
		}
	}
	return scaled
}
