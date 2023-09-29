package factory

import "fmt"

// Продукт
type Transport interface {
	Move()
}

// Конкретные продукты
type Car struct{}
type Airplane struct{}

func (c Car) Move() {
	fmt.Println("Автомобиль едет по дороге")
}

func (a Airplane) Move() {
	fmt.Println("Самолет летит в небе")
}

// Фабричный метод
type TransportFactory func() Transport

// Конкретные фабричные методы
func CreateCar() Transport {
	return Car{}
}

func CreateAirplane() Transport {
	return Airplane{}
}
