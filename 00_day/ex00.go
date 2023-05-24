package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scan() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Error of enter", err)
	}
	return str
}

func main() {
	str := scan()
	fmt.Print("lines: ", str) //
	// nbrCount := make(map[int]int)
	words := strings.Split(str, "\\n")
	fmt.Println("len: ", len(words))
	for i := 0; i < len(words); i++ {
		fmt.Println(words[i])
	}

}
