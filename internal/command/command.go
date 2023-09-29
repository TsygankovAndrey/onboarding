package command

import (
	"fmt"
)

type Command interface {
	Execute()
}

// Структуры для конкретных команд
type TurnOnLightCommand struct {
	Light *Light
}

func (c *TurnOnLightCommand) Execute() {
	c.Light.TurnOn()
}

type TurnOffLightCommand struct {
	Light *Light
}

func (c *TurnOffLightCommand) Execute() {
	c.Light.TurnOff()
}

// Получатель, который выполняет фактические действия
type Light struct {
	IsOn bool
}

func (l *Light) TurnOn() {
	l.IsOn = true
	fmt.Println("Свет включен")
}

func (l *Light) TurnOff() {
	l.IsOn = false
	fmt.Println("Свет выключен")
}

// Отправитель, который создает и выполняет команды

type RemoteControl struct {
	Command Command
}

func (r *RemoteControl) PressButton() {
	r.Command.Execute()
}
