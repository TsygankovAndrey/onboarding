// main pattern chan_of_responsibility
package main

import (
	"fmt"

	"patterns/internal/chain_of_responsibility"
)

// Клиентский код
func main() {
	// Цепочка обработчиков: Email -> SMS -> Push
	emailHandler := &chain_of_responsibility.EmailHandler{}
	smsHandler := &chain_of_responsibility.SMSHandler{}
	pushHandler := &chain_of_responsibility.PushHandler{}

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
