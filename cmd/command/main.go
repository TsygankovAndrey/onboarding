// main pattern command
package main

import "patterns/internal/command"

func main() {
	// Свет и команды для включения и выключения
	light := &command.Light{}
	turnOnCommand := &command.TurnOnLightCommand{Light: light}
	turnOffCommand := &command.TurnOffLightCommand{Light: light}

	// Создаем пульт и настраиваем его кнопки
	remote := &command.RemoteControl{}

	// Нажимаем кнопку для включения света
	remote.Command = turnOnCommand
	remote.PressButton()

	// Нажимаем кнопку для выключения света
	remote.Command = turnOffCommand
	remote.PressButton()
}
