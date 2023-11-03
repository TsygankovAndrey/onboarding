// HardDrive ...
package facade

import "fmt"

// HardDrive - компонент - жесткий диск
type HardDrive struct{}

func (hd HardDrive) ReadData() {
	fmt.Println("Чтение данных с жесткого диска")
}
