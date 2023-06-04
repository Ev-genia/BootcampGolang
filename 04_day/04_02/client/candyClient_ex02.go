package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type FlagArg struct {
	candy     string
	count     int
	money     int
	printHelp bool
}

type sPropertiesS struct {
	Change int    `json:"change"`
	Thanks string `json:"thanks"`
}

type sPropertiesF struct {
	Error string `json:"error"`
}

// Handling function
func buyCandy(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "error with ReadAll: %s\n", err)
	}
	fmt.Fprintf(w, "Body: %s\n", body)
}

// Print help
func printhelp() {
	fmt.Println("***\nThis program sends a request to the server, indicating with the help of flags the type of candy, the number of candies and the amount of money.")
	fmt.Println("mandatory flags:")
	fmt.Println("-k\t candy type")
	fmt.Println("-c\t number of candy")
	fmt.Println("-m\t amount of money")
	fmt.Println("-h\t show help")
	fmt.Println("Example: ./askCowClient -k AA -c 2 -m 50\n***")
	os.Exit(0)
}

// Flags handling
func initflags(flags *FlagArg) {

	flag.StringVar(&flags.candy, "k", "", "candy type")
	flag.IntVar(&flags.count, "c", 0, "number of candy")
	flag.IntVar(&flags.money, "m", 0, "money")
	flag.BoolVar(&flags.printHelp, "h", false, "show help")
	flag.Parse()

	if flags.printHelp {
		printhelp()
	}
}

// Reading certificate
func getCert(certfile, keyfile string) (c tls.Certificate, err error) {
	if certfile != "" && keyfile != "" {
		c, err = tls.LoadX509KeyPair(certfile, keyfile)
		if err != nil {
			fmt.Printf("Error loading key pair: %v\n", err)
		}
	} else {
		err = fmt.Errorf("I have no certificate")
	}
	return
}

// ClientCertReqFunc returns a function for tlsConfig.GetClientCertificate
func ClientCertReqFunc(certfile, keyfile string) func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
	c, err := getCert(certfile, keyfile)

	return func(certReq *tls.CertificateRequestInfo) (*tls.Certificate, error) {
		// fmt.Println("Received certificate request: sending certificate")
		if err != nil || certfile == "" {
			fmt.Println("I have no certificate")
		} else {
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		}
		return &c, nil
	}
}

// Create Http.Client
func getClient() *http.Client {
	cp := x509.NewCertPool()
	data, _ := ioutil.ReadFile("../cert/minica.pem")
	cp.AppendCertsFromPEM(data)

	config := &tls.Config{
		RootCAs:              cp,
		GetClientCertificate: ClientCertReqFunc("../cert/client/cert.pem", "../cert/client/key.pem"),
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}

// Print error
func checkError(err error) {
	if err != nil {
		fmt.Printf("Client error: %v\n", err)
		os.Exit(1)
	}
}

func printResponse(body []byte, statusCode int) {
	var successResp sPropertiesS
	var failedResp sPropertiesF

	switch statusCode {
	case 201:
		checkError(json.Unmarshal(body, &successResp))
		fmt.Printf("%s Your change is %d\n", successResp.Thanks, successResp.Change)
	default:
		checkError(json.Unmarshal(body, &failedResp))
		fmt.Printf("%s\n", failedResp.Error)
	}
}

func main() {
	var flags FlagArg

	initflags(&flags)
	client := getClient()
	s := fmt.Sprintf("{\"money\": %d, \"candyType\": \"%s\", \"candyCount\": %d}", flags.money, flags.candy, flags.count)
	resp, err := client.Post("https://candy.tld:3333/buy_candy", "application/json", bytes.NewBufferString(s))
	checkError(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed reading response body: %s", err)
	}
	printResponse(body, resp.StatusCode)
}
