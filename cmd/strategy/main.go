package main

import (
	"fmt"
	"patterns/internal/strategy"
)

func main() {
	data := []int{6, 42, 17, 29, 50, 8, 31, 12, 4, 23}

	// Создаем контекст с начальной стратегией сортировки пузырьком
	context := &strategy.SortContext{Strategy: strategy.BubbleSortStrategy{}}

	// Сортируем массив данных с использованием текущей стратегии (пузырьковой сортировки)
	sortedData := context.Sort(data)
	fmt.Println("BubbleSort:", sortedData)

	// Меняем стратегию на быструю сортировку и сортируем снова
	context.SetStrategy(strategy.QuickSortStrategy{})
	sortedData = context.Sort(data)
	fmt.Println("QuickSort:", sortedData)
}
