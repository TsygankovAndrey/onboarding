package main

import (
	"fmt"
)

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
	strategy SortingStrategy
}

func (sc *SortContext) SetStrategy(strategy SortingStrategy) {
	sc.strategy = strategy
}

func (sc *SortContext) Sort(data []int) []int {
	return sc.strategy.Sort(data)
}

func main() {
	data := []int{6, 42, 17, 29, 50, 8, 31, 12, 4, 23}

	// Создаем контекст с начальной стратегией сортировки пузырьком
	context := &SortContext{strategy: BubbleSortStrategy{}}

	// Сортируем массив данных с использованием текущей стратегии (пузырьковой сортировки)
	sortedData := context.Sort(data)
	fmt.Println("BubbleSort:", sortedData)

	// Меняем стратегию на быструю сортировку и сортируем снова
	context.SetStrategy(QuickSortStrategy{})
	sortedData = context.Sort(data)
	fmt.Println("QuickSort:", sortedData)
}
