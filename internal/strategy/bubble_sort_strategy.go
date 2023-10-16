// BubbleSortStrategy ...
package strategy

// BubbleSortStrategy - конкретная стратегия сортировки пузырьком
type BubbleSortStrategy struct{}

// Sort - реализация сортировки пузырьком
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
