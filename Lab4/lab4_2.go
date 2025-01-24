package main

import (
	"fmt"
	"os"
)

type NameCard struct {
	name []string
	age  []int
}

func main() {
	var card NameCard
	var value int
	fmt.Print("Введите количество записей в карте: ")
	fmt.Fscan(os.Stdin, &value)
	card.age = make([]int, value)
	card.name = make([]string, value)
	for i := 0; i < value; i++ {
		var name string
		var age int
		fmt.Print("Введите имя человека:")
		fmt.Fscan(os.Stdin, &name)
		fmt.Print("Введите его возраст: ")
		fmt.Fscan(os.Stdin, &age)
		card.age[i] = age
		card.name[i] = name
	}
	aver := AverAgeCard(card)
	fmt.Println("Средний возраст людей в карте: ", aver)
}

func AverAgeCard(cart NameCard) float32 {
	var dlina int
	var summa = 0
	dlina = len(cart.age)
	for i := 0; i < dlina; i++ {
		summa = summa + cart.age[i]
	}
	return float32(summa) / float32(dlina)
}
