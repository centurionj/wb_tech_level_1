/*
Реализовать все возможные способы остановки выполнения горутины.
*/

package main

import (
	"fmt"
	"time"
)

// Функция для записи данных в канал
func writeData(dataChan chan int, stopChan chan struct{}) {
	i := 0
	for {
		select {
		case dataChan <- i:
			i++
			time.Sleep(1 * time.Second)
		case <-stopChan:
			fmt.Println("writeData goroutine stopping")
			return
		}
	}
}

// Функция для чтения данных из канала
func readData(dataChan chan int, stopChan chan struct{}) {
	for {
		select {
		case data := <-dataChan:
			fmt.Println("Received:", data)
		case <-stopChan:
			fmt.Println("readData goroutine stopping")
			return
		}
	}
}

func main() {
	dataChan := make(chan int)
	stopChan := make(chan struct{})

	go writeData(dataChan, stopChan)
	go readData(dataChan, stopChan)

	time.Sleep(5 * time.Second)
	close(stopChan) // Сигнализация завершения

	time.Sleep(1 * time.Second)
}
