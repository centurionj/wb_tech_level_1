/*
Реализовать все возможные способы остановки выполнения горутины.
*/

package main

import (
	"context"
	"fmt"
	"time"
)

// Функция для записи данных в канал
func writeData(ctx context.Context, dataChan chan int) {
	i := 0
	for {
		select {
		case dataChan <- i:
			i++
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			fmt.Println("writeData goroutine stopping")
			return
		}
	}
}

// Функция для чтения данных из канала
func readData(ctx context.Context, dataChan chan int) {
	for {
		select {
		case data := <-dataChan:
			fmt.Println("Received:", data)
		case <-ctx.Done():
			fmt.Println("readData goroutine stopping")
			return
		}
	}
}

func main() {
	dataChan := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go writeData(ctx, dataChan)
	go readData(ctx, dataChan)

	time.Sleep(5 * time.Second)
	cancel() // Сигнализация завершения

	time.Sleep(1 * time.Second)
}
