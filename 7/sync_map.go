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
	var data sync.Map

	// Горутина для записи в sync.Map
	write := func(start int, end int) {
		for i := start; i < end; i++ {
			data.Store(i, i*i)
			time.Sleep(10 * time.Millisecond) // Имитация работы
		}
	}

	// Запуск нескольких горутин для записи данных
	go write(0, 50)
	go write(50, 100)

	// Подождем немного, чтобы горутины успели завершиться
	time.Sleep(1 * time.Second)

	// Выводим содержимое карты
	data.Range(func(key, value interface{}) bool {
		fmt.Printf("%d: %d\n", key.(int), value.(int))
		return true
	})
}
