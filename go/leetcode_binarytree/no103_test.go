package binarytree_test

import (
	"fmt"
	binarytree "leetcode/go/leetcode_binarytree"
	"testing"
)

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func zigzagLevelOrder(root *binarytree.TreeNode) [][]int {

	var res [][]int

	if root == nil {
		return res
	}

	treeNodeList := []*binarytree.TreeNode{root}

	for i := 0; len(treeNodeList) > 0; i++ {
		lenTreeNodeList := len(treeNodeList)
		ans := make([]int, 0)

		for _, treeNode := range treeNodeList {
			ans = append(ans, treeNode.Val)

			if treeNode.Left != nil {
				treeNodeList = append(treeNodeList, treeNode.Left)
			}

			if treeNode.Right != nil {
				treeNodeList = append(treeNodeList, treeNode.Right)
			}
		}
		if i%2 == 1 {
			for i, n := 0, len(ans); i < n/2; i++ {
				ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
			}
		}
		res = append(res, ans)
		treeNodeList = treeNodeList[lenTreeNodeList:]
	}
	return res

	// var res [][]int

	// if root == nil {
	// 	return res
	// }

	// queue := make([]*TreeNode, 0)
	// queue = append(queue, root)
	// isLeftStart := true

	// for len(queue) > 0 {
	// 	queueLength := len(queue)
	// 	fmt.Println(queueLength)
	// 	ans := make([]int, queueLength)

	// 	for i := 0; i < queueLength; i++ {
	// 		node := queue[i]
	// 		if node.Left != nil {
	// 			queue = append(queue, node.Left)
	// 		}
	// 		if node.Right != nil {
	// 			queue = append(queue, node.Right)
	// 		}
	// 		if isLeftStart {
	// 			ans[i] = node.Val
	// 		} else {
	// 			ans[queueLength-i-1] = node.Val
	// 		}
	// 	}
	// 	res = append(res, ans)
	// 	isLeftStart = !isLeftStart
	// 	queue = queue[queueLength:]
	// }
	// return res
}
func TestFunc103(t *testing.T) {

	// var root *binarytree.TreeNode
	// // [3,9,20,null,null,15,7],
	// root = &binarytree.TreeNode{
	// 	Val: 3,
	// 	Left: &binarytree.TreeNode{
	// 		Val:   9,
	// 		Left:  nil,
	// 		Right: nil,
	// 	},
	// 	Right: &binarytree.TreeNode{
	// 		Val: 20,
	// 		Left: &binarytree.TreeNode{
	// 			Val:   15,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 		Right: &binarytree.TreeNode{
	// 			Val:   7,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 	},
	// }

	res := zigzagLevelOrder(binarytree.Root)

	fmt.Println(res)
}
