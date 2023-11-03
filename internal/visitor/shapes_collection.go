// ShapesCollection ...
package visitor

// ShapesCollection - структура для хранения коллекции объектов
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
