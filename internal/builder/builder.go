// Builder ...
package builder

// Builder интерфейс определяет методы для построения зданий
type Builder interface {
	SetName(name string)
	SetAge(age int)
	Build() Building
}
