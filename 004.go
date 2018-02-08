package main

import (
	"fmt"
)

func insertSort(arr []int) {
	size := len(arr)

	for i := 1; i < size; i++ {
		j := i
		v := arr[i]

		for j > 0 && v < arr[j-1] {
			arr[j] = arr[j-1]
			j -= 1
		}
		arr[j] = v
	}
	fmt.Println(arr)
}

func main() {
	a := []int{3, 1, 5, 7, 2, 4, 9, 6}
	insertSort(a)
}
