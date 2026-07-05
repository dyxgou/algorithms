package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	s := NewStack[*TreeNode]()

	s.Insert(root)

	for !s.IsEmpty() {
		cur := s.Pop()

		if cur.Left != nil {
			s.Insert(cur.Left)
		}

		if cur.Right != nil {
			s.Insert(cur.Right)
		}

	}

	return 0
}
