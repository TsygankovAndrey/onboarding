// RemoteControl ...
package command

// RemoteControl - отправитель, который создает и выполняет команды
type RemoteControl struct {
	Command Command
}

func (r *RemoteControl) PressButton() {
	r.Command.Execute()
}
