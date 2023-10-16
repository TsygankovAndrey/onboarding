// CreateAirplane ...
package factory

// CreateAirplane - Конкретный фабричный метод для создания самолета
func CreateAirplane() Transport {
	return Airplane{}
}
