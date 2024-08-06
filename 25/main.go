/*
Реализовать собственную функцию sleep.
*/

package main

import (
	"fmt"
	"time"
)

func mySleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Начало сна")
	start := time.Now()
	mySleep(2 * time.Second)
	end := time.Now()
	fmt.Printf("Конец сна через %v\n", end.Sub(start))
}
