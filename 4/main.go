/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте. Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.
*/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Функция для воркера, который читает данные из канала и выводит их
func worker(id int, dataChan <-chan int, doneChan <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case data := <-dataChan:
			fmt.Printf("Worker %d received data: %d\n", id, data)
		case <-doneChan:
			fmt.Printf("Worker %d shutting down\n", id)
			return
		}
	}
}

func main() {
	// Определение количества воркеров
	numWorkers := 5

	// Канал для передачи данных
	dataChan := make(chan int)

	// Канал для уведомления о завершении работы
	doneChan := make(chan struct{})

	// Группа ожидания для воркеров
	var wg sync.WaitGroup

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChan, doneChan, &wg)
	}

	// Канал для получения системных сигналов
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Горутина для записи данных в канал
	go func() {
		for i := 1; ; i++ {
			select {
			case <-doneChan:
				return
			default:
				dataChan <- i
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Ожидание сигнала завершения
	sig := <-sigChan
	fmt.Printf("Received signal: %s. Shutting down...\n", sig)

	// Закрытие канала doneChan для уведомления воркеров о завершении
	close(doneChan)

	// Ожидание завершения всех воркеров
	wg.Wait()

	fmt.Println("All workers have been shut down.")
}
