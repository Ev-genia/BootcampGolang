package main

import "fmt"

func sendData(elem chan interface{}, output chan<- interface{}) {
	for val := range elem {
		output <- val
	}
}

func multiplex(elem ...chan interface{}) chan interface{} {
	output := make(chan interface{}, 20)

	for _, oneCh := range elem {
		go sendData(oneCh, output)
	}
	return output
}

func main() {
	ints := []int{2, 7, 9, 5, 3, 2}
	intsI := make([]interface{}, len(ints))
	for i, val := range ints {
		intsI[i] = val
	}
	var intCh = make(chan interface{}, 20)
	for _, val := range intsI {
		intCh <- val
	}

	ints2 := []int{11, 77, 99, 55, 33, 11}
	intsI2 := make([]interface{}, len(ints2))
	for i, val := range ints2 {
		intsI2[i] = val
	}
	var intCh2 = make(chan interface{}, 20)
	for _, val := range intsI2 {
		intCh2 <- val
	}

	// rezInt := multiplex(intCh, intCh2)
	// for i := 0; i < len(ints)+len(ints2); i++ {
	// 	fmt.Println(<-rezInt)
	// }
	// fmt.Println()
	strs := []string{"f", "ksf", "sjd"}
	strI := make([]interface{}, len(strs))
	for i, val := range strs {
		strI[i] = val
	}
	var strCh = make(chan interface{}, 20)
	for _, val := range strI {
		strCh <- val
	}

	strs2 := []string{"FGDD", "THH", "^__^"}
	strI2 := make([]interface{}, len(strs2))
	for i, val := range strs2 {
		strI2[i] = val
	}
	var strCh2 = make(chan interface{}, 20)
	for _, val := range strI2 {
		strCh <- val
	}

	rezStr := multiplex(strCh, strCh2, intCh, intCh2)
	for i := 0; i < len(strs)+len(strs2)+len(ints)+len(ints2); i++ {
		fmt.Println(<-rezStr)
	}
}
