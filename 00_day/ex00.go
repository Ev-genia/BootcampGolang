package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type sPair struct {
	key int
	val int
}

type PairList []sPair

type Interface interface {
	Len() int
}

func scan() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Error of enter", err)
	}
	return str
}

func makeMean(sumNbr float64, countNbr float64) float64 {
	mean := 0.0
	mean = float64(sumNbr) / float64(countNbr)
	fmt.Printf("Mean: %.2f\n", mean)
	return mean
}

func makeStandartDivitation(nbrsReadArr []int, mean float64, countNbr int) {
	var sDivitation float64
	for _, val := range nbrsReadArr {
		sDivitation += math.Pow(float64(val)-mean, 2)
	}
	sDivitation = sDivitation / float64(countNbr-1)
	sDivitation = math.Sqrt(sDivitation)
	fmt.Printf("SD: %.2f\n", sDivitation)
}

func makeMedian(nbrsReadArr []int) {
	var median float64
	sort.Slice(nbrsReadArr, func(i, j int) bool {
		return nbrsReadArr[i] < nbrsReadArr[j]
	})
	if len(nbrsReadArr)%2 == 0 {
		median = (float64(nbrsReadArr[len(nbrsReadArr)/2-1]) + float64(nbrsReadArr[len(nbrsReadArr)/2])) / 2
	} else {
		median = float64(nbrsReadArr[len(nbrsReadArr)/2])
	}
	fmt.Printf("Median: %.2f\n", median)
}

func makeMode(nbrsMap map[int]int) {
	var nbrsArr []sPair
	for key, val := range nbrsMap {
		nbrsArr = append(nbrsArr, sPair{key, val})
	}
	sort.Slice(nbrsArr, func(i, j int) bool {
		return nbrsArr[i].val > nbrsArr[j].val
	})
	mode := nbrsArr[0]
	fmt.Println("Mode: ", mode.key)
}

func main() {
	str := scan()
	countNbr := 0
	sumNbr := 0
	nbrsMap := make(map[int]int)
	var nbrsReadArr []int
	if str[len(str)-2:len(str)-1] >= "0" && str[len(str)-2:len(str)-1] <= "9" {
		num, err := strconv.Atoi(str[len(str)-2 : len(str)-1])
		if err == nil {
			nbrsMap[num] = 1
			countNbr++
			sumNbr += num
			nbrsReadArr = append(nbrsReadArr, num)
		}
	}
	words := strings.Split(str, "\\n")
	for _, i := range words {
		num, err := strconv.Atoi(i)
		if err == nil {
			check, ok := nbrsMap[num]
			if check == 0 && ok == false {
				nbrsMap[num] = 1
			} else {
				nbrsMap[num] += 1
			}
			countNbr++
			sumNbr += num
			nbrsReadArr = append(nbrsReadArr, num)
		}
	}

	// mean := 0.0
	// mean = float64(sumNbr) / float64(countNbr)
	// fmt.Printf("Mean: %.2f\n", mean)
	mean := makeMean(float64(sumNbr), float64(countNbr))

	// var nbrsArr []sPair
	// for key, val := range nbrsMap {
	// 	nbrsArr = append(nbrsArr, sPair{key, val})
	// }
	// sort.Slice(nbrsArr, func(i, j int) bool {
	// 	return nbrsArr[i].val > nbrsArr[j].val
	// })
	// mode := nbrsArr[0]
	// fmt.Println("Mode: ", mode.key)
	makeMode(nbrsMap)

	// var median float64
	// sort.Slice(nbrsReadArr, func(i, j int) bool {
	// 	return nbrsReadArr[i] < nbrsReadArr[j]
	// })
	// if len(nbrsReadArr)%2 == 0 {
	// 	median = (float64(nbrsReadArr[len(nbrsReadArr)/2-1]) + float64(nbrsReadArr[len(nbrsReadArr)/2])) / 2
	// } else {
	// 	median = float64(nbrsReadArr[len(nbrsReadArr)/2])
	// }
	// fmt.Printf("Median: %.2f\n", median)
	makeMedian(nbrsReadArr)

	// var sDivitation float64
	// for _, val := range nbrsReadArr {
	// 	sDivitation += math.Pow(float64(val)-mean, 2)
	// }
	// sDivitation = sDivitation / float64(countNbr-1)
	// sDivitation = math.Sqrt(sDivitation)
	// fmt.Printf("SD: %.2f\n", sDivitation)
	makeStandartDivitation(nbrsReadArr, mean, countNbr)
}
