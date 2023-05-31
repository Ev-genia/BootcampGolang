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
	fileName := pathArr[len(pathArr)-1]
	if pathDest[len(pathDest)-1] != '/' {
		pathDest += "/"
	}
	now := time.Now()
	sec := now.Unix()
	nameArh := pathDest + fileName[:len(fileName)-len(filepath.Ext(fileName))] + "_" + strconv.FormatInt(sec, 10) + ".tar.gz"
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

	if len(os.Args) < 2 {
		fmt.Println("Error of count of arguments")
		os.Exit(1)
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
			go executorA(path, pathDest)
		}
	} else {
		pathArr = os.Args[1:]
		for _, path := range pathArr {
			go executor(path)
		}
	}
	time.Sleep(5 * time.Second)
}
