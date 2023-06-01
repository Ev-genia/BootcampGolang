package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

type FlagArg struct {
	line      bool
	word      bool
	char      bool
	printHelp bool
	pathArg   string
	args      []string
}

// Подсчет
func countItems(file string, flags FlagArg) {

	// Открытие файла и инициализация сканера
	fileHandle, err := os.Open(file)
	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	// выбор функции подсчета
	if flags.char {
		fileScanner.Split(bufio.ScanRunes)
	} else if flags.line {
		fileScanner.Split(bufio.ScanLines)
	} else if flags.word {
		fileScanner.Split(bufio.ScanWords)
	}
	//  подсчет
	count := 0
	for fileScanner.Scan() {
		count++
	}
	if err := fileScanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d\t%s\n", count, file)
}

// Выполняем условия для флагов по сабджекту
func initFlagCondition(flags *FlagArg) {

	//  Ставим флаг по умолчанию
	if !((*flags).line || (*flags).word || (*flags).char) {
		(*flags).word = true
	}
	//  Обрабатываем множественные флаги
	if ((*flags).line && (*flags).word) ||
		((*flags).word && (*flags).char) ||
		((*flags).line && (*flags).char) {
		fmt.Println("Error: too many flags: flag -w is accepted")
		(*flags).word = true
		(*flags).line = false
		(*flags).char = false
	}
	//  Находим все имена файлов,  обрабатывам отсутствие имен файлов в аргументах
	(*flags).args = os.Args[1:]
	for i := 0; i < len((*flags).args); i++ {
		if (*flags).args[i][0] != '-' {
			(*flags).args = (*flags).args[i:]
			break
		} else if i == len((*flags).args)-1 {
			fmt.Println("Error: the file name is missing.")
			os.Exit(0)
		}
	}
}

// Инициализируем флаги и выводим хелп
func initflags(flags *FlagArg) {

	flag.BoolVar(&flags.line, "l", false, "подсчет строк")
	flag.BoolVar(&flags.word, "w", false, "подсчет слов (по умолчанию)")
	flag.BoolVar(&flags.char, "m", false, "подсчет символов")
	flag.BoolVar(&flags.printHelp, "h", false, "показать Хелп")
	flag.Parse()

	if (*flags).printHelp {
		fmt.Println("***\nПрограмма подсчитывает и выводит кол-во строк, слов и символов в зависимости от флага.")
		fmt.Println("доступные флаги:")
		fmt.Println("-w\t подсчет слов")
		fmt.Println("-l\t подсчет строк")
		fmt.Println("-m\t подсчет символов")
		fmt.Println("-h\t показать данный текст")
		fmt.Println("\nФайлы для подсчета указываются строго после флагов")
		fmt.Println("Без указания флагов посчитывает слова (-w). Без указания файла(ов) завершает работу.")
		fmt.Println("ПРИМЕР: ./myWc -w first_file.log second_file.txt\n***")
		os.Exit(0)
	}
}

// Основной поток
func main() {

	var flags FlagArg
	initflags(&flags)
	initFlagCondition(&flags)

	for _, file := range flags.args {
		go countItems(file, flags)
	}

	time.Sleep(3 * time.Second)
}
