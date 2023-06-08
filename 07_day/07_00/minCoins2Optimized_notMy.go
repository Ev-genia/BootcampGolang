package main

import "fmt"

type Change struct {
	val   int
	coins []int
}

func minCoins2Optimized(val int, coins []int) []int {
	if val <= 0 {
		return []int{}
	}
	arrVal := make([]Change, val+1)
	for i := 0; i < val+1; i++ {
		arrVal[i].val = i
	}
	for i := 0; i < val+1; i++ {
		for _, coin := range coins {
			if coin == arrVal[i].val {
				arrVal[i].coins = []int{coin}
				break
			} else if coin < arrVal[i].val {
				if len(arrVal[i].coins) == 0 || len(arrVal[i].coins) > len(arrVal[arrVal[i].val-coin].coins)+1 {
					minArr := make([]int, len(arrVal[arrVal[i].val-coin].coins))
					copy(minArr, arrVal[arrVal[i].val-coin].coins)
					arrVal[i].coins = append(minArr, coin)
				}
			} else {
				break
			}
		}
	}
	return arrVal[val].coins
}

func main() {
	// var capacity int
	// fmt.Println("Enter capacity of knapsack:")
	// fmt.Scan(&capacity)
	// items := []Present{{30, 6}, {14, 3}, {16, 4}, {9, 2}}
	// items := readData()
	capacity := 1
	coins := []int{1, 5, 10}
	bag := minCoins2Optimized(capacity, coins)
	fmt.Println(bag)
}
