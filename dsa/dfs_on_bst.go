package main

func search(curr *BinaryNode, needle int) bool {
	if curr == nil {
		return false
	}

	if curr.Value == needle {
		return true
	}

	if curr.Value < needle {
		return search(curr.right, needle)
	}

	return search(curr.left, needle)
}

func DFS(head *BinaryNode, needle int) bool {
	return search(head, needle)
}
