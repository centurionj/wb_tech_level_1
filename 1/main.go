/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/

package main

import (
	"fmt"
)

// Определение структуры Human
type Human struct {
	Name string
	Age  int
}

// Метод структуры Human
func (h Human) Introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", h.Name, h.Age)
}

// Метод структуры Human
func (h Human) Learn() {
	fmt.Println("I'm learning GO!")
}

// Определение структуры Action, которая встраивает Human
type Action struct {
	Human
	ActionName string
}

// Метод структуры Action
func (a Action) DoAction() {
	fmt.Printf("%s is performing action: %s\n", a.Name, a.ActionName) // можно вызвать поле Name вот так a.Human.Name
}

func main() {
	// Создание экземпляра структуры Action
	action := Action{
		Human: Human{
			Name: "Mike",
			Age:  21,
		},
		ActionName: "Play games",
	}

	// Вызов метода из структуры Human через экземпляр Action
	action.Introduce()
	action.Learn()

	// Вызов метода структуры Action
	action.DoAction()
}
