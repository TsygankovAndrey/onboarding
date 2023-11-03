// Sender ...
package chain_of_responsibility

// Sender - интерфейс для отправки уведомлений
type Sender interface {
	Send(message string) bool
}
