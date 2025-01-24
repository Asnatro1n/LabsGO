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
	var newperson string
	var newage int
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
	fmt.Print("Введите имя нового человека:")
	fmt.Fscan(os.Stdin, &newperson)
	fmt.Print("Введите возраст нового человека: ")
	fmt.Fscan(os.Stdin, &newage)
	card.age = append(card.age, newage)
	card.name = append(card.name, newperson)
	for i := 0; i < len(card.age); i++ {
		fmt.Println(card.name[i], " - ", card.age[i])
	}
}
