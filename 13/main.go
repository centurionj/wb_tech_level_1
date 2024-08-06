/*
Поменять местами два числа без создания временной переменной.
*/

package main

import "fmt"

func math(a, b int) {
	// Меняем местами без временной переменной
	a = a + b // a содержит сумму двух чисел.
	b = a - b // b содержит исходное значение a.
	a = a - b // a содержит исходное значение b.

	fmt.Printf("After swap: a = %d, b = %d\n", a, b)
}

func xor(a, b int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("After swap: a = %d, b = %d\n", a, b)
}

func main() {
	a := 5
	b := 10
	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)

	println("MATH")
	math(a, b)

	println("XOR")
	xor(a, b)
}
