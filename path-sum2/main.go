package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Right == nil && root.Left == nil && targetSum == root.Val {
		return [][]int{[]int{root.Val}}
	}
	targetSum -= root.Val
	arra1 := []int{root.Val}
	arra2 := []int{root.Val}
	var allArra [][]int
	if !hasPathSum(root.Left, targetSum, &arra1) || !hasPathSum(root.Right, targetSum, &arra2) {
		return [][]int{}
	}
	allArra = append(allArra, arra1, arra2)
	return allArra
}

func hasPathSum(root *TreeNode, targetSum int, arra *[]int) bool {
	if root == nil {
		return false
	}

	if targetSum-root.Val == 0 && root.Right == nil && root.Left == nil {
		*arra = append(*arra, root.Val)
		return true
	}
	targetSum -= root.Val

	if root.Left != nil || root.Right != nil {
		*arra = append(*arra, root.Val)
	}

	return hasPathSum(root.Right, targetSum, arra) || hasPathSum(root.Left, targetSum, arra)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Val == head.Next.Val {
		return true
	}
	if head.Next.Next == nil || head.Next.Next.Next == nil {
		return false
	}
	if head.Val == head.Next.Next.Next.Val {
		return true
	}

	return hasCycle(head.Next)

}
