// Pattern chan_of_responsibility
package chain_of_responsibility

import "fmt"

// NotificationHandler - интерфейс обработчика уведомлений
type NotificationHandler interface {
	Send(message string) bool
	SetNext(handler NotificationHandler)
}

// BaseHandler - базовая структура обработчика
type BaseHandler struct {
	nextHandler NotificationHandler
}

// SetNext устанавливает следующий обработчик в цепочке
func (bh *BaseHandler) SetNext(handler NotificationHandler) {
	bh.nextHandler = handler
}

// EmailHandler - обработчик электронной почты
type EmailHandler struct {
	BaseHandler
}

// Send обрабатывает сообщение по электронной почте, если может; иначе передает следующему обработчику.
func (eh *EmailHandler) Send(message string) bool {
	fmt.Println("Отправка по электронной почте:", message)
	if eh.nextHandler != nil {
		return eh.nextHandler.Send(message)
	}
	return false
}

// SMSHandler - обработчик SMS
type SMSHandler struct {
	BaseHandler
}

// Send обрабатывает SMS, если может; иначе передает следующему обработчику
func (sh *SMSHandler) Send(message string) bool {
	fmt.Println("Отправка SMS:", message)
	if sh.nextHandler != nil {
		return sh.nextHandler.Send(message)
	}
	return false
}

// PushHandler - Обработчик push-уведомлений
type PushHandler struct {
	BaseHandler
}

// Send обрабатывает push-уведомление, если может; иначе передает следующему обработчику
func (ph *PushHandler) Send(message string) bool {
	fmt.Println("Отправка push-уведомления:", message)
	if ph.nextHandler != nil {
		return ph.nextHandler.Send(message)
	}
	return false
}
