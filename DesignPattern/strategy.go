package main

import (
	"../DesignPattern/src/strategy"
)

func main() {
	arr := []int{4, 1, 5, 10, 3, 6}
	calculator := strategy.NewCalculator()
	bubble_sort := strategy.NewBubbleSort(arr)
	calculator.ChangeAlg(bubble_sort)
	calculator.Cal() // I am using util bubble sort
	quick_sort := strategy.NewQuickSort(arr)
	calculator.ChangeAlg(quick_sort)
	calculator.Cal() // I am using util quick sort
}
