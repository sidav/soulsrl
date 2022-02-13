package main

func getVectorRotatedLikeVector(vx, vy, rx, ry int) (int, int) {
	// rotates vx, vy relative to (1, 0) until this (1, 0) is (rx, ry)
	x, y := 1, 0
	for x != rx || y != ry {
		x, y = stupidRotateVector45(x, y)
		vx, vy = stupidRotateVector45(vx, vy)
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

// doesn't work well when x != 0, y != 0 and coords are not diagonal
func stupidRotateVector45(x, y int) (int, int) {
	initialLen := max(abs(x), abs(y))
	t := x
	x = x - y
	y = t + y
	return sign(x)*initialLen, sign(y)*initialLen
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

func scaleCoords(x, y, scale int) [][]int {
	var scaled [][]int
	for nx := 0; nx < scale; nx++ {
		for ny := 0; ny < scale; ny++ {
			scaled = append(scaled, []int{x + nx, y + ny})
		}
	}
	return scaled
}

func getSquareByVectorFromSquareOfSameSize(vx, vy, topleftx, toplefty, size int) [][]int {
	var rectCoords [][]int
	newtopleftx := topleftx + vx*size
	newtoplefty := toplefty + vy*size
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			rectCoords = append(rectCoords, []int{newtopleftx+x, newtoplefty+y})
		}
	}
	return rectCoords
}
