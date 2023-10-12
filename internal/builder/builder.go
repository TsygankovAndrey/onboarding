// Pattern builder
package builder

// Builder интерфейс определяет методы для построения зданий
type Builder interface {
	SetName(name string)
	SetAge(age int)
	Build() Building
}

// WoodenBuilder - реализация Builder для строительства деревянных зданий
type WoodenBuilder struct {
	name string
	age  int
}

type Building struct {
	Name string
	Builder
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

// Usecase - функция, демонстрирующая использование строителей
func Usecase() {
	woodenBuilder := &WoodenBuilder{}
	woodenBuilder.SetName("Andrey")
	woodenBuilder.SetAge(22)
	woodenBuilding := woodenBuilder.Build()
	woodenBuilding.Build()

	brickBuilder := &BrickBuilder{}
	brickBuilder.SetName("Сергей")
	brickBuilder.SetAge(28)
	brickBuilding := brickBuilder.Build()
	brickBuilding.Build()
}
