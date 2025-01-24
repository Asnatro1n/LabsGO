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
	var index int
	var imya string
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
	fmt.Print("Введите имя, которое требуется удалить с карты: ")
	fmt.Fscan(os.Stdin, &imya)
	for i := 0; i < value; i++ {
		if card.name[i] == imya {
			index = i
		}
	}
	card.name = append(card.name[0:index], card.name[index+1:]...)
	card.age = append(card.age[0:index], card.age[index+1:]...)
	for i := 0; i < len(card.age); i++ {
		fmt.Println(card.name[i], " - ", card.age[i])
	}
}
