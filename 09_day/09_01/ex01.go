package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func getBody(url string, chanWrite chan<- *string) {
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	str := string(body)
	log.Printf(str)
	chanWrite <- &str
}

func crawlWeb(chanIn chan string, wg *sync.WaitGroup) chan *string {
	chanWrite := make(chan *string, 50)

	for value := range chanIn {
		wg.Add(1)
		go func(str string) {
			defer wg.Done()
			getBody(str, chanWrite)
		}(value)
	}
	return chanWrite
}

func sendUrls(ctx context.Context, urls *[]string, ch chan<- string) {
	for _, url := range *urls {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			time.Sleep(time.Duration(1*1000) * time.Millisecond)
			ch <- url
		}
	}
	close(ch)
}

func startInitial() ([]string, []*string, chan string, *sync.WaitGroup) {
	urlArray := []string{"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com",
		"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com",
		"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com",
		"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com",
		"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com",
		"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com"}
	var bodyPtrArray []*string
	chanWrite := make(chan string, 100)
	return urlArray, bodyPtrArray, chanWrite, &sync.WaitGroup{}
}

func handlSignals(cancel context.CancelFunc) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT)
	for {
		sig := <-sigs
		switch sig {
		case os.Interrupt:
			fmt.Println("\nexiting")
			cancel()
			return
		}
	}
}

// Main goroutine
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go handlSignals(cancel)
	defer cancel()
	runtime.GOMAXPROCS(8)
	urlArray, bodyPtrArray, chanWrite, wg := startInitial()

	go sendUrls(ctx, &urlArray, chanWrite)
	chanRead := crawlWeb(chanWrite, wg)

	go func() {
		wg.Wait()
		close(chanRead)
	}()

	for t := range chanRead {
		bodyPtrArray = append(bodyPtrArray, t)
	}

	for _, Ptr := range bodyPtrArray {
		fmt.Println("\n", Ptr)
	}
}
