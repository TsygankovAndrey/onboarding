// EmailHandler ...
package chain_of_responsibility

import "fmt"

// EmailHandler - обработчик электронной почты
type EmailHandler struct {
	nextHandler Sender
}

// SetNext устанавливает следующий обработчик в цепочке
func (eh *EmailHandler) SetNext(handler Sender) {
	eh.nextHandler = handler
}

// Send обрабатывает сообщение по электронной почте, если может; иначе передает следующему обработчику.
func (eh *EmailHandler) Send(message string) bool {
	fmt.Println("Отправка по электронной почте:", message)
	if eh.nextHandler != nil {
		return eh.nextHandler.Send(message)
	}
	return true // Это обработчик по умолчанию, который всегда завершает отправку успешно
}
