package main

import (
	"fmt"
	"time"
)

func sleepSort(inputs []int) chan int {
	output := make(chan int)
	for _, in := range inputs {
		x := in
		go func(int, <-chan int) {
			time.Sleep(time.Duration(x) * time.Second)
			output <- x
		}(x, output)
	}
	return output
}

func main() {
	var input = []int{4, 7, 2, 9, 5, 3, 1}
	wait := sleepSort(input)
	defer close(wait)
	for i := 0; i < len(input); i++ {
		fmt.Println(<-wait)
	}
}
