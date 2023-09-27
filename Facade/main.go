package main

import "fmt"

type ComputerFacade struct {
	cpu       CPU
	hardDrive HardDrive
	memory    Memory
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       CPU{},
		hardDrive: HardDrive{},
		memory:    Memory{},
	}
}

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

func main() {
	computer := NewComputerFacade()
	computer.Start()
}
