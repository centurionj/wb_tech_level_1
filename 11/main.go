/*
Реализовать пересечение двух неупорядоченных множеств.
*/

package main

import (
	"fmt"
)

// Функция для нахождения пересечения двух множеств
func intersect(set1, set2 map[int]struct{}) map[int]struct{} {
	// Результирующее множество для пересечения
	intersection := make(map[int]struct{})

	// Проходим по первому множеству и проверяем наличие элементов во втором множестве
	for elem := range set1 {
		if _, exists := set2[elem]; exists {
			intersection[elem] = struct{}{}
		}
	}

	return intersection
}

func main() {
	// Инициализация двух множеств
	set1 := map[int]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
	}

	set2 := map[int]struct{}{
		4: {},
		5: {},
		6: {},
		7: {},
	}

	// Нахождение пересечения множеств
	result := intersect(set1, set2)

	// Вывод пересечения
	fmt.Println("Intersection:", result)
}
