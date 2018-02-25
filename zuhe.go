package main

import (
	"fmt"
	//"strings"
)

func zh(arr []int) interface{} {
	size := len(arr)
	if size == 1 {
		return arr
	}
	if size == 2 {
		return []interface{}{[]int{arr[0]}, []int{arr[1]}, []int{arr[0], arr[1]}}
	}
	if size > 3 {
		a := zh(arr[:size-1])
        result := []interface{}
        for _,value := range result{
            
        }
	}  
	return []int{}
}

func main() {
	arr := []int{1, 2}
	fmt.Println(zh(arr))
}
