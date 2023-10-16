// SMSHandler ...
package chain_of_responsibility

import "fmt"

// SMSHandler - обработчик SMS
type SMSHandler struct {
	nextHandler Sender
}

// SetNext устанавливает следующий обработчик в цепочке
func (sh *SMSHandler) SetNext(handler Sender) {
	sh.nextHandler = handler
}

// Send обрабатывает SMS, если может; иначе передает следующему обработчику
func (sh *SMSHandler) Send(message string) bool {
	fmt.Println("Отправка SMS:", message)
	if sh.nextHandler != nil {
		return sh.nextHandler.Send(message)
	}
	return true // Это обработчик по умолчанию, который всегда завершает отправку успешно
}
