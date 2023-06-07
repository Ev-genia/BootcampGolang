package main

import (
	"log"
	"net/http"
	"time"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
)

func main() {
	lmt := tollbooth.NewLimiter(10, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	lmt.SetMessage("429 Too Many Requests")
	http.Handle("/", tollbooth.LimitFuncHandler(lmt, home))

	http.Handle("/admin", tollbooth.LimitFuncHandler(lmt, admin))
	http.Handle("/snippet/create", tollbooth.LimitFuncHandler(lmt, createSnippet))
	http.Handle("/snippet/write", tollbooth.LimitFuncHandler(lmt, writeSnippet))
	http.Handle("/showpost", tollbooth.LimitFuncHandler(lmt, showSnippet))

	fileServer := http.FileServer(http.Dir("./"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Go Backend: { HTTPVersion = 1 }; serving on http://localhost:8888/")
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}
