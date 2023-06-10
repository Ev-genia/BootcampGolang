package main

import (
	//	"fmt"
	"fmt"
	"net/http"

	//	"net/url"
	"io/ioutil"
	"log"
	"runtime"
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
	//	log.Printf(str)
	chanWrite <- &str
}

func crawlWeb(chanIn chan string) chan *string {
	chanWrite := make(chan *string, 50)
	for value := range chanIn {
		// fmt.Printf(value)
		go getBody(value, chanWrite)
	}
	return chanWrite
}

func sendUrls(urls *[]string, ch chan<- string) {
	for _, url := range *urls {
		ch <- url
	}
	close(ch)
}

// Main goroutine
func main() {

	runtime.GOMAXPROCS(8)
	urlArray := []string{"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com",
		"https://www.itcodet.com", "https://pkg.go.dev/std", "https://golangify.com"}
	var bodyPtrArray []*string
	chanWrite := make(chan string, 50)

	go sendUrls(&urlArray, chanWrite)

	chanRead := crawlWeb(chanWrite)
	for range urlArray {
		bodyPtrArray = append(bodyPtrArray, <-chanRead)
	}
	close(chanRead)

	for _, Ptr := range bodyPtrArray {
		fmt.Println("\n", Ptr)
	}
}
