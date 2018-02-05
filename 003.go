package main

import (
	"fmt"
)

func merge(a []int) []int {
	size := len(a)
	if size <= 1 {
		return a
	}
	mid := size / 2
	left := merge(a[:mid])
	right := merge(a[mid:])
	fmt.Println(left, right)
	return mergeSort(left, right)
}

func mergeSort(a, b []int) (result []int) {
	lsize := len(a)
	right := len(b)

	var l, r int
	for l < lsize && r < right {
		if a[l] > b[r] {
			result = append(result, a[l])
			l += 1
		} else {
			result = append(result, b[r])
			r += 1
		}
	}
	result = append(result, a[l:]...)
	result = append(result, b[r:]...)
	return
}
func main() {
	arr := []int{3, 1, 5, 7, 2, 4, 9, 6, 10, 8, 20}
	fmt.Println(merge(arr))
}
