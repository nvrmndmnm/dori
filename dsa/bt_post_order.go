package main

func btPostWalk(curr *BinaryNode) []int {
	path := []int{}

	if curr != nil {
		path = append(path, btPostWalk(curr.left)...)
		path = append(path, btPostWalk(curr.right)...)
		path = append(path, curr.Value)
	}

	return path
}

func PostOrderTraversal(head *BinaryNode) []int {
	return btPostWalk(head)
}
