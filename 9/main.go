/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/

package main

import (
	"fmt"
	"time"
)

// Функция для записи данных в первый канал
func sendData(inputChan chan<- int, numbers []int) {
	defer close(inputChan) // Закрываем канал после отправки всех данных

	for _, num := range numbers {
		inputChan <- num
	}
}

// Функция для обработки данных и отправки в второй канал
func processData(inputChan <-chan int, outputChan chan<- int) {
	defer close(outputChan) // Закрываем канал после обработки всех данных

	for num := range inputChan {
		result := num * 2
		outputChan <- result
	}
}

// Функция для чтения данных из второго канала и вывода на экран
func printData(outputChan <-chan int) {
	for result := range outputChan {
		fmt.Println(result)
	}
}

func main() {
	// Создание двух каналов
	inputChan := make(chan int)
	outputChan := make(chan int)

	// Исходный массив чисел
	numbers := []int{1, 2, 3, 4, 5}

	// Запуск горутины для записи данных в первый канал
	go sendData(inputChan, numbers)

	// Запуск горутины для обработки данных и отправки во второй канал
	go processData(inputChan, outputChan)

	// Запуск горутины для чтения данных из второго канала и вывода на экран
	go printData(outputChan)

	// Ожидание завершения всех горутин
	time.Sleep(1 * time.Second) // Подождем немного, чтобы все данные успели обработаться
}
