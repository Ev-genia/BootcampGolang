package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type sItem struct {
	// XmlIngredients xml.Name `json:"-" xml:"item"`
	ItemName  string `json:"ingredient_name" xml:"itemname"`
	ItemCount string `json:"ingredient_count" xml:"itemcount"`
	ItemUnit  string `json:"ingredient_unit,omitempty" xml:"itemunit"`
}

type sDataCake struct {
	NameCake    string  `json:"name" xml:"name"`
	TimeCake    string  `json:"time" xml:"stovetime"`
	Ingredients []sItem `json:"ingredients" xml:"ingredients"`
	// XmlIngredients xml.Name `json:"-" xml:"item"`
}

type sCake struct {
	Recipes []sDataCake `json:"cake" xml:"cake"`
	XmlName xml.Name    `json:"-" xml:"recipes"`
}

func showAllCommands() {
	flag.PrintDefaults()
	os.Exit(0)
}

func readJson(fileName string) sCake {
	var cakes sCake
	dataFromFile, err := ioutil.ReadFile(fileName)
	if err != nil && err != io.EOF {
		_, err := fmt.Fprint(os.Stderr, "Error of reading *.json file\n")
		if err != nil {
			fmt.Println("Error of Fprint")
			os.Exit(1)
		}
	}
	if !json.Valid(dataFromFile) {
		fmt.Println("json is not valid")
		os.Exit(1)
	}
	err = json.Unmarshal(dataFromFile, &cakes)
	if err != nil {
		fmt.Println("err in Unmarshl:", err)
	}
	return cakes
}

func readXml(fileName string) sCake {
	var cakes sCake

	dataFromFile, err := ioutil.ReadFile(fileName)
	if err != nil && err != io.EOF {
		_, err := fmt.Fprint(os.Stderr, "Error of reading *.xml file\n")
		if err != nil {
			fmt.Println("Error of Fprint")
			os.Exit(1)
		}
	}
	err = xml.Unmarshal(dataFromFile, &cakes)
	if err != nil {
		fmt.Println("err in Unmarshl:", err)
	}
	return cakes
}

func main() {
	var fileName string
	var printUsage bool
	var cakes sCake
	flag.StringVar(&fileName, "f", "", "enter of file ")
	flag.BoolVar(&printUsage, "h", false, "show all commands")
	flag.Parse()
	if printUsage {
		showAllCommands()
	}
	if fileName == "" {
		fmt.Println("Enter name of file")
		fmt.Scan(&fileName)
	}
	// fmt.Println("fileName: ", fileName) //
	if len(fileName) > 3 && fileName[len(fileName)-4:] == "json" {
		// fmt.Println("extention: ", fileName[len(fileName)-4:]) //
		cakes = readJson(fileName)
		fmt.Println("cakes: ", cakes) //
		// enc := json.Encoder(cakes)
	} else if len(fileName) > 2 && fileName[len(fileName)-3:] == "xml" {
		fmt.Println("extention: ", fileName[len(fileName)-3:]) //
		cakes = readXml(fileName)
		fmt.Println("cakes: ", cakes) //
	} else {
		fmt.Println("Error of enter of type")
		os.Exit(1)
	}
}
