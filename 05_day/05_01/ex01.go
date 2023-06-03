package main

import (
	"fmt"
)

func reverseArr(tempArr []*TreeNode) {
	for left, right := 0, len(tempArr)-1; left < right; left, right = left+1, right-1 {
		tempArr[left], tempArr[right] = tempArr[right], tempArr[left]
	}
}

func (root *TreeNode) getLevel() int {
	if root == nil {
		return 0
	}
	return root.Left.getLevel() + root.Right.getLevel() + 1
}

func (root *TreeNode) unrollGarland() {
	if root == nil {
		return
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	rez := []bool{}
	count := 0
	for len(queue) > 0 {
		tempArr := []*TreeNode{}
		for _, x := range queue {
			rez = append(rez, x.HasToy)
			if x.Left != nil {
				tempArr = append(tempArr, x.Left)
			}
			if x.Right != nil {
				tempArr = append(tempArr, x.Right)
			}
		}
		if count%2 != 0 {
			reverseArr(tempArr)
		}
		queue = tempArr
		count++
	}
	fmt.Println(rez)
}
