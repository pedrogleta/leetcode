package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	items []int
}

func inorderTraversal(root *TreeNode) []int {

	stack := Stack{}
	stack.traverseInorderly(root)
	return stack.items
}

func (s *Stack) traverseInorderly(node *TreeNode) {
	for node == nil {
		return
	}

	s.traverseInorderly(node.Left)
	s.items = append(s.items, node.Val)
	s.traverseInorderly(node.Right)
}
