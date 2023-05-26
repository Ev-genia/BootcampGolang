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

func main() {
	str := scan()
	countNbr := 0
	sumNbr := 0
	// fmt.Println("str[len(str)]: ", str[len(str)-2:len(str)-1])
	// fmt.Println("TypeOf(str[len(str)]): ", reflect.TypeOf(str[len(str)-2:len(str)-1]))
	nbrsMap := make(map[int]int)
	var nbrsReadArr []int
	// j := 0
	if str[len(str)-2:len(str)-1] >= "0" && str[len(str)-2:len(str)-1] <= "9" {
		num, err := strconv.Atoi(str[len(str)-2 : len(str)-1])
		if err == nil {
			// fmt.Println("num: ", num)
			// fmt.Println("TypeOf(num): ", reflect.TypeOf(num))
			nbrsMap[num] = 1
			countNbr++
			sumNbr += num
			nbrsReadArr = append(nbrsReadArr, num)
		}
	}
	// str += "\\n"              //
	// fmt.Print("lines: ", str) //
	words := strings.Split(str, "\\n")
	// fmt.Println("len: ", len(words)) //
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

	// for key, val := range nbrsMap { //
	// 	fmt.Println("key: ", key, "val: ", val) //
	// } //
	// fmt.Println(nbrsMap)
	// fmt.Println("countNbr: ", countNbr) //
	// fmt.Println("sumNbr: ", sumNbr)     //

	mean := 0.0
	mean = float64(sumNbr) / float64(countNbr)
	fmt.Printf("Mean: %.2f\n", mean)

	// fmt.Println("len map: ", len(nbrsMap)) //

	var nbrsArr []sPair
	// nbrsArr := make([]sPair, len(nbrsMap))
	// j := 0
	for key, val := range nbrsMap {
		// 	nbrsArr[j].key = key
		// 	nbrsArr[j].val = val
		// 	j++
		nbrsArr = append(nbrsArr, sPair{key, val})
	}

	// // fmt.Println("Arr:")            //
	// // for _, elem := range nbrsArr { //
	// // 	fmt.Println(elem, " ") //
	// // } //

	sort.Slice(nbrsArr, func(i, j int) bool {
		return nbrsArr[i].val > nbrsArr[j].val
	})

	// // fmt.Println("Arr sort:")       //
	// // for _, elem := range nbrsArr { //
	// // 	fmt.Println(elem, " ") //
	// // } //

	mode := nbrsArr[0]
	fmt.Println("Mode: ", mode.key)

	// // fmt.Println("len arr: ", len(nbrsArr)) //
	// // fmt.Println("len/2 arr: ", len(nbrsArr)/2) //
	var median float64
	// if len(nbrsArr)%2 == 0 {
	// 	median = (float64(nbrsArr[len(nbrsArr)/2-1].key) + float64(nbrsArr[len(nbrsArr)/2].key)) / 2
	// } else {
	// 	median = float64(nbrsArr[len(nbrsArr)/2].key)
	// }
	// fmt.Println("Median map: ", median)
	// fmt.Println("read arr: ", nbrsReadArr)

	sort.Slice(nbrsReadArr, func(i, j int) bool {
		return nbrsReadArr[i] < nbrsReadArr[j]
	})
	// fmt.Println("read arr sort: ", nbrsReadArr)
	if len(nbrsReadArr)%2 == 0 {
		median = (float64(nbrsReadArr[len(nbrsReadArr)/2-1]) + float64(nbrsReadArr[len(nbrsReadArr)/2])) / 2
	} else {
		median = float64(nbrsReadArr[len(nbrsReadArr)/2])
	}
	fmt.Printf("Median: %.2f\n", median)
	var sDivitation float64
	for _, val := range nbrsReadArr {
		sDivitation += math.Pow(float64(val)-mean, 2)
	}
	sDivitation = sDivitation / float64(countNbr-1)
	sDivitation = math.Sqrt(sDivitation)
	fmt.Printf("SD: %.2f\n", sDivitation)
}
