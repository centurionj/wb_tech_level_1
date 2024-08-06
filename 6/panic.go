/*
Реализовать все возможные способы остановки выполнения горутины.
*/

package main

import (
	"fmt"
	"time"
)

// Функция для записи данных в канал с паникой
func writeData(dataChan chan int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("writeData goroutine stopped due to panic:", r)
		}
	}()
	i := 0
	for {
		dataChan <- i
		i++
		time.Sleep(1 * time.Second)
	}
}

// Функция для чтения данных из канала с паникой
func readData(dataChan chan int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("readData goroutine stopped due to panic:", r)
		}
	}()
	for data := range dataChan {
		fmt.Println("Received:", data)
	}
}

func main() {
	dataChan := make(chan int)

	go writeData(dataChan)
	go readData(dataChan)

	time.Sleep(5 * time.Second)
	panic("manual stop") // Вызов паники для завершения всех горутин

	time.Sleep(1 * time.Second)
}
