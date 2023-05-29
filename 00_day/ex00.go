package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	// "os"
	"sort"
	"strconv"
	// "strings"
	// "flag"
)

type sMethod struct {
	meanMethod   bool
	medianMethod bool
	sdMethod     bool
	modeMethod   bool
	allMethods   bool
}

const (
	minNbr = -100000
	maxNbr = 100000
)

type sPair struct {
	key int
	val int
}

type PairList []sPair

// type Interface interface {
// 	Len() int
// }

// func scan() string {
// 	in := bufio.NewReader(os.Stdin)
// 	str, err := in.ReadString('\n')
// 	fmt.Println("str in scan: ", str) //
// 	if err != nil {
// 		fmt.Println("Error of enter", err)
// 	}
// 	return str
// }

/*
func readData(str string) (int, int, map[int]int, []int) {
	countNbr := 0
	sumNbr := 0
	nbrsMap := make(map[int]int)
	var nbrsReadArr []int
	// if str[len(str)-2:len(str)-1] >= "0" && str[len(str)-2:len(str)-1] <= "9" {
	// 	num, err := strconv.Atoi(str[len(str)-2 : len(str)-1])
	// 	if err == nil && num >= -100000 && num <= 100000 {
	// 		nbrsMap[num] = 1
	// 		countNbr++
	// 		sumNbr += num
	// 		nbrsReadArr = append(nbrsReadArr, num)
	// 	} else {
	// 		fmt.Println("wrong agrument")
	// 		os.Exit(1)
	// 	}
	// }
	words := strings.Split(str, "\\n")
	for _, i := range words[:len(words)-1] {
		num, err := strconv.Atoi(i)
		fmt.Println("map: ", nbrsMap) //
		fmt.Println("err: ", err)     //

		if err == nil { //&& (num >= -100000 && num <= 100000) {
			check, _ := nbrsMap[num]
			if check == 0 { //&& ok == false {
				nbrsMap[num] = 1
			} else {
				nbrsMap[num] += 1
			}
			countNbr++
			sumNbr += num
			nbrsReadArr = append(nbrsReadArr, num)
		} else {
			fmt.Println("wrong agrument")
			os.Exit(1)
		}
	}
	lastWordArr := words[len(words)-1:]
	lastWord := lastWordArr[0]
	strings.TrimSpace(lastWord)
	fmt.Print("lastWord: ", lastWord)
	fmt.Println("map at the end: ", nbrsMap) //
	number, err := strconv.Atoi(lastWord)
	fmt.Println("err Atoi number: ", err)
	fmt.Println("number: ", number)

	return countNbr, sumNbr, nbrsMap, nbrsReadArr
}
*/

/*
func readData(str string) (int, int, map[int]int, []int) {
	countNbr := 0
	sumNbr := 0

	fmt.Print("str: ", str)
	var nbrsReadArr []int
	strArr := strings.Split(str, "\\n")
	fmt.Println("strArr: ", strArr)
	for _, s := range strArr {
		i, _ := strconv.Atoi(s)
		nbrsReadArr = append(nbrsReadArr, i)
	}
	fmt.Println("nbrsReadArr: ", nbrsReadArr)

	nbrsMap := make(map[int]int)

	return countNbr, sumNbr, nbrsMap, nbrsReadArr
}
*/

func scaningNbr() (bool, float64, bool) {
	endRead := false
	addToArr := true
	in := bufio.NewReader(os.Stdin)
	nbr, err := in.ReadString('\n')
	if err != nil && err != io.EOF {
		addToArr = false
		_, err := fmt.Fprint(os.Stderr, "Error of argument\n")
		if err != nil {
			return addToArr, 0, false
		}
	} else if io.EOF == err {
		endRead = true
	}
	val := strings.TrimSpace(nbr)
	if addToArr && val == "" {
		if err != io.EOF {
			_, err := fmt.Fprint(os.Stderr, "Input is empty, press Ctrl+d if you want to end enter\n")
			if err != nil {
				return addToArr, 0, false
			}
		}
		addToArr = false
	}
	valInt, err := strconv.Atoi(val)
	valFloat := float64(valInt)
	if addToArr && err != nil {
		addToArr = false
		_, err := fmt.Fprint(os.Stderr, "Wrong number\n")
		if err != nil {
			return addToArr, 0, false
		}
	}
	if addToArr && valFloat < minNbr {
		addToArr = false
		_, err := fmt.Fprint(os.Stderr, "Input argument smaller then %.0d\n", minNbr)
		if err != nil {
			return addToArr, 0, false
		}
	}
	if addToArr && valFloat > maxNbr {
		addToArr = false
		_, err := fmt.Fprint(os.Stderr, "Input argument bigger then %.0d\n", maxNbr)
		if err != nil {
			return addToArr, 0, false
		}
	}
	return addToArr, valFloat, endRead
}

