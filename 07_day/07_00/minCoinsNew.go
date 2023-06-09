package main

import (
	"fmt"
	"sort"
)

// проверка вхождения числа в массив
func FindInt(value int, coins *[] int)(int, int, bool) {
	for i, val1 := range *coins {
		if  value == val1 {
			return i, value, true
		}
	}
	return -1, -1, false
}

// Главная функция
func minCoinsNew(val int, coins []int) ([]int) {
	rez := make([]int, 0, 200)
	if val < 0 {
		fmt.Print("Error: Negative value.\n")
		return rez
	} else if val == 0 {
		return rez
	}

	var coinsValidated [] int

	for _, val1 := range coins {
		_, _, isExist := FindInt(val1, &coinsValidated)
		if !(isExist) {
			coinsValidated = append(coinsValidated, val1)
		}
	}
	sort.Ints(coinsValidated)
	if coinsValidated[0] < 0 {
		fmt.Print("Error: Negative denomination.\n")
		return rez
	}
	var freqCoins [100] int

	i := len(coinsValidated) - 1
	for i >= 0 {
		freqCoins[i] = val / coinsValidated[i]
			if freqCoins[i] > 0 && i > 0 && val - coinsValidated[i] * (freqCoins[i] - 1) == coinsValidated[i-1] * 2 {
				freqCoins[i] -= 1
				freqCoins[i - 1] = 2
				break
			}
		val = val % coinsValidated[i]
		i--
	}
	i = len(coinsValidated) - 1
	for i >= 0 {
		value := coinsValidated[i]
		for j := 0; j < freqCoins[i]; j++ {
			rez = append(rez, value)
		}
		i--
	}
	return rez

}

// func main(){
// 	fmt.Println(minCoinsNew(285,[]int{1,2,5,10}))
// }