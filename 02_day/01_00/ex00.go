package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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

func recursiveFind(pathArg string, typeFind string) {
	filepath.WalkDir(pathArg, func(path string, file fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if file.IsDir() && (typeFind == "d" || typeFind == "") {
			fmt.Println(path)
		}
		if !file.IsDir() && (typeFind == "f" || typeFind == "") {
			fmt.Println(path)
		}
		fi, err := file.Info()
		if (fi.Mode()&fs.ModeSymlink != 0) && typeFind == "sl" {
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
	flag.StringVar(&extArg, "ext", "", "find only files with a certain extension")
	flag.BoolVar(&printUsage, "h", false, "show all commands")
	flag.Parse()
	if printUsage {
		showAllCommands()
	}
	args := os.Args[1:]
	pathArg = args[len(args)-1]
	if len(args) == 1 {
		recursiveFind(pathArg, "")
	} else {
		flags.allFlags = !(flags.symbolicLink || flags.dir || flags.file)
		if flags.allFlags || flags.dir {
			recursiveFind(pathArg, "d")
		}
		if flags.allFlags || flags.file {
			recursiveFind(pathArg, "f")
		}
		if flags.allFlags || flags.symbolicLink {
			recursiveFind(pathArg, "sl")
		}
	}
}
