package main

import (
	//	"fmt"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	//	"net/url"
	"io/ioutil"
	"log"
	"runtime"
)

func getBody(ctx context.Context, url string, chanWrite chan *string) {
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
	//	log.Printf(str)
	select {
	case val, ok := <-chanWrite:
		if !ok {
			// канал закрыт
			fmt.Println("stop in: ", val)
			break
		}
	default:
		chanWrite <- &str
	}

}

func crawlWeb(ctx context.Context, chanIn chan string) chan *string {
	chanWrite := make(chan *string, 1)
	for value := range chanIn {
		select {
		case <-ctx.Done():
			// fmt.Println("too long")
			fmt.Println("value: ", value)
			close(chanWrite)
			return nil

		case <-time.After(200 * time.Millisecond):
			fmt.Printf("value defauilt: ", value)
			go getBody(ctx, value, chanWrite)
		}
	}
	return chanWrite
}

func sendUrls(ctx context.Context, urls *[]string, ch chan<- string) {
	for i, url := range *urls {
		fmt.Println("url: ", url)
		select {
		case <-ctx.Done():
			// fmt.Println("too long")
			// close(ch)
			return
		case <-time.After(100 * time.Millisecond):
			fmt.Println("i: ", i)
			// default:
			ch <- url
		}
	}
	close(ch)
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
			// defer os.Exit(0)
			return
		}
	}
}

// Main goroutine
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go handlSignals(cancel)

	runtime.GOMAXPROCS(8)
	urlArray := []string{"https://www.itcodet.com", "https://pkg.go.dev/std",
		"https://golangify.com", "https://www.itcodet.com", "https://pkg.go.dev/std",
		"https://golangify.com", "https://www.itcodet.com", "https://pkg.go.dev/std",
		"https://golangify.com", "https://www.itcodet.com", "https://pkg.go.dev/std",
		"https://golangify.com", "https://www.itcodet.com", "https://pkg.go.dev/std",
		"https://golangify.com", "https://www.itcodet.com", "https://pkg.go.dev/std"}
	var bodyPtrArray []*string
	chanWrite := make(chan string, 2)
	chanRead := make(chan *string, 10)

	go sendUrls(ctx, &urlArray, chanWrite)

	select {
	case _, ok := <-chanWrite:
		if ok {
			chanRead = crawlWeb(ctx, chanWrite)
		}
	default:
		fmt.Println("chanal is close ")
	}
	for range urlArray {
		bodyPtrArray = append(bodyPtrArray, <-chanRead)
	}
	close(chanRead)

	for _, Ptr := range bodyPtrArray {
		fmt.Println("\n", Ptr)
	}
}
