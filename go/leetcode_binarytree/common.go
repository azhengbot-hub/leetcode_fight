package binarytree

// TreeNode treenode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Root *binarytree.TreeNode
// [3,9,20,null,null,15,7],
var Root = &TreeNode{
	Val: 3,
	Left: &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	},
	Right: &TreeNode{
		Val: 20,
		Left: &TreeNode{
			Val:   15,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	},
}
