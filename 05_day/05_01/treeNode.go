package main

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func createNode(val bool) *TreeNode {
	node := TreeNode{val, nil, nil}
	return &node
}
