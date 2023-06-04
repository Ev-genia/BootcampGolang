package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type sFlag struct {
	symbolicLink bool
	dir          bool
	file         bool
	allFlags     bool
}
type FileMode uint32

func showAllCommands() {
	flag.PrintDefaults()
	os.Exit(0)
}

func recursiveFind(pathArg string, typeFind string, extArg string) {
	filepath.WalkDir(pathArg, func(path string, file fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fi, err := file.Info()
		if file.IsDir() && (typeFind == "d" || typeFind == "") {
			fmt.Println(path)
		}
		if !file.IsDir() && ((typeFind == "f" || typeFind == "") || (typeFind == "d" && extArg != "")) {
			if strings.EqualFold(extArg, "") {
				fmt.Println(path)
			} else if strings.EqualFold(filepath.Ext(path), extArg) && extArg != "" {
				fmt.Println(path)
			}
		}
		if (fi.Mode()&fs.ModeSymlink != 0) && (typeFind == "sl" || typeFind == "") {
			eval, err := filepath.EvalSymlinks(fi.Name())
			if err != nil && os.IsNotExist(err) {
				fmt.Printf("%s -> [broken]\n", path)
			} else {
				fmt.Println("%s -> %s\n", path, eval)
			}
		}
		return nil
	})
}

func main() {
	var (
		flags      = sFlag{false, false, false, true}
		printUsage bool
		extArg     string
		pathArg    string
	)
	flag.BoolVar(&flags.symbolicLink, "sl", false, "find symbolic link")
	flag.BoolVar(&flags.dir, "d", false, "find directory")
	flag.BoolVar(&flags.file, "f", false, "find file")
	flag.StringVar(&extArg, "ext", "", "find only files with a certain extension, work only with -f")
	flag.BoolVar(&printUsage, "h", false, "show all commands")
	flag.Parse()
	if printUsage {
		showAllCommands()
	}
	args := os.Args[1:]
	if len(args) > 0 {
		pathArg = args[len(args)-1]
	} else {
		pathArg = "."
	}
	if _, err := os.Stat(pathArg); os.IsNotExist(err) {
		fmt.Printf("%s is not exist\n", pathArg)
	}
	if len(args) == 1 {
		recursiveFind(pathArg, "", extArg)
	} else {
		flags.allFlags = !(flags.symbolicLink || flags.dir || flags.file)
		if flags.allFlags || flags.dir {
			recursiveFind(pathArg, "d", extArg)
		}
		if !flags.file {
			extArg = ""
		}
		if !(strings.EqualFold(extArg, "")) {
			extArg = "." + extArg
		}
		if flags.allFlags || flags.file || (extArg != "" && flags.file) {
			recursiveFind(pathArg, "f", extArg)
		}
		if flags.allFlags || flags.symbolicLink {
			recursiveFind(pathArg, "sl", extArg)
		}
	}
}
