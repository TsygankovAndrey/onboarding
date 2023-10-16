// Circle ...
package visitor

// Circle - конкретный элемент - круг
type Circle struct {
	Radius float64
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}
