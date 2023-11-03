// Light ...
package command

import "fmt"

// Light - получатель, который выполняет фактические действия
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
