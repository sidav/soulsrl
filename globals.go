package main

func readKeyToVector(key string) (int, int) {
	switch key {
	case "q":
		return -1, -1
	case "UP", "w":
		return 0, -1
	case "e":
		return 1, -1
	case "RIGHT", "d":
		return 1, 0
	case "c":
		return 1, 1
	case "DOWN", "x":
		return 0, 1
	case "z":
		return -1, 1
	case "LEFT", "a":
		return -1, 0
	}
	return 0, 0
}
