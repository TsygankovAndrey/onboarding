// WoodenBuilder ...
package builder

// WoodenBuilder - реализация Builder для строительства деревянных зданий
type WoodenBuilder struct {
	name string
	age  int
}

// SetName устанавливает имя строителя
func (wb *WoodenBuilder) SetName(name string) {
	wb.name = name
}

// SetAge устанавливает возраст строителя
func (wb *WoodenBuilder) SetAge(age int) {
	wb.age = age
}

// Build создает деревянное здание и возвращает его
func (wb *WoodenBuilder) Build() Building {
	return Building{
		Builder: wb,
		Name:    "Деревянная изба",
	}
}
