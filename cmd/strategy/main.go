// main pattern strategy
package main

import (
	"fmt"

	"patterns/internal/strategy"
)

func main() {
	// Создаем контекст сортировки
	context := &strategy.SortContext{}

	// Данные, которые нужно отсортировать
	data := []int{5, 1, 4, 2, 8}

	// Устанавливаем стратегию сортировки пузырьком
	context.SetStrategy(&strategy.BubbleSortStrategy{})

	// Используем контекст для сортировки данных
	sortedData := context.Sort(data)
	fmt.Println("Отсортированные данные (пузырьком):", sortedData)

	// Теперь меняем стратегию на QuickSortStrategy
	context.SetStrategy(&strategy.QuickSortStrategy{})

	// Используем контекст для сортировки данных с новой стратегией
	sortedData = context.Sort(data)
	fmt.Println("Отсортированные данные (быстрая сортировка):", sortedData)
}
