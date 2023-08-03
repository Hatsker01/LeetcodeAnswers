package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}

	// Step 2: Create the second node with value 2 and attach it to the right of the root node
	root.Right = &TreeNode{Val: 2}

	// Step 3: Create the third node with value 3 and attach it to the left of the second node
	root.Right.Left = &TreeNode{Val: 3}

	fmt.Println(inorderTraversal(root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)

	smth(root, &result)
	return result
}

func smth(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	smth(root.Left, result)
	*result = append(*result, root.Val)
	smth(root.Right, result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
