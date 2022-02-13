package main

import "math"

func getVectorRotatedLikeVector(vx, vy, rx, ry int) (int, int) {
	// rotates vx, vy relative to (1, 0) until this (1, 0) is (rx, ry)
	//x, y := 1, 0
	//for x != rx || y != ry {
	//	x, y = rotateIntVector(x, y, 45)
	//	vx, vy = rotateIntVector(vx, vy, 45)
	//}
	atan := math.Atan2(float64(ry), float64(rx))
	return rotateIntVector(vx, vy, int(atan*180/3.14159265))
}

func rotateIntVector(x, y, angle int) (int, int) {
	t := x
	rads := float64(angle) * 3.14159265358979323 / 180
	sin := math.Sin(rads)
	cos := math.Cos(rads)
	fx := float64(x)*cos - float64(y)*sin
	fy := float64(t)*sin + float64(y)*cos
	x = int(math.Round(fx))
	y = int(math.Round(fy))
	return x, y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
