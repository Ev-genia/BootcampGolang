package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func sleepSort(inputs []int, sig <-chan os.Signal) chan int {

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
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	var input = []int{4, 7, 2, 9, 5, 3, 1}
	wait := sleepSort(input, sigs)
	defer close(wait)
	for i := 0; i < len(input); i++ {
		fmt.Println(<-wait)
	}
	sig := <-sigs
	fmt.Println("got of signal: ", sig)
}
