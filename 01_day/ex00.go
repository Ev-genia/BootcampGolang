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
	ItemName  string `json:"ingredient_name" xml:"itemname"`
	ItemCount string `json:"ingredient_count" xml:"itemcount"`
	ItemUnit  string `json:"ingredient_unit,omitempty" xml:"itemunit"`
}

type sXmlIngredients struct {
	Items []sItem `xml:"item"`
}

type sDataCake struct {
	NameCake       string          `json:"name" xml:"name"`
	TimeCake       string          `json:"time" xml:"stovetime"`
	Ingredients    []sItem         `json:"ingredients" xml:"-"`
	XMLIngredients sXmlIngredients `json:"-" xml:"ingredients"`
}

type sCake struct {
	Recipes []sDataCake `json:"cake" xml:"cake"`
	XmlName xml.Name    `json:"-" xml:"recipes"`
}

func readData(dataFromFile []byte, dataType string) sCake {
	var cakes sCake

	switch dataType {
	case "json":
		cakes = readJson(dataFromFile)
	case "xml":
		cakes = readXml(dataFromFile)
	}
	return cakes
}

type DBReader interface {
	readData(dataFromFile []byte, dataType string) sCake
}

func showAllCommands() {
	flag.PrintDefaults()
	os.Exit(0)
}

func readJson(dataFromFile []byte) sCake {
	var cakes sCake

	if !json.Valid(dataFromFile) {
		fmt.Println("json is not valid")
		os.Exit(1)
	}
	err := json.Unmarshal(dataFromFile, &cakes)
	if err != nil {
		fmt.Println("err in Unmarshl:", err)
		os.Exit(1)
	}
	return cakes
}

func readXml(dataFromFile []byte) sCake {
	var cakes sCake

	err := xml.Unmarshal(dataFromFile, &cakes)
	if err != nil {
		fmt.Println("err in Unmarshl:", err)
		os.Exit(1)
	}
	return cakes
}

func getDataFromFile(fileName string) []byte {
	dataFromFile, err := ioutil.ReadFile(fileName)
	if err != nil && err != io.EOF {
		_, err := fmt.Fprint(os.Stderr, "Error of reading *.xml file\n")
		if err != nil {
			fmt.Println("Error of Fprint")
			os.Exit(1)
		}
	}
	return dataFromFile
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
	dataFromFile := getDataFromFile(fileName)
	if len(fileName) > 3 && fileName[len(fileName)-4:] == "json" {
		cakes = readData(dataFromFile, "json")
		fmt.Println("cakes: ", cakes) //
	} else if len(fileName) > 2 && fileName[len(fileName)-3:] == "xml" {
		cakes = readData(dataFromFile, "xml")
		fmt.Println("cakes: ", cakes) //
	} else {
		fmt.Println("Error of enter of type")
		os.Exit(1)
	}
}
