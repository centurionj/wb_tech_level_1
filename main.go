/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

package main

import (
	"fmt"
	"strings"
)

// Функция для проверки уникальности символов в строке
func allUniqueChars(s string) bool {
	// Приводим строку к нижнему регистру для регистронезависимой проверки
	lowerStr := strings.ToLower(s)

	// Создаем map для отслеживания встречающихся символов
	charMap := make(map[rune]bool)

	// Проходим по всем символам строки
	for _, char := range lowerStr {
		// Если символ уже есть в map, значит символ не уникален
		if charMap[char] {
			return false
		}
		// Добавляем символ в map
		charMap[char] = true
	}

	// Если не нашли повторяющихся символов, возвращаем true
	return true
}

func main() {
	// Примеры для проверки
	testStrings := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, str := range testStrings {
		fmt.Printf("%s: %v\n", str, allUniqueChars(str))
	}
}
