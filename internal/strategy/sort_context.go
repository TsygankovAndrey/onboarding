// SortContext ...
package strategy

// SortContext - класс, который использует выбранную стратегию для сортировки
type SortContext struct {
	Strategy SortingStrategy
}

// SetStrategy - устанавливает стратегию сортировки
func (sc *SortContext) SetStrategy(strategy SortingStrategy) {
	sc.Strategy = strategy
}

// Sort - выполняет сортировку с использованием текущей стратегии
func (sc *SortContext) Sort(data []int) []int {
	return sc.Strategy.Sort(data)
}
