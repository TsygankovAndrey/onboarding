// CreateCar ...
package factory

// CreateCar - Конкретный фабричный метод для создания автомобиля
func CreateCar() Transport {
	return Car{}
}
