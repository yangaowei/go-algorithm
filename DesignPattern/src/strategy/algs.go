package strategy

import (
	"fmt"
)

type IAlg interface {
	used()
}

type BubbleSort struct {
	data []int
}

func (self *BubbleSort) used() {
	fmt.Printf("origin bubblesort data: %v\n", self.data)
}

func NewBubbleSort(arr []int) *BubbleSort {
	bSort := &BubbleSort{arr}

	return bSort
}

type QuickSort struct {
	data []int
}

func NewQuickSort(arr []int) *QuickSort {
	qSort := &QuickSort{arr}
	return qSort
}

func partition(a []int, low int, hight int) int {
	//fmt.Println(a, low, hight)
	privotKey := a[low]
	for j := low; j <= hight; j++ {
		//fmt.Println(a, low, hight)
		for low < hight && a[hight] >= privotKey {
			hight -= 1
		}
		a[low], a[hight] = a[hight], a[low]
		for low < hight && a[low] <= privotKey {
			low += 1
			a[low], a[hight] = a[hight], a[low]
		}
		a[low], a[hight] = a[hight], a[low]

	}
	return low
}

func quickSort(a []int, low int, hight int) {
	if low < hight {
		privotLoc := partition(a, low, hight)
		quickSort(a, low, privotLoc-1) //递归对低子表递归排序
		quickSort(a, privotLoc+1, hight)
	}
}

func (self *QuickSort) used() {
	fmt.Printf("origin quicksort data: %v\n", self.data)
	quickSort(self.data, 0, len(self.data)-1)
	fmt.Printf("sort quicksort data: %v\n", self.data)
}