func makeMean(sumNbr float64, countNbr int) float64 {
	mean := 0.0
	mean = float64(sumNbr) / float64(countNbr)
	// fmt.Printf("Mean: %.2f\n", mean)
	return mean
}

func makeStandartDivitation(nbrsReadArr []float64, mean float64, countNbr int) {
	var sDivitation float64
	for _, val := range nbrsReadArr {
		sDivitation += math.Pow(float64(val)-mean, 2)
	}
	sDivitation = sDivitation / float64(countNbr-1)
	sDivitation = math.Sqrt(sDivitation)
	fmt.Printf("SD: %.2f\n", sDivitation)
}

func makeMedian(nbrsReadArr []float64) {
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

func showAllCommands() {
	flag.PrintDefaults()
	os.Exit(0)
}

func handler(nbrsReadArr []float64, nbrsMap map[int]int, countNbr int, sumNbr float64, method sMethod) {
	mean := makeMean(sumNbr, countNbr)
	if method.allMethods || method.meanMethod || method.sdMethod {
		if method.meanMethod || method.allMethods {
			fmt.Printf("Mean: %.2f\n", mean)
		}
	}
	if method.allMethods || method.medianMethod {
		makeMedian(nbrsReadArr)
	}
	if method.allMethods || method.modeMethod {
		makeMode(nbrsMap)
	}
	if method.allMethods || method.sdMethod {
		makeStandartDivitation(nbrsReadArr, mean, countNbr)
	}
}

func main() {
	// str := scan()

	method := sMethod{false, false, false, false, true}
	var printUsage bool
	flag.BoolVar(&method.meanMethod, "mean", false, "show mean of arguments")
	flag.BoolVar(&method.medianMethod, "median", false, "show median of arguments")
	flag.BoolVar(&method.modeMethod, "mode", false, "show mode of  arguments")
	flag.BoolVar(&method.sdMethod, "standart devitation", false, "show standart devitation of arguments")
	flag.BoolVar(&printUsage, "help", false, "show all commands")
	flag.Parse()
	if printUsage {
		showAllCommands()
	}
	method.allMethods = !(method.meanMethod || method.medianMethod || method.modeMethod || method.sdMethod)

	var nbrsReadArr []float64
	nbrsMap := make(map[int]int)
	countNbr := 0
	sumNbr := 0.0
	fmt.Printf("Enter a number > %.0d and < %.0d:\n", minNbr, maxNbr)
	for {
		addToArr, valFloat, endRead := scaningNbr()
		if addToArr {
			nbrsReadArr = append(nbrsReadArr, valFloat)
			sumNbr += valFloat
			countNbr++
			check, _ := nbrsMap[int(valFloat)]
			if check == 0 { //&& ok == false {
				nbrsMap[int(valFloat)] = 1
			} else {
				nbrsMap[int(valFloat)] += 1
			}
		}
		if endRead {
			fmt.Printf("\n")
			break
		}
	}
	if len(nbrsReadArr) != 0 {
		handler(nbrsReadArr, nbrsMap, countNbr, sumNbr, method)
	} else {
		fmt.Println("Values no entered")
	}

	// // countNbr, sumNbr, nbrsMap, nbrsReadArr := readData(str)
	// //add check -100000 и 100000
	// fmt.Println("\nPlease, choose one of the available commands:")
	// fmt.Println("1.\tShow mean\n2.\tShow median\n3.\tShow mode")
	// fmt.Println("4.\tShow standart divitation\n5.\tShow all metrics\n6.\tExit program")
	// fmt.Println("Please, choose one and enter command number:")
	// fmt.Println()

	// var method int
	// for {
	// 	fmt.Scan(&method)
	// 	if method >= 1 && method <= 6 {
	// 		break
	// 	} else {
	// 		fmt.Println("Enter digit from 1 to 6")
	// 	}
	// }

	// mean := makeMean(float64(sumNbr), float64(countNbr))

	// switch method {
	// case 1:
	// 	fmt.Printf("Mean: %.2f\n", mean)
	// case 2:
	// 	makeMedian(nbrsReadArr)
	// case 3:
	// 	makeMode(nbrsMap)
	// case 4:
	// 	makeStandartDivitation(nbrsReadArr, mean, countNbr)
	// case 5:
	// 	fmt.Printf("Mean: %.2f\n", mean)
	// 	makeMedian(nbrsReadArr)
	// 	makeMode(nbrsMap)
	// 	makeStandartDivitation(nbrsReadArr, mean, countNbr)
	// case 6:
	// 	return
	// default:
	// 	fmt.Println("Unknown symbol")
	// }

	// // makeMode(nbrsMap)
	// // makeMedian(nbrsReadArr)
	// // makeStandartDivitation(nbrsReadArr, mean, countNbr)
}
