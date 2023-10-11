package builder

type Builder interface {
	SetName(name string)
	SetAge(age int)
	Build() Building
}

type WoodenBuilder struct {
	name string
	age  int
}

type Building struct {
	Name string
	Builder
}

func (wb *WoodenBuilder) SetName(name string) {
	wb.name = name
}

func (wb *WoodenBuilder) SetAge(age int) {
	wb.age = age
}

func (wb *WoodenBuilder) Build() Building {
	return Building{
		Builder: wb,
		Name:    "Деревянная изба",
	}
}

type BrickBuilder struct {
	name string
	age  int
}

func (bb *BrickBuilder) SetName(name string) {
	bb.name = name
}

func (bb *BrickBuilder) SetAge(age int) {
	bb.age = age
}

func (bb *BrickBuilder) Build() Building {
	return Building{
		Builder: bb,
		Name:    "Кирпичный дом",
	}
}

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
