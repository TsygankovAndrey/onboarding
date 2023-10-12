// Pattern strategy
package strategy

// Стратегия - интерфейс для всех конкретных стратегий
type SortingStrategy interface {
	Sort(data []int) []int
}

// Конкретная стратегия 1 - сортировка пузырьком
type BubbleSortStrategy struct{}

func (bs BubbleSortStrategy) Sort(data []int) []int {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

// Конкретная стратегия 2 - быстрая сортировка
type QuickSortStrategy struct{}

func (qs QuickSortStrategy) Sort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	pivot := data[len(data)-1]
	less := []int{}
	equal := []int{}
	more := []int{}

	for _, item := range data {
		switch {
		case item < pivot:
			less = append(less, item)
		case item == pivot:
			equal = append(equal, item)
		case item > pivot:
			more = append(more, item)
		}
	}

	less = qs.Sort(less)
	more = qs.Sort(more)

	return append(append(less, equal...), more...)
}

// Контекст - класс, который использует выбранную стратегию для сортировки
type SortContext struct {
	Strategy SortingStrategy
}

func (sc *SortContext) SetStrategy(strategy SortingStrategy) {
	sc.Strategy = strategy
}

func (sc *SortContext) Sort(data []int) []int {
	return sc.Strategy.Sort(data)
}
