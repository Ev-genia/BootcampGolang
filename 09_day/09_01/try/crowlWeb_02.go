package main

import (
	//	"fmt"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	//	"net/url"
	"io/ioutil"
	"log"
	"runtime"
)

func getBody(url string, chanWrite chan<- *string) {
	log.Println("enter to getBody, url: ", url)
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
	log.Println("get body | len str: ", len(str))
	chanWrite <- &str
	log.Println("get body | len of chanal: ", len(chanWrite))
}

func crawlWeb(ctx context.Context, chanIn chan string) chan *string {
	chanWrite := make(chan *string, 50)
	var wg sync.WaitGroup

	for value := range chanIn {
		log.Println("crawlWeb | value: ", value)
		wg.Add(1)
		go getBody(value, chanWrite)
		log.Println("\ncrawlWeb |after getBody, len chanW: ", len(chanWrite))
		wg.Done()
	}
	log.Println("\ncrawlWeb | before wg.Wait, len chanW: ", len(chanWrite))
	wg.Wait()
	log.Println("crawlWeb | after wg.Wait, len chanW: ", len(chanWrite))
	return chanWrite
}

func sendUrls(ctx context.Context, urls *[]string, ch chan<- string) {
	for i, url := range *urls {
		time.Sleep(time.Duration(100) * time.Millisecond)
		// log.Println("ctx err: ", ctx.Err(), ", i: ", i)
		if ctx.Err() != nil {
			log.Println("ctx in sendUrls: ", ctx.Err(), ", i: ", i)
			// ch <- url
			return
		}
		log.Println("ch <- url in sendUrls: ", url, ", i: ", i)
		ch <- url
	}
	log.Println("closing ch in sendUrls")
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
		"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com",
		"https://golangify.com", "https://www.itcodet.com", "https://pkg.go.dev/std",
		"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com",
		"https://golangify.com", "https://www.itcodet.com", "https://pkg.go.dev/std",
		"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com",
		"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com",
		"https://golangify.com", "https://www.itcodet.com", "https://pkg.go.dev/std",
		"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com",
		"https://golangify.com", "https://www.itcodet.com", "https://pkg.go.dev/std",
		"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com"}
	log.Println("len: ", len(urlArray))
	var bodyPtrArray []*string
	chanWrite := make(chan string)

	go sendUrls(ctx, &urlArray, chanWrite)

	chanRead := crawlWeb(ctx, chanWrite)
	log.Println("len chanRead in main: ", len(chanRead))

	for ptr := range chanRead {
		bodyPtrArray = append(bodyPtrArray, ptr)
	}
	defer close(chanRead)
	count := 0
	for _, Ptr := range bodyPtrArray {
		fmt.Println(Ptr, " : ", count)
		count++
	}
}
