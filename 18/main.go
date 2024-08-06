/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
*/

package main

import (
	"fmt"
	"sync"
)

// Counter представляет счетчик с синхронизацией.
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment увеличивает значение счетчика на 1.
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// GetValue возвращает текущее значение счетчика.
func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	// Запускаем несколько горутин, которые инкрементируют счетчик.
	numGoroutines := 100
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	// Ждем завершения всех горутин.
	wg.Wait()

	// Выводим итоговое значение счетчика.
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.GetValue())
}
