package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
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
	// канал для получения этих уведомлений
	// (мы также создадим канал, чтобы уведомить нас,
	// когда программа может выйти).
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// регистрирует данный канал для
	// получения уведомлений об указанных сигналах
	signal.Notify(sigs, syscall.SIGINT)

	// горутина выполняет блокировку приема сигналов
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	var input = []int{4, 7, 2, 9, 5, 3, 1}
	wait := sleepSort(input)
	defer close(wait)
	for i := 0; i < len(input); i++ {
		fmt.Println(<-wait)
	}

	// Программа будет ждать здесь, пока не получит ожидаемый сигнал
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
