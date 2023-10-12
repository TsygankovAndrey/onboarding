// Pattern facade
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
	return &ComputerFacade{
		cpu:       CPU{},
		hardDrive: HardDrive{},
		memory:    Memory{},
	}
}

// Start запускает компьютер, скрывая сложность внутренних компонентов
func (cf *ComputerFacade) Start() {
	fmt.Println("Запуск компьютера")
	cf.cpu.Start()
	cf.memory.Load()
	cf.hardDrive.ReadData()
	fmt.Println("Компьютер готов к использованию")
}

type CPU struct {
}

func (c CPU) Start() {
	fmt.Println("Запуск процессора")
}

type HardDrive struct {
}

func (hd HardDrive) ReadData() {
	fmt.Println("Чтение данных с жесткого диска")
}

type Memory struct {
}

func (m Memory) Load() {
	fmt.Println("Загрузка данных в память")
}
