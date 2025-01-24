package main

import (
	"fmt"
	mathutils "mathutils"
	"os"
)

func main() {
	var chislo int
	fmt.Print("Введите целое число для вычисления его факториала: ")
	fmt.Fscan(os.Stdin, &chislo)
	mathutils.Factorial(chislo)
	factor := mathutils.Factorial(chislo)
	fmt.Println("Факториал данного числа:", factor)
}
