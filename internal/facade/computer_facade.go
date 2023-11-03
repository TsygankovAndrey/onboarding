// ComputerFacade ...
package facade

import "fmt"

// ComputerFacade - фасад компьютера, предоставляющий упрощенный интерфейс для работы с компонентами
type ComputerFacade struct {
	cpu       CPU
	hardDrive HardDrive
	memory    Memory
}

// NewComputerFacade создает новый экземпляр ComputerFacade
func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{}
}

// Start запускает компьютер, скрывая сложность внутренних компонентов
func (cf *ComputerFacade) Start() {
	fmt.Println("Запуск компьютера")
	cf.cpu.Start()
	cf.memory.Load()
	cf.hardDrive.ReadData()
	fmt.Println("Компьютер готов к использованию")
}
