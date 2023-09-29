package main

import (
	"fmt"
)

// Интерфейс обработчика
type NotificationHandler interface {
	Send(message string) bool
	SetNext(handler NotificationHandler)
}

// Конкретные обработчики
type EmailHandler struct {
	nextHandler NotificationHandler
}

func (eh *EmailHandler) Send(message string) bool {
	fmt.Println("Отправка по электронной почте:", message)
	return true
}

func (eh *EmailHandler) SetNext(handler NotificationHandler) {
	eh.nextHandler = handler
}

type SMSHandler struct {
	nextHandler NotificationHandler
}

func (sh *SMSHandler) Send(message string) bool {
	fmt.Println("Отправка SMS:", message)
	return true
}

func (sh *SMSHandler) SetNext(handler NotificationHandler) {
	sh.nextHandler = handler
}

type PushHandler struct {
	nextHandler NotificationHandler
}

func (ph *PushHandler) Send(message string) bool {
	fmt.Println("Отправка push-уведомления:", message)
	return true
}

func (ph *PushHandler) SetNext(handler NotificationHandler) {
	ph.nextHandler = handler
}

// Клиентский код
func main() {
	// Цепочка обработчиков: Email -> SMS -> Push
	emailHandler := &EmailHandler{}
	smsHandler := &SMSHandler{}
	pushHandler := &PushHandler{}

	emailHandler.SetNext(smsHandler)
	smsHandler.SetNext(pushHandler)

	// Отправляем уведомление и передаем его по цепочке обработчиков
	message := "Важное уведомление!"
	if emailHandler.Send(message) {
		fmt.Println("Уведомление успешно отправлено.")
	} else {
		fmt.Println("Не удалось отправить уведомление.")
	}
}
