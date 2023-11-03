// AreaVisitor ...
package visitor

// AreaVisitor - конкретный посетитель для вычисления площади
type AreaVisitor struct {
	TotalArea float64
}

func (av *AreaVisitor) VisitCircle(circle *Circle) {
	area := 3.14 * circle.Radius * circle.Radius
	av.TotalArea += area
}

func (av *AreaVisitor) VisitRectangle(rectangle *Rectangle) {
	area := rectangle.Width * rectangle.Height
	av.TotalArea += area
}
