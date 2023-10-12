// main pattern factory
package main

import "patterns/internal/factory"

func main() {
	// Используем фабричные методы для создания объектов
	car := factory.CreateCar()
	airplane := factory.CreateAirplane()

	// Вызываем методы объектов
	car.Move()
	airplane.Move()
}
