// main pattern chan_of_responsibility
package main

import (
	"fmt"

	"patterns/internal/chain_of_responsibility"
)

// Клиентский код
func main() {
	pushHandler := &chain_of_responsibility.PushHandler{}
	smsHandler := &chain_of_responsibility.SMSHandler{}
	emailHandler := &chain_of_responsibility.EmailHandler{}
	emailHandler.SetNext(smsHandler)
	smsHandler.SetNext(pushHandler)
	message := "Важное уведомление!"
	if emailHandler.Send(message) {
		fmt.Println("Уведомление успешно отправлено.")
	} else {
		fmt.Println("Не удалось отправить уведомление.")
	}
}
