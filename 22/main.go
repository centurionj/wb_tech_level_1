/*
Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Инициализация больших чисел a и b
	a := new(big.Int)
	b := new(big.Int)

	// Установка значений для a и b
	a.SetString("1048577", 10) // 2^20 + 1
	b.SetString("2097153", 10) // 2^21 + 1

	// Создание переменных для хранения результатов операций
	sum := new(big.Int)
	diff := new(big.Int)
	prod := new(big.Int)
	quot := new(big.Int)
	rem := new(big.Int)

	// Выполнение арифметических операций
	sum.Add(a, b)
	diff.Sub(a, b)
	prod.Mul(a, b)
	quot.Div(a, b)
	rem.Mod(a, b)

	// Вывод результатов
	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n", b.String())
	fmt.Printf("a + b = %s\n", sum.String())
	fmt.Printf("a - b = %s\n", diff.String())
	fmt.Printf("a * b = %s\n", prod.String())
	fmt.Printf("a / b = %s\n", quot.String())
	fmt.Printf("a %% b = %s\n", rem.String())
}
