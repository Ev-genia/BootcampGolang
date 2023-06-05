package main

import (
	"container/heap"
	"fmt"
	"log"
)

type Item struct {
	value    Present
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return (pq[i].value.Value > pq[j].value.Value) || (pq[i].value.Value == pq[j].value.Value && pq[i].value.Size < pq[j].value.Size)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, present Present, priority int) {
	item.value.Value = present.Value
	item.value.Size = present.Size
	item.priority = priority
	heap.Fix(pq, item.index)
}

func readData() (data []Present) {
	fmt.Println("Enter two number: value and size of present. Enter 0 0 for exit")
	var elem Present
	for {
		_, err := fmt.Scanf("%d %d", &elem.Value, &elem.Size)
		if err != nil {
			log.Fatal("Error Scanf: ", err)
		}
		if elem.Value == 0 && elem.Size == 0 {
			break
		}
		data = append(data, elem)
	}
	return data
}

func checkMinSize(elem1 Present, elem2 Present) (Present, Present) {
	if elem1.Size < elem2.Size {
		return elem1, elem2
	} else {
		return elem2, elem1
	}
}

func getNCoolestPresents(items []Present, n int) (coolestPresents []Present, err error) {
	if n < 0 || n > len(items) {
		log.Fatal("n is negative or is larger than the size of the slice\n")
	}
	pq := make(PriorityQueue, len(items))
	i := 0
	for priority, value := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	i = 0
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if i < n {
			coolestPresents = append(coolestPresents, item.value)
		}
		i++
	}
	return coolestPresents, nil
}

func main() {
	var n int
	fmt.Println("Enter count of first prestnts - n:")
	fmt.Scan(&n)
	// items := []Present{{5, 2}, {4, 5}, {5, 3}, {5, 1}}
	items := readData()
	rez, err := getNCoolestPresents(items, n)
	if err != nil {
		log.Fatal("Error of getting coolest presents: %s\n", err)
	}
	fmt.Println(rez)
}
