/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/

package main

import (
	"fmt"
)

func reverseString(s string) string {
	// Преобразуем строку в срез рун (символов Unicode)
	runes := []rune(s)

	// Инициализируем указатели на начало и конец среза рун
	left := 0
	right := len(runes) - 1

	// Переворачиваем срез рун
	for left < right {
		// Меняем местами символы
		runes[left], runes[right] = runes[right], runes[left]
		// Двигаем указатели
		left++
		right--
	}

	// Преобразуем срез рун обратно в строку
	return string(runes)
}

func main() {
	// Пример строки
	input := "главрыба"

	// Переворачиваем строку
	reversed := reverseString(input)

	// Выводим результат
	fmt.Println("Исходная строка:", input)
	fmt.Println("Перевернутая строка:", reversed)
}
