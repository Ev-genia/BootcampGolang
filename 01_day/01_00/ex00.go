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

// type Recipes struct {
// 	XMLName xml.Name 		`xml:"recipes"`
// 	RecipesXml [] struct {
// 		Name        string 	`xml:"name"`
// 		TimeCooking	string	`xml:"stovetime"`
// 		Ingredients  struct {
// 			Items [] struct {
// 				Name	string	`xml:"itemname"`
// 				Count	string	`xml:"itemcount"`
// 				Unit	string	`xml:"itemunit"`
// 			}`xml:"item"`
// 		}`xml:"ingredients"`
// 	}`xml:"cake"`
// }

type JsonRecipes struct {
	Cake []JsonCake `json:"cake"`
}

type JsonCake struct {
	Name        string `json:"name"`
	TimeCooking string `json:"time"`
	Items       []Item `json:"ingredients"`
}

type Item struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit,omitempty" xml:"itemunit"`
}

type IngredientsList struct {
	Items []Item `xml:"item"`
}

type Cake struct {
	Name        string          `xml:"name"`
	TimeCooking string          `xml:"stovetime"`
	Ingredients IngredientsList `xml:"ingredients"`
}

type Recipes struct {
	XMLName    xml.Name `xml:"recipes"`
	RecipesXml []Cake   `xml:"cake"`
}

type DBReader interface {
	readDb([]byte) Recipes
}

type jsonRead struct{}

type xmlRead struct{}

func (j *jsonRead) readDb(dataFromFile []byte) Recipes {
	var cakesJson JsonRecipes
	if !json.Valid(dataFromFile) {
		fmt.Println("json is not valid")
		os.Exit(1)
	}
	err := json.Unmarshal(dataFromFile, &cakesJson)
	if err != nil {
		fmt.Println("err in Unmarshl:", err)
		os.Exit(1)
	}
	cakes := convertJsonRecipeToXmlRecipe(cakesJson)
	return cakes
}

func convertJsonRecipeToXmlRecipe(jr JsonRecipes) Recipes {
	var cakes Recipes

	for _, temp := range jr.Cake {
		tempXmlCake := Cake{}
		tempXmlCake.Name = temp.Name
		tempXmlCake.TimeCooking = temp.TimeCooking
		tempXmlCake.Ingredients.Items = temp.Items
		cakes.RecipesXml = append(cakes.RecipesXml, tempXmlCake)
	}
	return cakes
}

func convertXmlRecipeToJsonRecipe(xr Recipes) JsonRecipes {
	var cakes JsonRecipes

	for _, temp := range xr.RecipesXml {
		tempJsonCake := JsonCake{}
		tempJsonCake.Name = temp.Name
		tempJsonCake.TimeCooking = temp.TimeCooking
		tempJsonCake.Items = temp.Ingredients.Items
		cakes.Cake = append(cakes.Cake, tempJsonCake)
	}
	return cakes
}

func (x *xmlRead) readDb(dataFromFile []byte) Recipes {
	var cakes Recipes

	err := xml.Unmarshal(dataFromFile, &cakes)
	if err != nil {
		fmt.Println("err in Unmarshl:", err)
		os.Exit(1)
	}
	return cakes
}

func pushDataToFile(data []byte, fileNameWrite string) {
	err := ioutil.WriteFile(fileNameWrite, data, 0777)
	if err != nil && err != io.EOF {
		_, err := fmt.Fprint(os.Stderr, "Error at writing\n")
		if err != nil {
			fmt.Println("Error of Writefile")
			os.Exit(1)
		}
	}
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

func writeRecipe(recipes Recipes, typeWrite string, fileName string) {
	if typeWrite == "Json" {
		data, errwrite := json.MarshalIndent(convertXmlRecipeToJsonRecipe(recipes), "", "    ")
		if errwrite == nil {
			pushDataToFile(data, fileName)
		} else {
			fmt.Println("Error at Marshalling to Json")
		}
	} else if typeWrite == "Xml" {
		data, errwrite := xml.MarshalIndent(recipes, "", "    ")
		if errwrite == nil {
			pushDataToFile(data, fileName)
		} else {
			fmt.Println("Error at Marshalling to Xml")
		}
	}
}

func showAllCommands() {
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	var (
		cakesXml   Recipes
		fileName   string
		x          = xmlRead{}
		j          = jsonRead{}
		printUsage bool
	)
	// Обработка флагов
	flag.StringVar(&fileName, "f", "", "enter of file ")
	flag.BoolVar(&printUsage, "h", false, "show all commands")
	flag.Parse()
	if printUsage {
		showAllCommands()
	}
	if fileName == "" {
		fmt.Println("Enter name of file:")
		fmt.Scan(&fileName)
	}
	if len(fileName) > 3 && fileName[len(fileName)-4:] == "json" {
		cakesXml = j.readDb(getDataFromFile(fileName))
	} else if len(fileName) > 2 && fileName[len(fileName)-3:] == "xml" {
		cakesXml = x.readDb(getDataFromFile(fileName))
	} else {
		fmt.Println("Error of enter of type")
		os.Exit(1)
	}
	writeRecipe(cakesXml, "Json", "outRecipe.json")
	writeRecipe(cakesXml, "Xml", "outRecipe.xml")
}
