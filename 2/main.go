/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

package main

import (
	"fmt"
	"sync"
)

// Функция для расчета квадрата числа
func square(num int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()
	square := num * num
	resultChan <- square
}

func main() {
	// Исходный массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	// Канал для передачи результатов
	resultChan := make(chan int, len(numbers))

	// Группа ожидания для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запуск горутины для каждого числа
	for _, num := range numbers {
		wg.Add(1)
		go square(num, &wg, resultChan)
	}

	// Запуск горутины для закрытия канала после завершения всех вычислений
	go func() { wg.Wait(); close(resultChan) }()

	// Чтение и вывод результатов из канала
	for result := range resultChan {
		fmt.Println(result)
	}
}
