package main

import (
	"fmt"
)

func sum(arr []int, target int) (result []int) {
	size := len(arr)
	if size == 1 {
		return arr
	}
	m := make(map[int]int)
	for _, v := range arr {
		if i, ok := m[v]; ok {
			result = append(result, i)
			result = append(result, v)
			break
		} else {
			m[target-v] = v
		}
	}

	return
}

func main() {
	arr := []int{2, 5, 1, 7, 10, 0, 8}

	fmt.Println(sum(arr, 10))
}
