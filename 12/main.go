/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
*/

package main

import (
	"fmt"
)

func main() {
	// Имеется последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество
	set := make(map[string]struct{})

	// Добавляем строки в множество
	for _, str := range strings {
		set[str] = struct{}{}
	}

	// Выводим множество
	fmt.Println("Unique strings in the set:")
	for key := range set {
		fmt.Println(key)
	}
}
