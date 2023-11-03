// main pattern visitor
package main

import (
	"fmt"

	"patterns/internal/visitor"
)

func main() {
	areaVisitor := &visitor.AreaVisitor{}
	shapes := &visitor.ShapesCollection{}
	shapes.AddShape(&visitor.Circle{Radius: 5.0})
	shapes.AddShape(&visitor.Rectangle{Width: 3.0, Height: 4.0})
	shapes.AddShape(&visitor.Circle{Radius: 2.0})
	shapes.Accept(areaVisitor)
	fmt.Printf("Общая площадь: %.2f\n", areaVisitor.TotalArea)
}
