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

func getBody(ctx context.Context, url string, chanWrite chan<- *string) {
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

func crawlWeb(ctx context.Context, chanIn chan string, wg *sync.WaitGroup) chan *string {
	chanWrite := make(chan *string, 50)
	for {
		select {
		case <-ctx.Done():
			close(chanWrite)
			break
		case value, ok := <-chanIn:
			if ok {
				wg.Add(1)
				go func(str string) {
					defer wg.Done()
					getBody(ctx, str, chanWrite)
				}(value)
			} else {
				return chanWrite
			}
		}
	}

}

func sendUrls(ctx context.Context, urls *[]string, ch chan<- string) {
	for _, url := range *urls {
		switch ctx.Err() != nil {

		}
		time.Sleep(time.Duration(1*500) * time.Millisecond)
		ch <- url
	}
	close(ch)
}

func handlSignals(chW *chan string) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT)
	for {
		sig := <-sigs
		switch sig {
		case os.Interrupt:
			fmt.Println("\nexiting")
			close(*chW)
			//			 os.Exit(0)
			// cancel()
			return
		}
	}
}

// func handlSignals(cancel context.CancelFunc) {
// 	sigs := make(chan os.Signal)
// 	signal.Notify(sigs, syscall.SIGINT)
// 	for {
// 		sig := <-sigs
// 		switch sig {
// 		case os.Interrupt:
// 			 fmt.Println("\nexiting")
// 			cancel()
// //			 os.Exit(0)
// 			return
// 		}
// 	}
// }

func startInitial() ([]string, []*string, chan string) {
	urlArray := []string{"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com",
		// "https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com",
		// "https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com",
		"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com",
		"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com",
		"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com"}
	var bodyPtrArray []*string
	chanWrite := make(chan string, 100)
	return urlArray, bodyPtrArray, chanWrite
}

// Main goroutine
func main() {
	runtime.GOMAXPROCS(8)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	urlArray, bodyPtrArray, chanWrite := startInitial()
	go handlSignals(&chanWrite)
	wg := &sync.WaitGroup{}
	go sendUrls(ctx, &urlArray, chanWrite)
	chanRead := crawlWeb(ctx, chanWrite, wg)

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
