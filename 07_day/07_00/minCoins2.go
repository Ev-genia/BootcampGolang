package main

type Change struct {
	val   int
	coins []int
}

func minCoins2(money int, coins []int) []int {
	if money <= 0 {
		return []int{}
	}
	arrVal := make([]Change, money+1)
	for i := 0; i < money+1; i++ {
		arrVal[i].val = i
	}
	for i := 0; i < money+1; i++ {
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
	return arrVal[money].coins
}
