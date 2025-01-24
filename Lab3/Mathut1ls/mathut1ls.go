package mathutils

func factorial(x int) int {
	var factorial int = 1
	a := x
	proverka := x
	for x > 1 {
		x--
		factorial *= (a * x)
		a -= 2
		x--
	}
	if proverka > 0 {
		return factorial
	} else {
		return 0
	}
}
