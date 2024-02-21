package main

func btPreWalk(curr *BinaryNode) []int {
	path := []int{}

	if curr != nil {
		path = append(path, curr.Value)
		path = append(path, btPreWalk(curr.left)...)
		path = append(path, btPreWalk(curr.right)...)
	}

	return path
}

func PreOrderTraversal(head *BinaryNode) []int {
	return btPreWalk(head)
}
