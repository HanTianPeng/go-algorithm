/*
Author: Conk
Function:
CreateTime: 2021/5/24 5:47 下午
UpdateTime: 2021/5/24 5:47 下午
*/

package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// PreOrderTraversal: 前序遍历<根 -> 左 -> 右>: A -> B -> D -> E -> C-> F
func PreOrderTraversal(root *TreeNode) []int {
	var result []int
	var preOrderTraversalNode func(*TreeNode)
	preOrderTraversalNode = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 写入值
		result = append(result, node.Val)
		// 遍历左子树
		preOrderTraversalNode(node.Left)
		// 遍历右子树
		preOrderTraversalNode(node.Right)
	}
	preOrderTraversalNode(root)
	return result
}

func main() {

}