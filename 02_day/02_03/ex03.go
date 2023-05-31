package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Run "tar -czvf (archive name).tar.gz (pathtofile)” in the Terminal
// to compress a file or folder.
// tar xf archive.tar.gz  # for .tar.gz files

// tar [-ключи] [название архива, который будет создан] [что паковать\куда паковать]
// с (create) - создать файл архива
// v (verbose) - показать информацию о выполнении
// f (file) - указывает что нужно создавать файл с именем, которое задается после ключей (в нашем примере file.tar или file.tar.gz), если не указать этот ключ, то будет использовано имя по умолчанию или возникнут проблемы.
// z (gzip) - архивировать файл с помощью gzip

func showAllCommands() {
	flag.PrintDefaults()
	os.Exit(0)
}

func executor(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("%s is not exist\n", path)
		os.Exit(1)
	}
	now := time.Now()
	sec := now.Unix()
	nameArh := path[:len(path)-len(filepath.Ext(path))] + "_" + strconv.FormatInt(sec, 10) + ".tar.gz"
	cmd := exec.Command("tar", "-czf", nameArh, path)
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func executorA(path string, pathDest string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("%s is not exist\n", path)
		os.Exit(1)
	}
	if _, err := os.Stat(pathDest); os.IsNotExist(err) {
		fmt.Printf("%s is not exist\n", pathDest)
		os.Exit(1)
	}
	pathArr := strings.Split(path, "/")
	fmt.Println("pathArr[end]: ", pathArr[len(pathArr)-1])
	fileName := pathArr[len(pathArr)-1]
	nameWithOutExt := fileName[:len(fileName)-len(filepath.Ext(fileName))]
	fmt.Println("nameWithOutExt", nameWithOutExt)
	if pathDest[len(pathDest)-1] != '/' {
		pathDest += "/"
	}
	now := time.Now()
	sec := now.Unix()
	nameArh := pathDest + fileName[:len(fileName)-len(filepath.Ext(fileName))] + "_" + strconv.FormatInt(sec, 10) + ".tar.gz"
	fmt.Println("nameArh", nameArh)
	cmd := exec.Command("tar", "-czf", nameArh, path)
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	var (
		aFlag      bool
		printUsage bool
		pathDest   string
		pathArr    []string
	)
	flag.BoolVar(&aFlag, "a", false, "enter the path for the archive")
	flag.BoolVar(&printUsage, "h", false, "show commands")
	flag.Parse()
	if printUsage {
		showAllCommands()
	}

	if aFlag {
		if len(os.Args) < 4 {
			fmt.Println("Error of count of arguments")
			os.Exit(1)
		}
		pathDest = os.Args[2]
		if _, err := os.Stat(pathDest); os.IsNotExist(err) {
			fmt.Printf("%s is not exist\n", pathDest)
			os.Exit(1)
		}
		pathArr = os.Args[3:]
		for _, path := range pathArr {
			executorA(path, pathDest)
		}
	} else {
		pathArr = os.Args[1:]
		for _, path := range pathArr {
			go executor(path)
		}
	}
	time.Sleep(5 * time.Second)
}
