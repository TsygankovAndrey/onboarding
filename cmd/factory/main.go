// main pattern factory
package main

import "patterns/internal/factory"

func main() {
	car := factory.CreateCar()
	airplane := factory.CreateAirplane()
	car.Move()
	airplane.Move()
}
