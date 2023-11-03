// QuickSortStrategy ...
package strategy

// QuickSortStrategy - конкретная стратегия быстрой сортировки
type QuickSortStrategy struct{}

// Sort - реализация быстрой сортировки
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
