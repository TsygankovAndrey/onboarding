// Memory ...
package facade

import "fmt"

// Memory - компонент - оперативная память
type Memory struct{}

func (m Memory) Load() {
	fmt.Println("Загрузка данных в память")
}
