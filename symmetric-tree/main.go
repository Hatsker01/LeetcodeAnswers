package main

func main() {

}
func mirror(left *TreeNode, right *TreeNode) bool {
	if right == nil && left == nil {
		return true
	}

	if right == nil || left == nil {
		return false
	}

	return right.Val == left.Val && mirror(right.Left, left.Right) && mirror(right.Right, left.Left)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return mirror(root.Left, root.Right)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
