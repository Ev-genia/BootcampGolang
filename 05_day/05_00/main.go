package main

import "fmt"

func main() {
	root := createNode(false)           //(1)
	root.Left = createNode(true)        //(2)
	root.Right = createNode(true)       //(3)
	root.Left.Left = createNode(false)  //(6)
	root.Left.Right = createNode(true)  //(4)
	root.Right.Left = createNode(false) //(5)
	root.Right.Right = createNode(true) //(7)
	rez := root.areToysBalanced()
	fmt.Println("Equal: ", rez)
}
