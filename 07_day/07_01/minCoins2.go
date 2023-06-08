package main

import (
	"log"
	"sort"
)

func minCoins2(val int, coins []int) (res []int) {
	if val < 1 {
		log.Print("Negative value\n")
		return
	}

	coinsMap := make(map[int]bool)
	for _, v := range coins {
		coinsMap[v] = true
	}

	money := []int{}
	for v := range coinsMap {
		money = append(money, v)
	}

	sort.Slice(money, func(i, j int) bool {
		return money[i] < money[j]
	})

	i := len(money) - 1
	for i >= 0 {
		if coins[i] < 1 {
			log.Print("Negative coin value\n")
			return []int{}
		}
		for val >= money[i] {
			val -= money[i]
			res = append(res, money[i])
		}
		i -= 1
	}

	return
}
