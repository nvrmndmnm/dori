package main

func btInWalk(curr *BinaryNode) []int {
	path := []int{}

	if curr != nil {
		path = append(path, btInWalk(curr.left)...)
		path = append(path, curr.Value)
		path = append(path, btInWalk(curr.right)...)
	}

	return path
}

func InOrderTraversal(head *BinaryNode) []int {
	return btInWalk(head)
}
