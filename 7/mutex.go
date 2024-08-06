/*
Реализовать конкурентную запись данных в map.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Создание карты и мьютекса
	data := make(map[int]int)
	var mu sync.Mutex

	// Горутина для записи в карту
	write := func(start int, end int) {
		for i := start; i < end; i++ {
			mu.Lock()
			data[i] = i * i
			mu.Unlock()
			time.Sleep(10 * time.Millisecond) // Имитация работы
		}
	}

	// Запуск нескольких горутин для записи данных
	go write(0, 50)
	go write(50, 100)

	// Подождем немного, чтобы горутины успели завершиться
	time.Sleep(1 * time.Second)

	// Выводим содержимое карты
	mu.Lock()
	for k, v := range data {
		fmt.Printf("%d: %d\n", k, v)
	}
	mu.Unlock()
}
