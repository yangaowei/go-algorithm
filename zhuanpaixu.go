package main

import (
	"fmt"
)

func swap(arr []string, n1, n2 int) {
	arr[n1], arr[n2] = arr[n2], arr[n1]
}

func AllRange(arr []string, start, length int) {

	if start == length-1 {
		fmt.Println(arr)
	} else {
		for index, _ := range arr {
			//fmt.Println(index, v)
			// if index+1 == length {
			// 	break
			// }
			swap(arr, start, index)
			AllRange(arr, start+1, length)
			//swap(arr, start, index)
		}
	}
}

func main() {
	arr := []string{"a", "b", "c"}
	AllRange(arr, 0, 3)
}
