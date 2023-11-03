// Airplane ...
package factory

import "fmt"

// Airplane - конкретный продукт - самолет
type Airplane struct{}

func (a Airplane) Move() {
	fmt.Println("Самолет летит в небе")
}
