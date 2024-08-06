/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(2^2+3^2+4^2...) с использованием конкурентных вычислений.
*/

package main

import (
	"fmt"
	"sync"
)

// Функция для расчета квадрата числа и добавления его к общей сумме
func squareAndSum(num int, wg *sync.WaitGroup, sum *int, mu *sync.Mutex) {
	defer wg.Done()
	square := num * num
	mu.Lock()
	*sum += square
	mu.Unlock()
}

func main() {
	// Исходный массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	// Переменная для хранения суммы
	sum := 0

	// Мьютекс для защиты доступа к переменной суммы
	var mu sync.Mutex

	// Группа ожидания для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запуск горутины для каждого числа
	for _, num := range numbers {
		wg.Add(1)
		go squareAndSum(num, &wg, &sum, &mu)
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод результата
	fmt.Println("Sum of squares:", sum)
}
