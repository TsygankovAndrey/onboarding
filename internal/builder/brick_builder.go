// BrickBuilder ...
package builder

// BrickBuilder - реализация Builder для строительства кирпичных зданий
type BrickBuilder struct {
	name string
	age  int
}

// SetName устанавливает имя строителя
func (bb *BrickBuilder) SetName(name string) {
	bb.name = name
}

// SetAge устанавливает возраст строителя
func (bb *BrickBuilder) SetAge(age int) {
	bb.age = age
}

// Build создает кирпичное здание и возвращает его
func (bb *BrickBuilder) Build() Building {
	return Building{
		Builder: bb,
		Name:    "Кирпичный дом",
	}
}
