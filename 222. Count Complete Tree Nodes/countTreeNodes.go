package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	visited := []*TreeNode{}

	if root == nil {
		return 0
	}
	return len(recurse(root, visited))
}

func recurse(root *TreeNode, visited []*TreeNode) []*TreeNode {

	visited = append(visited, root)

	if root.Left != nil {
		visited = recurse(root.Left, visited)
	}
	if root.Right != nil {
		visited = recurse(root.Right, visited)
	}
	return visited
}

func main() {

}
