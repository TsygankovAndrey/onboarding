package main

import (
	"fmt"
)

// Интерфейс Посетителя
type Visitor interface {
	VisitCircle(circle *Circle)
	VisitRectangle(rectangle *Rectangle)
}

// Конкретный Посетитель
type AreaVisitor struct {
	TotalArea float64
}

func (av *AreaVisitor) VisitCircle(circle *Circle) {
	area := 3.14 * circle.Radius * circle.Radius
	av.TotalArea += area
}

// Интерфейс Элемента
type Shape interface {
	Accept(visitor Visitor)
}

// Конкретные Элементы
type Circle struct {
	Radius float64
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

func (av *AreaVisitor) VisitRectangle(rectangle *Rectangle) {
	area := rectangle.Width * rectangle.Height
	av.TotalArea += area
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Accept(visitor Visitor) {
	visitor.VisitRectangle(r)
}

// Структура объектов
type ShapesCollection struct {
	Shapes []Shape
}

func (sc *ShapesCollection) AddShape(shape Shape) {
	sc.Shapes = append(sc.Shapes, shape)
}

func (sc *ShapesCollection) Accept(visitor Visitor) {
	for _, shape := range sc.Shapes {
		shape.Accept(visitor)
	}
}

func main() {
	// Экземпляр посетителя
	areaVisitor := &AreaVisitor{}

	// Создаем коллекцию геометрических фигур и добавляем круги и прямоугольники
	shapes := &ShapesCollection{}
	shapes.AddShape(&Circle{Radius: 5.0})
	shapes.AddShape(&Rectangle{Width: 3.0, Height: 4.0})
	shapes.AddShape(&Circle{Radius: 2.0})

	// Применяем посетителя к коллекции фигур
	shapes.Accept(areaVisitor)

	// Выводим общую площадь всех фигур
	fmt.Printf("Общая площадь: %.2f\n", areaVisitor.TotalArea)
}
