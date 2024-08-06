/*
Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.
*/

package main

import (
	"fmt"
)

// Устанавливает i-й бит в 1
func setBit(value int64, i int) int64 {
	return value | (1 << i)
}

// Устанавливает i-й бит в 0
func clearBit(value int64, i int) int64 {
	return value & ^(1 << i)
}

// Проверяет значение i-го бита
func getBit(value int64, i int) int {
	return int((value >> i) & 1)
}

func main() {
	var value int64 = 0 // Начальное значение переменной

	fmt.Println("Initial value:", value)

	i := 5 // Установка 5-го бита

	// Устанавливаем i-й бит в 1
	value = setBit(value, i)
	fmt.Printf("Value after setting %d-th bit to 1: %d\n", i, value)

	// Устанавливаем i-й бит в 0
	value = clearBit(value, i)
	fmt.Printf("Value after clearing %d-th bit: %d\n", i, value)

	// Проверяем значение i-го бита
	bitValue := getBit(value, i)
	fmt.Printf("Value of %d-th bit: %d\n", i, bitValue)
}
