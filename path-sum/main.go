package main

import "fmt"

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 11}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right.Right.Right = &TreeNode{Val: 1}

	targetSum := 55
	fmt.Println(hasPathSum(root, targetSum)) // Output: false
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val

	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}

	return hasPathSum(root.Right, targetSum) || hasPathSum(root.Left, targetSum)
}

// func Sum(right, left *TreeNode, sum *int) {
// 	if right == nil && left == nil {
// 		return
// 	}
// 	if right == nil {
// 		*sum += left.Val
// 		Sum(left.Right, left.Left, sum)
// 		return
// 	} else if left == nil {
// 		*sum += right.Val
// 		Sum(right.Right, right.Left, sum)
// 		return
// 	} else {
// 		*sum += right.Val + left.Val
// 	}
// 	Sum(right.Right, right.Left, sum)
// 	Sum(left.Right, left.Left, sum)
// }

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}
