// PushHandler ...
package chain_of_responsibility

import "fmt"

// PushHandler - Обработчик push-уведомлений
type PushHandler struct {
	nextHandler Sender
}

// SetNext устанавливает следующий обработчик в цепочке
func (ph *PushHandler) SetNext(handler Sender) {
	ph.nextHandler = handler
}

// Send обрабатывает push-уведомление, если может; иначе передает следующему обработчику
func (ph *PushHandler) Send(message string) bool {
	fmt.Println("Отправка push-уведомления:", message)
	if ph.nextHandler != nil {
		return ph.nextHandler.Send(message)
	}
	return true // Это обработчик по умолчанию, который всегда завершает отправку успешно
}
