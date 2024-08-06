/*
Удалить i-ый элемент из слайса.
*/

package main

import (
	"fmt"
)

func remove(slice []int, i int) []int {
	// Проверяем, чтобы индекс был в пределах допустимого диапазона
	if i < 0 || i >= len(slice) {
		return slice
	}
	// Удаляем элемент путем создания нового слайса
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	// Пример слайса
	slice := []int{0, 1, 2, 3, 4, 5}

	// Индекс элемента, который нужно удалить
	i := 2

	fmt.Printf("Оригинальный слайс: %v\n", slice)

	// Удаление элемента
	newSlice := remove(slice, i)

	fmt.Printf("Слайс после удаления %d-го элемента: %v\n", i, newSlice)
}
