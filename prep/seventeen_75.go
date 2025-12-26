package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
   res := []int{}

	 var traverse func(node *TreeNode)
	 traverse = func(node *TreeNode) {
		if node == nil {
			return;
		}

		traverse(node.Left)
		res = append(res, node.Val)
		traverse(node.Right)
	 }
	 traverse(root)
	 
	 return res
}