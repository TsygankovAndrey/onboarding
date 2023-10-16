// Rectangle ...
package visitor

// Rectangle - конкретный элемент - прямоугольник
type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Accept(visitor Visitor) {
	visitor.VisitRectangle(r)
}
