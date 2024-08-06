/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
*/

package main

import (
	"fmt"
	"time"
)

// Функция для записи данных в канал
func writeData(dataChan chan int) {
	i := 0
	for {
		dataChan <- i
		i++
		time.Sleep(1 * time.Second)
	}
}

// Функция для чтения данных из канала
func readData(dataChan chan int) {
	for data := range dataChan {
		fmt.Println("Received:", data)
	}
}

func main() {
	// Канал для передачи данных
	dataChan := make(chan int)

	// Таймер для завершения программы через N секунд
	N := 5
	timer := time.NewTimer(time.Duration(N) * time.Second)

	// Запуск горутины для записи данных в канал
	go writeData(dataChan)

	// Запуск горутины для чтения данных из канала
	go readData(dataChan)

	// Блокировка до истечения таймера
	<-timer.C

	// Закрытие канала для завершения чтения данных
	close(dataChan)

	fmt.Println("Program finished")
}
