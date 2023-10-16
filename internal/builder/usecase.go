// Usecase ...
package builder

// Usecase - функция, демонстрирующая использование строителей
func Usecase() {
	woodenBuilder := &WoodenBuilder{}
	woodenBuilder.SetName("Andrey")
	woodenBuilder.SetAge(22)
	woodenBuilding := woodenBuilder.Build()
	woodenBuilding.Builder.Build()

	brickBuilder := &BrickBuilder{}
	brickBuilder.SetName("Сергей")
	brickBuilder.SetAge(28)
	brickBuilding := brickBuilder.Build()
	brickBuilding.Builder.Build()
}
