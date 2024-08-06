/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

package main

import (
	"fmt"
	"strings"
)

// Функция для переворачивания слов в строке
func reverseWords(s string) string {
	// Разбиваем строку на слова по пробелам
	words := strings.Fields(s)

	// Переворачиваем срез слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Объединяем слова обратно в строку с пробелами
	return strings.Join(words, " ")
}

func main() {
	// Пример строки
	input := "snow dog sun"

	// Переворачиваем слова в строке
	reversed := reverseWords(input)

	// Выводим результат
	fmt.Println("Исходная строка:", input)
	fmt.Println("Строка с перевернутыми словами:", reversed)
}
