// Pattern factory
package factory

import "fmt"

// Transport - продукт
type Transport interface {
	Move()
}

// Car, Airplane - конкретные продукты
type Car struct{}
type Airplane struct{}

func (c Car) Move() {
	fmt.Println("Автомобиль едет по дороге")
}

func (a Airplane) Move() {
	fmt.Println("Самолет летит в небе")
}

// TransportFactory - фабричный метод
type TransportFactory func() Transport

// CreateCar, CreateAirplane - Конкретные фабричные методы
func CreateCar() Transport {
	return Car{}
}

func CreateAirplane() Transport {
	return Airplane{}
}
