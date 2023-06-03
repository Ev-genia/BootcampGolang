package main

func main() {
	root := createNode(true)            //(1)            //
	root.Left = createNode(true)        //(2)        //
	root.Right = createNode(false)      //(3)       //
	root.Left.Left = createNode(true)   //(6)   //
	root.Left.Right = createNode(false) //(4)  //
	root.Right.Left = createNode(true)  //(5)  //
	root.Right.Right = createNode(true) //(7) //
	root.unrollGarland()
}
