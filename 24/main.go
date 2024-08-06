/*
Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.
*/

package main

import (
	"fmt"
	"math"
)

// Определение структуры Point
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Метод для нахождения расстояния до другой точки
func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки
	p1 := NewPoint(3, 4)
	p2 := NewPoint(2, 3)

	// Вычисляем расстояние между ними
	distance := p1.Distance(p2)

	// Выводим результат
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
