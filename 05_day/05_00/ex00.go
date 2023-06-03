package main

func (root *TreeNode) GetTreeNodeVal() int {
	if root == nil {
		return 0
	}
	ret := 0
	if root.HasToy {
		ret = 1
	}
	return root.Left.GetTreeNodeVal() + root.Right.GetTreeNodeVal() + ret
}

func (root *TreeNode) areToysBalanced() bool {
	left := root.Left.GetTreeNodeVal()
	right := root.Right.GetTreeNodeVal()
	return left == right
}
