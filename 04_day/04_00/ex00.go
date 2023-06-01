package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Required struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

type sResponses struct {
	Description string  `json:"description"`
	Schema      sSchema `json:"schema"`
}

type sSchema struct {
	TypeObj     string       `json:"type"`
	PropertiesS sPropertiesS `json:"properties,omitempty"`
	PropertiesF sPropertiesF `json:"properties,omitempty"`
}

type sPropertiesS struct {
	Change int    `json:"change"`
	Thanks string `json:"thanks"`
}

type sPropertiesF struct {
	Error string `json:"error,omitempty"`
}

type sCandy struct {
	CandyType string
	Price     int
}

func candyArrInit(candyArr []sCandy) {
	candyArr[0].CandyType = "CE"
	candyArr[0].Price = 10
	candyArr[1].CandyType = "AA"
	candyArr[1].Price = 15
	candyArr[2].CandyType = "NT"
	candyArr[2].Price = 17
	candyArr[3].CandyType = "DE"
	candyArr[3].Price = 21
	candyArr[4].CandyType = "YR"
	candyArr[4].Price = 23
}

func success(w http.ResponseWriter, ret sPropertiesS, diff int) {
	ret.Change = diff
	ret.Thanks = "Thank you!"
	byteRestp, err := json.Marshal(ret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error with json.Marshal: %s\n", err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(byteRestp)
}

func getType(candies []sCandy, candyType string) int {
	numType := 0

	for i, val := range candies {
		if val.CandyType == candyType {
			numType = i
			break
		}
	}
	return numType
}

func getData(r *http.Request, elem *Required) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error with ReadAll: %s\n", err)
	}
	err = json.Unmarshal(body, &elem)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error Unmarshal: %s\n", err)
	}
}

func noMoney(w http.ResponseWriter, ret sPropertiesF, diff int) {
	diff *= -1
	respSrt := "You need { " + strconv.Itoa(diff) + " } more money!\n"
	w.WriteHeader(http.StatusPaymentRequired)
	ret.Error = "not enough money"
	byteRestp, err := json.Marshal(ret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error with json.Marshal: %s\n", err)
	}
	fmt.Fprintf(os.Stderr, "%s\n", byteRestp)
	w.Write([]byte(respSrt))
}

func errorCase(w http.ResponseWriter, numType int, ret sPropertiesF) {
	w.WriteHeader(http.StatusBadRequest)
	if numType == 0 {
		ret.Error = "Error type candy\n"
	} else {
		ret.Error = "Error count candy"
	}
	byteRestp, err := json.Marshal(ret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error with json.Marshal: %s\n", err)
	}
	w.Write(byteRestp)
}

func buyCandy(w http.ResponseWriter, r *http.Request) {
	var elem Required
	var retS sPropertiesS
	var retF sPropertiesF

	candies := make([]sCandy, 5, 5)
	candyArrInit(candies)
	getData(r, &elem)
	numType := getType(candies, elem.CandyType)
	w.Header().Set("Content-Type", "application/json")
	if numType != 0 && elem.CandyCount > 0 {
		diff := elem.Money - elem.CandyCount*candies[numType].Price
		if diff >= 0 {
			success(w, retS, diff)
		} else {
			noMoney(w, retF, diff)
		}
	} else {
		errorCase(w, numType, retF)
	}
}

func main() {
	http.HandleFunc("/buy_candy", buyCandy)
	log.Fatal(http.ListenAndServe("localhost:3333", nil))
}
