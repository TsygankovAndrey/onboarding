package main

import (
	"fmt"
	"patterns/internal/visitor"
)

func main() {
	// Экземпляр посетителя
	areaVisitor := &visitor.AreaVisitor{}

	// Создаем коллекцию геометрических фигур и добавляем круги и прямоугольники
	shapes := &visitor.ShapesCollection{}
	shapes.AddShape(&visitor.Circle{Radius: 5.0})
	shapes.AddShape(&visitor.Rectangle{Width: 3.0, Height: 4.0})
	shapes.AddShape(&visitor.Circle{Radius: 2.0})

	// Применяем посетителя к коллекции фигур
	shapes.Accept(areaVisitor)

	// Выводим общую площадь всех фигур
	fmt.Printf("Общая площадь: %.2f\n", areaVisitor.TotalArea)
}
