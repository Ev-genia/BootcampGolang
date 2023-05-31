package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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
	// var mode int
	// max := 0
	// for _, i := range nbrsArr {
	// 	freq := nbrsMap[i.key]
	// 	if freq > max {
	// 		mode = i.key
	// 		max = freq
	// 	}
	// }
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
			if check == 0 {
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
}
