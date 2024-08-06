/*
Реализовать паттерн «адаптер» на любом примере.
*/

package main

import "fmt"

// Celsius это интерфейс для работы с температурой в градусах Цельсия
type Celsius interface {
	GetTemperatureCelsius() float64
}

// Реализация Celsius
type ConcreteCelsius struct {
	temperature float64
}

func (c *ConcreteCelsius) GetTemperatureCelsius() float64 {
	return c.temperature
}

// Fahrenheit это интерфейс для работы с температурой в градусах Фаренгейта
type Fahrenheit interface {
	GetTemperatureFahrenheit() float64
}

// Реализация Fahrenheit
type ConcreteFahrenheit struct {
	temperature float64
}

func (f *ConcreteFahrenheit) GetTemperatureFahrenheit() float64 {
	return f.temperature
}

// TemperatureAdapter адаптирует Fahrenheit к Celsius
type TemperatureAdapter struct {
	fahrenheit Fahrenheit
}

func (a *TemperatureAdapter) GetTemperatureCelsius() float64 {
	fTemp := a.fahrenheit.GetTemperatureFahrenheit()
	// Преобразование Фаренгейта в Цельсий
	return (fTemp - 32) * 5 / 9
}

func main() {
	// Создаем объект температур в Фаренгейтах
	fahrenheit := &ConcreteFahrenheit{temperature: 100}

	// Создаем адаптер для преобразования Фаренгейтов в Цельсии
	adapter := &TemperatureAdapter{fahrenheit: fahrenheit}

	// Используем адаптер для получения температуры в Цельсиях
	fmt.Printf("Temperature in Celsius: %.2f\n", adapter.GetTemperatureCelsius())
}
