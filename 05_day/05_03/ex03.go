package main

import (
	"errors"
	"fmt"
	"log"
)

func maxNbr(f, s int) int {
	if f > s {
		return f
	}
	return s
}

func restoreAnswer(dynamic [][]int, items []Present, countItems int, capacity int, rez *[]Present) error {
	if dynamic[countItems][capacity] == 0 {
		err := errors.New("No one item doesn't enter in bag")
		return err
	}
	if dynamic[countItems-1][capacity] == dynamic[countItems][capacity] {
		restoreAnswer(dynamic, items, countItems-1, capacity, rez)
	} else {
		restoreAnswer(dynamic, items, countItems-1, capacity-items[countItems-1].Size, rez)
		*rez = append(*rez, items[countItems-1])
	}
	return nil
}

func grabPresents(items []Present, capacity int) []Present {
	coolPresents := []Present{}
	dynamic := make([][]int, len(items)+1)
	for i := 0; i < len(items)+1; i++ {
		dynamic[i] = make([]int, capacity+1)
		for j := 0; j < capacity+1; j++ {
			if i == 0 || j == 0 {
				dynamic[i][j] = 0
			} else {
				if j >= items[i-1].Size {
					dynamic[i][j] = maxNbr(dynamic[i-1][j], dynamic[i-1][j-items[i-1].Size]+items[i-1].Value)
				} else {
					dynamic[i][j] = dynamic[i-1][j]
				}
			}
		}
	}
	err := restoreAnswer(dynamic, items, len(items), capacity, &coolPresents)
	if err != nil {
		log.Fatal("No one item doesn't enter in bag")
	}
	return coolPresents
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

func main() {
	var capacity int
	fmt.Println("Enter capacity of knapsack:")
	fmt.Scan(&capacity)
	// items := []Present{{30, 6}, {14, 3}, {16, 4}, {9, 2}}
	items := readData()
	bag := grabPresents(items, capacity)
	fmt.Println(bag)
}
