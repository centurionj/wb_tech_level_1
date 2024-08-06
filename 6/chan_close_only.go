/*
Реализовать все возможные способы остановки выполнения горутины.
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
	fmt.Println("readData goroutine stopping")
}

func main() {
	dataChan := make(chan int)

	go writeData(dataChan)
	go readData(dataChan)

	time.Sleep(5 * time.Second)
	close(dataChan) // Закрытие канала

	time.Sleep(1 * time.Second)
}
