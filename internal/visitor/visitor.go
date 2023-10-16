// Visitor ...
package visitor

// Visitor - интерфейс посетителя
type Visitor interface {
	VisitCircle(circle *Circle)
	VisitRectangle(rectangle *Rectangle)
}

// Shape - интерфейс элемента
type Shape interface {
	Accept(visitor Visitor)
}
