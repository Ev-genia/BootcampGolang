package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func sleepSort(ctx context.Context, inputs []int) chan int {
	output := make(chan int)
	for _, in := range inputs {
		select {
		case <-ctx.Done():
			fmt.Println("too long")
			return output
		default:
			x := in
			go func(int, <-chan int) {
				time.Sleep(time.Duration(x) * time.Second)
				output <- x
			}(x, output)
		}

	}
	return output
}

func handlSignals(cancel context.CancelFunc) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT)
	for {
		sig := <-sigs
		switch sig {
		case os.Interrupt:
			// fmt.Println("\nexiting")
			cancel()
			os.Exit(0)
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go handlSignals(cancel)

	var input = []int{4, 7, 2, 9, 5, 3, 1, 6, 4, 3, 2, 4, 6, 7, 9, 9, 33}
	wait := sleepSort(ctx, input)
	defer close(wait)
	for i := 0; i < len(input); i++ {
		fmt.Println(<-wait)
	}
}
