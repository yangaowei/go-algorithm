package main

import (
	"fmt"
)

func sum(arr []int) int {
	r := 0
	for _, v := range arr {
		r += v
	}
	return r
}

func version01(arr []int) {
	max := 0
	var sub []int
	size := len(arr)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			tmp := sum(arr[i:j])
			if tmp > max {
				max = tmp
				sub = arr[i:j]
			}
		}
	}
	fmt.Println(max, sub)
}

func version02(a []int) int {
	max := 0
	b := 0
	size := len(a)
	for i := 0; i < size; i++ {
		if b < 0 {
			b = a[i]
		} else {
			b += a[i]
		}
		if max < b {
			max = b
		}
	}
	return max
}

func main() {
	arr := []int{1, -2, 3, 10, -4, 7, 2, -5}

	version01(arr)
	fmt.Println(version02([]int{-1, -2, -3}))
}
