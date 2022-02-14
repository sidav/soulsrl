package main

import (
	"fmt"
	"github.com/sidav/sidavgorandom/fibrandom"
	"math"
	"soulsrl/game_log"
)

var (
	io  consoleIO
	rnd *fibrandom.FibRandom
	log *game_log.GameLog
	exitGame bool
)

func main() {
	x, y := 2, 0
	fx, fy := float64(x), float64(y)
	for i := 0; i <= 8; i++ {
		fmt.Printf("Rotated: %d, %d == %.1f, %.1f\n", x, y, fx, fy)
		// x, y = rotateIntVector(x, y, 45)
		x, y = stupidRotateVector45(x, y)
		ft := fx
		fx = fx * math.Cos(math.Pi/4) - fy * math.Sin(math.Pi/4) + 0.1
		fy = ft * math.Sin(math.Pi/4) + fy * math.Cos(math.Pi/4) + 0.1
	}
	io.init()
	log = &game_log.GameLog{}
	log.Init(3)
	rnd = &fibrandom.FibRandom{}
	rnd.InitDefault()
	defer io.close()

	b := newBattlefield()
	for !exitGame {
		b.combatGameLoop()
	}
}
