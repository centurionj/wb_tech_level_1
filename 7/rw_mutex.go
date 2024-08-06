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
	// Создание карты и RWMutex
	data := make(map[int]int)
	var rwmu sync.RWMutex

	// Горутина для записи в карту
	write := func(start int, end int) {
		for i := start; i < end; i++ {
			rwmu.Lock()
			data[i] = i * i
			rwmu.Unlock()
			time.Sleep(10 * time.Millisecond) // Имитация работы
		}
	}

	// Горутина для чтения из карты
	read := func() {
		for {
			rwmu.RLock()
			for k, v := range data {
				fmt.Printf("%d: %d\n", k, v)
			}
			rwmu.RUnlock()
			time.Sleep(50 * time.Millisecond) // Имитация работы
		}
	}

	// Запуск нескольких горутин для записи данных
	go write(0, 50)
	go write(50, 100)

	// Запуск горутины для чтения данных
	go read()

	// Подождем немного, чтобы горутины успели завершиться
	time.Sleep(3 * time.Second)
}
