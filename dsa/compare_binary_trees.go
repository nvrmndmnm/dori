package main

func Compare(a *BinaryNode, b *BinaryNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Value != b.Value {
		return false
	}

	return Compare(a.left, b.left) && Compare(a.right, b.right)
}
