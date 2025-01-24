package main

import (
	"fmt"
	"os"
)

type Shape interface {
	float32() float32
}

type Rectangle struct {
	length float32
	width  float32
}

func (a Rectangle) float32() float32 {
	Ploshad := a.length * a.width
	return Ploshad
}

type Circle struct {
	radius float32
}

func (b Circle) float32() float32 {
	const pi float32 = 3.1415
	return b.radius * b.radius * pi
}

func Area(s Shape) {
	fmt.Println("Площадь фигуры: ", s.float32())
}

func main() {
	var circle Circle
	var dlina float32
	var shirina float32
	fmt.Print("Введите длину прямоугольника: ")
	fmt.Fscan(os.Stdin, &dlina)
	fmt.Print("Введите ширину прямоугольника: ")
	fmt.Fscan(os.Stdin, &shirina)
	rectangle := Rectangle{dlina, shirina}
	Area(rectangle)
	fmt.Print("Введите радиус круга: ")
	fmt.Fscan(os.Stdin, &circle.radius)
	Area(circle)
}
