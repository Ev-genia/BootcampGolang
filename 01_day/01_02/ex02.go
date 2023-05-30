package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type dataStruct struct {
	fileName    string
	dirIndex    int
	isExistFlag bool
}

// Загрузка данных в массив строк
func fileLoad(fileName string) []string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	sarr := strings.Split(string(content), "\n")
	return sarr
}

// Сжатие данных этап 2(второй файл)
func compressNew(sarr []string, dirArr *[]string) []dataStruct {
	fileArr := make([]dataStruct, 0, 10000)
	var tempDataStruct dataStruct
	for _, value := range sarr {
		dir, file := path.Split(value)
		tempDataStruct.fileName = file
		tempDataStruct.dirIndex = -1
		for k := 0; k < len(*dirArr); k++ {
			if strings.EqualFold(dir, (*dirArr)[k]) {
				tempDataStruct.dirIndex = k
				break
			}
		}
		if tempDataStruct.dirIndex == -1 {
			*dirArr = append(*dirArr, dir)
			tempDataStruct.dirIndex = len(*dirArr) - 1
		}
		fileArr = append(fileArr, tempDataStruct)
	}
	return fileArr
}

// Сжатие данных этап 1(первый файл)
func compress(sarr []string) ([]string, []dataStruct) {
	fileArr := make([]dataStruct, 0, 10000)
	dirArr := make([]string, 0, 1000)
	var tempDataStruct dataStruct

	for _, value := range sarr {
		dir, file := path.Split(value)
		tempDataStruct.fileName = file
		tempDataStruct.dirIndex = -1
		for k := 0; k < len(dirArr); k++ {
			if strings.EqualFold(dir, dirArr[k]) {
				tempDataStruct.dirIndex = k
				break
			}
		}
		if tempDataStruct.dirIndex == -1 {
			dirArr = append(dirArr, dir)
			tempDataStruct.dirIndex = len(dirArr) - 1
		}
		fileArr = append(fileArr, tempDataStruct)
	}
	return dirArr, fileArr
}

// Сравнение загруженных данных
func compareFs(Old *[]dataStruct, New *[]dataStruct) {

	for i := 0; i < len(*Old); i++ {
		for j := 0; j < len(*New); j++ {
			if (*Old)[i].dirIndex == (*New)[j].dirIndex &&
				strings.EqualFold((*Old)[i].fileName, (*New)[j].fileName) {
				(*Old)[i].isExistFlag = true
				(*New)[j].isExistFlag = true
				break
			}
		}
	}
}

// Вывод изменений файловой системы
func printChangesFs(Old *[]dataStruct, New *[]dataStruct, dirArr *[]string) {
	k := 0
	for i := 0; i < len(*New); i++ {
		if !((*New)[i].isExistFlag) {
			fmt.Printf("ADDED %s%s\n", (*dirArr)[(*New)[i].dirIndex], (*New)[i].fileName)
			k++
		}
	}
	for i := 0; i < len(*Old); i++ {
		if !((*Old)[i].isExistFlag) {
			fmt.Printf("REMOVED %s%s\n", (*dirArr)[(*Old)[i].dirIndex], (*Old)[i].fileName)
			k++
		}
	}
	if k == 0 {
		fmt.Println("FileSystem Unchanged")
	}
}

func inputFileNameCheck(old *string, new *string) {
	if *old == "" {
		fmt.Println("Enter name of OLD file:")
		fmt.Scanln(old)
	}
	if *new == "" {
		fmt.Println("Enter name of NEW file:")
		fmt.Scanln(new)
	}
}

func showAllCommands() {
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	var (
		oldFile    string
		newFile    string
		printUsage bool
	)

	// Обработка флагов
	flag.StringVar(&oldFile, "old", "", "enter of file ")
	flag.StringVar(&newFile, "new", "", "enter of file ")
	flag.BoolVar(&printUsage, "h", false, "show all commands")
	flag.Parse()
	if printUsage {
		showAllCommands()
	}
	inputFileNameCheck(&oldFile, &newFile)

	//Поочередно загружаем файлы и сжимаем данные
	dirArr, fileArrOld := compress(fileLoad(oldFile))
	fileArrNew := compressNew(fileLoad(newFile), &dirArr)

	//  Сравниваем базы и выводим
	compareFs(&fileArrOld, &fileArrNew)
	printChangesFs(&fileArrOld, &fileArrNew, &dirArr)
}
