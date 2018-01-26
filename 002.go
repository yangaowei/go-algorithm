package main

import (
	"fmt"
)

func findMax(arr []int, index int) (maxIndex int) {
	min := index
	for j := index; j < len(arr); j++ {
		if arr[j] > arr[min] {
			min = j
		}
	}
	return min
}

func selectSort(arr []int) (result []int) {
	for index, _ := range arr {
		i := findMax(arr, index)
		if index != i {
			arr[index], arr[i] = arr[i], arr[index]
		}
		fmt.Println(i)
	}
	return arr
}

func main() {
	a := []int{3, 1, 5, 7, 2, 4, 9, 6}
	fmt.Println(selectSort(a))
}
