/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/

package main

import (
	"fmt"
	"reflect"
)

// Функция для определения типа переменной
func getType(v interface{}) string {
	// Используем reflect.TypeOf для получения типа переменной
	t := reflect.TypeOf(v)
	return t.String()
}

func main() {
	// Примеры переменных разных типов
	var intVar int = 42
	var strVar string = "hello"
	var boolVar bool = true
	var chVar chan int = make(chan int)

	// Примеры переменных типа interface{}
	var iVar interface{}

	// Установка переменной типа interface{} в разные значения
	iVar = intVar
	fmt.Printf("Type of iVar: %s\n", getType(iVar))

	iVar = strVar
	fmt.Printf("Type of iVar: %s\n", getType(iVar))

	iVar = boolVar
	fmt.Printf("Type of iVar: %s\n", getType(iVar))

	iVar = chVar
	fmt.Printf("Type of iVar: %s\n", getType(iVar))
}
