// Car ...
package factory

import "fmt"

// Car - конкретный продукт - автомобиль
type Car struct{}

func (c Car) Move() {
	fmt.Println("Автомобиль едет по дороге")
}
