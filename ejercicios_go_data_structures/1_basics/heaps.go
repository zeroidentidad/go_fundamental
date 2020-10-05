package main

// importing package container/heap
import (
	"container/heap"
	"fmt"
)

// integerHeap type para struct methods
type IntegerHeap []int

// IntegerHeap method - obtiene la longitud de integerHeap
func (iheap IntegerHeap) Len() int { return len(iheap) }

// IntegerHeap method - comprueba si el elemento del índice i
// es menor que el índice j
func (iheap IntegerHeap) Less(i, j int) bool { return iheap[i] < iheap[j] }

// IntegerHeap method - intercambia el elemento del índice i al j
func (iheap IntegerHeap) Swap(i, j int) { iheap[i], iheap[j] = iheap[j], iheap[i] }

//IntegerHeap method - agrega el elemento al montón
func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

//IntegerHeap method - saca el elemento del montón
func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}

	heap.Init(intHeap)
	heap.Push(intHeap, 8)
	heap.Push(intHeap, 3)
	fmt.Printf("minimo: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
