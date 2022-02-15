package main

import (
	"fmt"
	"github.com/sidav/sidavgorandom/fibrandom"
	"soulsrl/game_log"
	"soulsrl/geometry"
)

var (
	io       consoleIO
	rnd      *fibrandom.FibRandom
	log      *game_log.GameLog
	exitGame bool
)

func main() {
	doTestingStuff()

	log = &game_log.GameLog{}
	log.Init(3)
	rnd = &fibrandom.FibRandom{}
	rnd.InitDefault()

	b := newBattlefield()

	io.init()
	defer io.close()

	for !exitGame {
		b.combatGameLoop()
	}
}

func doTestingStuff() {
	//x, y := 2, 0
	//fx, fy := float64(x), float64(y)
	//for i := 0; i <= 8; i++ {
	//	fmt.Printf("Rotated: %d, %d == %.1f, %.1f\n", x, y, fx, fy)
	//	// x, y = rotateIntVector(x, y, 45)
	//	x, y = geometry.StupidRotateVector45(x, y)
	//	ft := fx
	//	fx = fx*math.Cos(math.Pi/4) - fy*math.Sin(math.Pi/4) + 0.1
	//	fy = ft*math.Sin(math.Pi/4) + fy*math.Cos(math.Pi/4) + 0.1
	//}
	x, y := 0, 0
	size := 3
	//for atx := 0; atx < 10; atx++ {
	//	for aty := 0; aty < 10; aty++ {
	//		fmt.Printf("Dist to %d, %d: %d\n", atx, aty, geometry.DistanceBetweenSquares(x, y, size, atx, aty, size))
	//	}
	//}
	atx, aty := 3, 3
	_, foundx, foundy := geometry.FindCoordsForNeighbouringSquareOfSameSizeContainingCoords(x, y, size, atx, aty)
	fmt.Printf("Closest: %d, %d\n", foundx, foundy)
}
