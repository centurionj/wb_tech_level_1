/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

package main

import (
	"fmt"
	"math"
)

// Функция для группировки температур в диапазоны
func groupTemperatures(temperatures []float64, step float64) map[float64][]float64 {
	groups := make(map[float64][]float64)

	for _, temp := range temperatures {
		// Определяем диапазон, к которому принадлежит температура
		rangeStart := math.Floor(temp/step) * step
		groups[rangeStart] = append(groups[rangeStart], temp)
	}

	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10.0

	// Группируем температуры
	grouped := groupTemperatures(temperatures, step)

	// Выводим результат
	for key, val := range grouped {
		fmt.Printf("%.0f: %v\n", key, val)
	}
}
