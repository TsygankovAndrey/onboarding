// SortingStrategy ...
package strategy

// SortingStrategy - интерфейс для всех конкретных стратегий сортировки
type SortingStrategy interface {
	Sort(data []int) []int
}
