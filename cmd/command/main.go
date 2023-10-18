// main pattern command
package main

import "patterns/internal/command"

func main() {
	light := &command.Light{}
	turnOnCommand := &command.TurnOnLightCommand{Light: light}
	turnOffCommand := &command.TurnOffLightCommand{Light: light}
	remote := &command.RemoteControl{}
	remote.Command = turnOnCommand
	remote.PressButton()
	remote.Command = turnOffCommand
	remote.PressButton()
}
