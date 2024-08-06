/*
Реализовать бинарный поиск встроенными методами языка.
*/

package main

import (
	"fmt"
	"sort"
)

// BinarySearch осуществляет бинарный поиск в отсортированном срезе.
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// Возвращаем -1, если элемент не найден
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 10

	// Сначала сортируем массив, если он не отсортирован
	sort.Ints(arr)

	// Выполняем бинарный поиск
	index := BinarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Элемент %d найден на позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
