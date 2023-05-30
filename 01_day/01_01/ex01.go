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

type DBreader interface {
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

func getXml(fileName string, x xmlRead, j jsonRead) Recipes {
	var cakesXml Recipes
	if len(fileName) > 3 && fileName[len(fileName)-4:] == "json" {
		cakesXml = j.readDb(getDataFromFile(fileName))
	} else if len(fileName) > 2 && fileName[len(fileName)-3:] == "xml" {
		cakesXml = x.readDb(getDataFromFile(fileName))
	} else {
		fmt.Println("Error of enter of type")
		os.Exit(1)
	}

	return cakesXml
}

func checkRemoveCake(cakeOld []string, cakeNew []string, rez *[]string) {
	for i := 0; i < len(cakeOld); i++ {
		countCake := 0
		for j := 0; j < len(cakeNew); j++ {
			if cakeOld[i] == cakeNew[j] {
				countCake++
			}
		}
		if countCake == 0 {
			*rez = append(*rez, "REMOVED cake \""+cakeOld[i]+"\"\n")
		}
	}
}

func checkAddCake(cakeOld []string, cakeNew []string, rez *[]string) {
	for i := 0; i < len(cakeNew); i++ {
		countCake := 0
		for j := 0; j < len(cakeOld); j++ {
			if cakeNew[i] == cakeOld[j] {
				countCake++
			}
		}
		if countCake == 0 {
			*rez = append(*rez, "ADDED cake \""+cakeNew[i]+"\"\n")
		}
	}
}

func checkAddIngr(ingOld []string, ingNew []string, nameCake string, rez *[]string) {
	for i := 0; i < len(ingNew); i++ {
		countCake := 0
		for j := 0; j < len(ingOld); j++ {
			if ingNew[i] == ingOld[j] {
				countCake++
			}
		}
		if countCake == 0 {
			*rez = append(*rez, "ADDED ingredient \""+ingNew[i]+"\" for cake \""+nameCake+"\"\n")
		}
	}
}

func checkRemoveIngr(ingOld []string, ingNew []string, nameCake string, rez *[]string) {
	for i := 0; i < len(ingOld); i++ {
		countCake := 0
		for j := 0; j < len(ingNew); j++ {
			if ingOld[i] == ingNew[j] {
				countCake++
			}
		}
		if countCake == 0 {
			*rez = append(*rez, "REMOVED ingredient \""+ingOld[i]+"\" for cake  \""+nameCake+"\"\n")
		}
	}
}

func checkItems(itemOld Item, itemNew Item, nameCake string, rez *[]string) {
	if itemOld.Count != itemNew.Count {
		*rez = append(*rez, "CHANGED unit count for ingredient \""+itemOld.Name+"\" for cake \""+nameCake+"\" - \""+itemNew.Count+"\" instead of\""+itemOld.Count+"\"\n")
	}
	if itemOld.Unit != itemNew.Unit {
		*rez = append(*rez, "CHANGED unit for ingredient \""+itemOld.Name+"\" for cake \""+nameCake+"\" - \""+itemNew.Unit+"\" instead of \""+itemOld.Unit+"\"\n")
	}
}

func checkIngridients(oldX Cake, newX Cake, rez *[]string) {
	if oldX.TimeCooking != newX.TimeCooking {
		*rez = append(*rez, "CHANGED cooking time for cake \""+oldX.Name+"\" - \""+newX.TimeCooking+"\" instead of \""+oldX.TimeCooking+"\"\n")
	}
	var ingrOld []string
	var ingrNew []string
	for _, i := range oldX.Ingredients.Items {
		ingrOld = append(ingrOld, i.Name)
	}
	for _, i := range newX.Ingredients.Items {
		ingrNew = append(ingrNew, i.Name)
	}

	checkRemoveIngr(ingrOld, ingrNew, oldX.Name, rez)
	checkAddIngr(ingrOld, ingrNew, oldX.Name, rez)
	for i, itemOld := range oldX.Ingredients.Items {
		for j, itemNew := range newX.Ingredients.Items {
			if oldX.Ingredients.Items[i].Name == newX.Ingredients.Items[j].Name {
				checkItems(itemOld, itemNew, oldX.Name, rez)
			}
		}
	}
}

func comparison(cakesXmlOld Recipes, cakesXmlNew Recipes) {
	var rez []string
	var cakeOld []string
	var cakeNew []string

	for _, oldX := range cakesXmlOld.RecipesXml {
		cakeOld = append(cakeOld, oldX.Name)
	}
	for _, newX := range cakesXmlNew.RecipesXml {
		cakeNew = append(cakeNew, newX.Name)
	}
	checkRemoveCake(cakeOld, cakeNew, &rez)
	checkAddCake(cakeOld, cakeNew, &rez)
	for _, oldX := range cakesXmlOld.RecipesXml {
		for _, newX := range cakesXmlNew.RecipesXml {
			if oldX.Name == newX.Name {
				checkIngridients(oldX, newX, &rez)
			}
		}
	}
	for _, r := range rez {
		fmt.Print(r)
	}
}

func showAllCommands() {
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	var fileNameOld string
	var fileNameNew string
	x := xmlRead{}
	j := jsonRead{}
	var printUsage bool

	// Обработка флагов
	flag.StringVar(&fileNameOld, "old", "", "enter old base")
	flag.StringVar(&fileNameNew, "new", "", "enter new base")
	flag.BoolVar(&printUsage, "h", false, "show all commands")
	flag.Parse()
	if printUsage {
		showAllCommands()
	}
	if fileNameOld == "" {
		fmt.Println("Enter name of old base:")
		fmt.Scan(&fileNameOld)
	}
	if fileNameNew == "" {
		fmt.Println("Enter name of new base:")
		fmt.Scan(&fileNameNew)
	}
	cakesXmlOld := getXml(fileNameOld, x, j)
	cakesXmlNew := getXml(fileNameNew, x, j)
	comparison(cakesXmlOld, cakesXmlNew)
}
