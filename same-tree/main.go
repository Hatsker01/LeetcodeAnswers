package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	

	if p.Val != q.Val {
		return false
	}
	if p.Right.Val!=q.Right.Val{
		return false
	}
	if p.Left.Val!=q.Left.Val{
		return false
	}

	return true
}
