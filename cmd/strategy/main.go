// main pattern strategy
package main

import (
	"fmt"

	"patterns/internal/strategy"
)

func main() {
	context := &strategy.SortContext{}
	data := []int{5, 1, 4, 2, 8}
	context.SetStrategy(&strategy.BubbleSortStrategy{})
	sortedData := context.Sort(data)
	fmt.Println("Отсортированные данные (пузырьком):", sortedData)
	context.SetStrategy(&strategy.QuickSortStrategy{})
	sortedData = context.Sort(data)
	fmt.Println("Отсортированные данные (быстрая сортировка):", sortedData)
}
