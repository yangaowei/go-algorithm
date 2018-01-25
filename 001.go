package main

import (
	"fmt"
)

func binary_search(a []int, search int) (index int) {
	index = -1
	low := 0
	height := len(a) - 1
	for low < height {
		mid := (low + height) / 2
		if a[mid] < search {
			low = mid + 1
		} else if a[mid] > search {
			height = mid - 1
		} else {
			index = mid
			break
		}
	}
	return
}

func main() {

	arr := []int{3, 5, 7, 9, 10, 46, 333}

	fmt.Println(binary_search(arr, 9))
	fmt.Println(binary_search(arr, 111))
}
