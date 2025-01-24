package main

import (
	"fmt"
	"os"
)

type Circle struct {
	radius float32
}

func main() {
	var kryg Circle
	fmt.Print("Введите радиус круга: ")
	fmt.Fscan(os.Stdin, &kryg.radius)
	SquareCircle(kryg)
}

func SquareCircle(a Circle) {
	const pi float32 = 3.1415
	S := a.radius * a.radius * pi
	fmt.Println("Площадь данного круга: ", S)
}
