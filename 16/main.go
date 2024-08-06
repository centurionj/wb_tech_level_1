/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

package main

import (
	"fmt"
)

func quicksort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// Выбираем опорный элемент
	pivot := arr[len(arr)/2]
	left := 0
	right := len(arr) - 1

	for left <= right {
		// Найти элементы, которые нужно обменять
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}

		// Если найдено что обменивать, обменять элементы
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	// Рекурсивно сортируем подмассивы
	if right > 0 {
		quicksort(arr[:right+1])
	}
	if left < len(arr) {
		quicksort(arr[left:])
	}
}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, -1, 0, -100}
	fmt.Println("Original array:", arr)
	quicksort(arr)
	fmt.Println("Sorted array:", arr)
}
