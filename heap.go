package main

import (
	"fmt"
	"math/rand"
	"time"
)

///////////
// heap 特征
//
//  parent(t) = (t- 1) >> 1， 即t的父节点下标=(t-1)/2,  例如： t = 6，parent(t) = 2
//  left(t) = t << 1 + 1, 即t的左孩子的节点下标=t * 2 + 1, 例如： t= 2,left(t) = 5
//  reght(t) = t << 1 + 2, 即t的右孩子的节点下标=t * 2 + 2, 例如：t = 2,right(t) = 6
//
// 堆的主要操作有：
//      构建堆
//      移除堆的根节点
//      往堆中插入一个节点
//      调整堆
///////////

func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

type Heap struct {
	data   []int
	length int
}

func (heap *Heap) swap(n1, n2 int) {
	heap.data[n1], heap.data[n2] = heap.data[n2], heap.data[n1]
}

func (heap *Heap) getParentIndex(node int) int {
	return (node - 1) >> 1
}

func (heap *Heap) getLeftChildIndex(node int) int {
	return (node << 1) + 1
}

func (heap *Heap) getRightChildIndex(node int) int {
	return (node << 1) + 2
}

func (heap *Heap) print() string {
	return fmt.Sprintf("data:%v", heap.data)
}

func (heap *Heap) sortPrint(node int) {
	//fmt.Println(node, "node", heap.length)
	if node < heap.length {
		fmt.Println(heap.data[node])
		heap.sortPrint(heap.getLeftChildIndex(node))
		heap.sortPrint(heap.getRightChildIndex(node))
	}
}

// type MaxHeap struct {
// 	Heap
// }

func (heap *Heap) buildHeapDown() {
	start := heap.getParentIndex(heap.length - 1)
	for start >= 0 {
		heap.adjustDownHeap(start, len(heap.data))
		start -= 1
	}
}

func (heap *Heap) buildHeapUp() {
	start := 0
	for start < heap.length {
		heap.adjustUpHeap(start)
		start += 1
	}
}

func (heap *Heap) adjustDownHeap(node int, size int) {
	for {
		right := heap.getRightChildIndex(node)
		left := heap.getLeftChildIndex(node)
		max := node
		if right < size && heap.data[right] > heap.data[max] {
			max = right
		}
		if left < size && heap.data[left] > heap.data[max] {
			max = left
		}
		if max != node {
			heap.swap(node, max)
			node = max
		} else {
			break
		}
	}
}

func (heap *Heap) adjustUpHeap(node int) {
	parent := heap.getParentIndex(node)
	if parent >= 0 && heap.data[parent] < heap.data[node] {
		heap.swap(node, parent)
		heap.adjustUpHeap(parent)
	}
}

func main() {
	data := []int{14, 84, 74, 7, 35, 88, 88, 81, 94, 61, 12, 86, 41, 59, 24, 69, 20, 74, 89, 58, 61}
	// for {
	// 	i := RandInt(0, 100)
	// 	data = append(data, i)
	// 	if len(data) > 20 {
	// 		break
	// 	}
	//}
	fmt.Println(data)
	heap := Heap{data: data, length: len(data)}
	//heap.sortPrint(0)
	heap.buildHeapDown()
	fmt.Println(heap.data)
	for i := len(heap.data) - 1; i > 0; i-- {
		heap.swap(0, i)
		heap.adjustDownHeap(0, i)
	}
	fmt.Println(heap.data)
}
