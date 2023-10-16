// CPU ...
package facade

import "fmt"

// CPU - компонент - процессор
type CPU struct{}

func (c CPU) Start() {
	fmt.Println("Запуск процессора")
}
